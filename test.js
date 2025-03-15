// const API_URL = "http://localhost:3000/protected"; 
// const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5AZXhhbXBsZS5jb20iLCJleHAiOjE3NDIwNDk2OTYsInBsYW4iOiJmcmVlIn0.OPF94bkOMnX0s92RtREfDGcPIQrtf7ObIqLmxcCyCfw"; 
// const MAX_REQUESTS = 250; 

// async function sendRequest(i) {
//     try {
//         const response = await fetch(API_URL, {
//             method: "GET",
//             headers: { "Authorization": `Bearer ${TOKEN}` }
//         });

//         console.log(`#${i} - Status: ${response.status}`);
//     } catch (error) {
//         console.log(`#${i} ❌ Request Failed:`, error.message);
//     }
// }

// async function testRateLimit() {
//     console.log(`🚀 Testing Rate Limit: Sending ${MAX_REQUESTS} requests...`);
    
//     for (let i = 1; i <= MAX_REQUESTS; i++) {
//         await sendRequest(i);
//     }
    
//     console.log(`✅ Test Completed`);
// }

// testRateLimit();



const API_URL = "http://localhost:3000/api/hotels"; // Targeting API Gateway
const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5AZXhhbXBsZS5jb20iLCJleHAiOjE3NDIxMzY1OTIsInBsYW4iOiJmcmVlIn0.nzQwh5Y1-LmxbBNpErfMpXHLqPkAilw3OeH5UDtuZBY"; // Replace with a valid JWT token
const MAX_REQUESTS = 50; // Adjust to test rate limiting
const API_KEY= "17ccd5efae37b3f7453339431932c663"

async function sendRequest(i) {
    try {
        const response = await fetch(API_URL, {
            method: "GET",
            headers: { "Authorization": `Bearer ${TOKEN}` , "X-API-Key" : API_KEY }
        });

        console.log(`#${i} - Status: ${response.status}`);

        if (response.status === 429) {
            console.log(`#${i} ❌ Rate limit exceeded!`);
        }

    } catch (error) {
        console.log(`#${i} ❌ Request Failed:`, error.message);
    }
}

async function testRateLimit() {
    console.log(`🚀 Testing Rate Limit on /api/*: Sending ${MAX_REQUESTS} requests...`);
    
    for (let i = 1; i <= MAX_REQUESTS; i++) {
        await sendRequest(i);
    }

    console.log(`✅ Test Completed`);
}

testRateLimit();
