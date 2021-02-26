package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/api"
	"github.com/spiderman930706/gin_admin/global"
	"log"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if global.Config.DEBUG {
					log.Println("[Recovery from panic]", string(httpRequest), err, string(debug.Stack()))
				}
				api.ErrorWithMessage(err.(error).Error(), c)
			}
		}()
		c.Next()
	}
}
