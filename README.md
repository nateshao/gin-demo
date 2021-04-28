# gin
gin一下

出自：https://www.liwenzhou.com/posts/Go/Gin_framework/#autoid-0-2-2

GET用来获取资源
POST用来新建资源
PUT用来更新资源
DELETE用来删除资源。

只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

func main() {
r := gin.Default()
r.GET("/book", func(c *gin.Context) {
c.JSON(200, gin.H{
"message": "GET",
})
})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
}