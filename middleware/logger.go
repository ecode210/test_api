package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Logger - This is how every API call will be logged in the console
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"%s [%s] %s %s %d %s",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}
