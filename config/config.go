package config

import (
	"os"
	"fmt"
)


var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func PrintConfig() {
	fmt.Println("Loaded JWT Secret:", string(JWTSecret))
}
