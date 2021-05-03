package main

func main() {

	// 普通路由
	/*
		r.GET("/index", func(c *gin.Context) {...})
		r.GET("/login", func(c *gin.Context) {...})
		r.POST("/login", func(c *gin.Context) {...})
	*/
	// 此外，还有一个可以匹配所有请求方法的Any方法如下：	r.Any("/test", func(c *gin.Context) {...})

	/*
		为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回views/404.html页面。

		r.NoRoute(func(c *gin.Context) {
				c.HTML(http.StatusNotFound, "views/404.html", nil)
			})
	*/

	/**************************** 路由组 ************************************/
	/*
		路由组
	我们可以将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由，这只是为了看着清晰，你用不用{}包裹功能上没什么区别。

	func main() {
		r := gin.Default()
		userGroup := r.Group("/user")
		{
			userGroup.GET("/index", func(c *gin.Context) {...})
			userGroup.GET("/login", func(c *gin.Context) {...})
			userGroup.POST("/login", func(c *gin.Context) {...})

		}
		shopGroup := r.Group("/shop")
		{
			shopGroup.GET("/index", func(c *gin.Context) {...})
			shopGroup.GET("/cart", func(c *gin.Context) {...})
			shopGroup.POST("/checkout", func(c *gin.Context) {...})
		}
		r.Run()
	}
	路由组也是支持嵌套的，例如：

	shopGroup := r.Group("/shop")
		{
			shopGroup.GET("/index", func(c *gin.Context) {...})
			shopGroup.GET("/cart", func(c *gin.Context) {...})
			shopGroup.POST("/checkout", func(c *gin.Context) {...})
			// 嵌套路由组
			xx := shopGroup.Group("xx")
			xx.GET("/oo", func(c *gin.Context) {...})
		}
	通常我们将路由分组用在划分业务逻辑或划分API版本时。
	*/

}
