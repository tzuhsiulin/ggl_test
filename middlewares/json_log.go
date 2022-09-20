package middlewares

import (
	"strings"
	"time"

	"ggl_test/utils/log"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now().UnixMicro()

		// Process Request
		c.Next()

		// Stop timer
		end := time.Now().UnixMicro()
		duration := end - start

		entry := log.GetLogger().
			With("clientIp", getClientIP(c)).
			With("method", c.Request.Method).
			With("path", c.Request.RequestURI).
			With("status", c.Writer.Status()).
			With("referrer", c.Request.Referer()).
			With("requestId", requestid.Get(c)).
			With("duration", duration)

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

func getClientIP(c *gin.Context) string {
	// first check the X-Forwarded-For header
	requester := c.Request.Header.Get("X-Forwarded-For")
	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}
