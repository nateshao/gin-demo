package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

//func main() {
//	r := gin.Default()
// 1，引用。2，渲染
//	r.LoadHTMLGlob("templates/**/*")
//	r.GET("/post/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK,"posts/index.html",gin.H{
//			"title":"posts/index",
//		})
//	})
//	r.GET("users/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK,"users/index.html",gin.H{
//			"title":"users/index",
//		})
//
//	})
//	r.Run(":8080")
//}
//func main() {
//	r := gin.Default()
//	r.LoadHTMLGlob("templates/**/*")
//	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
//	r.GET("/posts/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "posts/index.html", gin.H{
//			"title": "posts/index",
//		})
//	})
//
//	r.GET("users/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "users/index.html", gin.H{
//			"title": "users/index",
//		})
//	})
//
//	r.Run(":8080")
//}

/**************************************************************************************/
//func main() {
//	router := gin.Default()
//	router.SetFuncMap(template.FuncMap{
//		"safe": func(str string) template.HTML{
//			return template.HTML(str)
//		},
//	})
//	router.LoadHTMLFiles("./index.tmpl")
//
//	router.GET("/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://nateshao.gitee.io'>千羽的博客</a>")
//	})
//
//	router.Run(":8080")
//}

/********************************************************************************/
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func indexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func homeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)
	r.Run()
}
