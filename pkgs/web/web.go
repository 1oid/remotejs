package web

import (
	"fmt"
	"github.com/chromedp/cdproto/runtime"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebRunner(port string, callFunc func(t string) (*runtime.RemoteObject, error)) {
	gin.SetMode(gin.ReleaseMode)
	route := gin.New()
	route.Use(gin.Recovery())

	route.POST("/remote", func(context *gin.Context) {
		evalString := context.PostForm("eval")

		if result, err := callFunc(evalString); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"value": result.Value,
				"type":  result.Type,
			})
		}

	})

	addr := fmt.Sprintf(`0.0.0.0:%s`, port)
	fmt.Printf("[*] web server start: %s\n", addr)
	route.Run(addr)
}
