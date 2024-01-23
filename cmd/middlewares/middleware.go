package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func ValidateRequest(c *gin.Context) {
	// Validate the request

	// Sanitize inputs

	// Call the next middleware or endpoint handler
	c.Next()
}

func TraceRequest(c *gin.Context) {
	// Do some request tracing before request is processed
	beforeRequest(c)

	// Call the next middleware or endpoint handler
	c.Next()

	// Do some request tracing after request is processed
	afterRequest(c)
}

func SecureRequest(c *gin.Context) {
	// Add some security features to the request

	// Call the next middleware or endpoint handler
	c.Next()
}

func AuthenticateRequest(c *gin.Context) {
	// Authenticate the request

	// Call the next middleware or endpoint handler
	c.Next()
}

func beforeRequest(c *gin.Context) {
	start := time.Now()

	// Log the request start time
	fmt.Printf("Started %s %s\n", c.Request.Method, c.Request.URL.Path)

	// Add start time to the request context
	c.Set("startTime", start)
}

func afterRequest(c *gin.Context) {
	// Get the start time from the request context
	startTime, exists := c.Get("startTime")
	if !exists {
		startTime = time.Now()
	}

	// Calculate the request duration
	duration := time.Since(startTime.(time.Time))

	// Log the request completion time and duration
	fmt.Printf("Completed %s %s in %v\n", c.Request.Method, c.Request.URL.Path, duration)
}
