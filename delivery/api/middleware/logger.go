package middleware

import (
	"net"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// Logger is our custom middleware for logging in a structured way
func Logger(logger logrus.FieldLogger) gin.HandlerFunc {
	hostname, _ := os.Hostname()
	localIP := GetLocalIP()
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errors := c.Errors.ByType(gin.ErrorTypePrivate).String()
		userAgent := c.Request.UserAgent()
		contentType := c.Request.Header.Get("content-type")
		timeFormatted := end.Format("2006-01-02 15:04:05")

		logger.WithFields(logrus.Fields{
			"time":         timeFormatted,
			"hostname":     hostname,
			"local-ip":     localIP,
			"method":       method,
			"path":         path,
			"content-type": contentType,
			"latency":      latency,
			"client-ip":    clientIP,
			"errors":       errors,
			"status":       statusCode,
			"user-agent":   userAgent,
		}).Info("http request")
	}
}

// GetLocalIP returns the non loopback local IP of the host
// http://stackoverflow.com/a/31551220/916440
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
