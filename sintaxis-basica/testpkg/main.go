package main

// webview-go example
//import webview "github.com/webview/webview_go"
//
//func main() {
//	w := webview.New(true)
//	defer w.Destroy()
//	w.SetTitle("Hello, webview-go!")
//	w.SetSize(800, 600, webview.HintNone)
//	w.Navigate("https://github.com/webview/webview")
//	w.Run()
//}

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	fmt.Println("Corriendo APP")

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
