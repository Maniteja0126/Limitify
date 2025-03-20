package middleware

import (
	"limitify/config"
	"limitify/models"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func APIGateway() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.GetHeader("X-API-Key")

		if apikey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing API Key"})
			c.Abort()
			return
		}

		backendURL, err := config.RedisClient.Get(config.Ctx, "apikey:"+apikey).Result()
		if err != nil {
			var key models.APIKey
			if err := config.DB.Where("api_key = ?", apikey).First(&key).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
				c.Abort()
				return
			}
			backendURL = key.BackendUrl
			config.RedisClient.Set(config.Ctx, "apikey:"+apikey, backendURL, time.Hour*24)
		}

		failureKey := "failures:" + apikey
		failures, _ := config.RedisClient.Get(config.Ctx, failureKey).Int()
		if failures >= 5 {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service temporarily unavailable"})
			c.Abort()
			return
		}

		originalPath := c.Request.URL.Path
		forwardPath := strings.TrimPrefix(originalPath, "/api")
		forwardPath = strings.TrimLeft(forwardPath, "/")

		if !strings.HasSuffix(backendURL, "/") {
			backendURL += "/"
		}

		// fullURL := backendURL + forwardPath
		// fmt.Println("🔹 Forwarding request to:", fullURL)

		targetURL, err := url.Parse(backendURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid backend URL"})
			c.Abort()
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Proxy error: %v", err)
			config.RedisClient.Incr(config.Ctx, failureKey)
			config.RedisClient.Expire(config.Ctx, failureKey, time.Hour)

			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach backend service"})
		}

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = "/" + forwardPath

			authHeader := c.GetHeader("Authorization")
			if authHeader != "" {
				req.Header.Set("Authorization", authHeader)
			}

			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Referer", "https://yourdomain.com")
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Connection", "close")
			req.Header.Set("X-Forwarded-For", c.ClientIP())

			req.Header.Del("X-API-Key")

		}

		proxy.ModifyResponse = func(res *http.Response) error {
			if res.StatusCode >= 500 {
				log.Printf("Upstream service error: %d", res.StatusCode)
				config.RedisClient.Incr(config.Ctx, failureKey)
				config.RedisClient.Expire(config.Ctx, failureKey, time.Hour)
			}
			return nil
		}

		proxy.ServeHTTP(c.Writer, c.Request)

		analyticsKey := "analytics:" + apikey + ":" + time.Now().Format("200601021504")
		config.RedisClient.Incr(config.Ctx, analyticsKey)
		config.RedisClient.Expire(config.Ctx, analyticsKey, time.Hour*24)
	}
}
