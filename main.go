package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pwhb/lear/pkg/router"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	router.ConfigurePoemsRoutes(r)
	r.Run()

}

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*")
// 	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
// 	router.GET("/index", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.tmpl", gin.H{
// 			"title": "Main website",
// 		})
// 	})
// 	router.Run(":8080")
// }
