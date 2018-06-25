package routers

import (
	"net/http"
	"test/testgin/controllers"

	"github.com/gin-gonic/gin"
)

// 跨域
const (
	CORSHeaders = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
	CORSMethods = "POST, GET, OPTIONS, PUT, DELETE, UPDATE"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", CORSHeaders)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", CORSMethods)
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
func SetupRouter() *gin.Engine {

	//创建gin引擎
	router := gin.Default()

	//静态
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("/", "static/client.html")

	testgin := router.Group("/testgin")
	{
		//body编码
		testgin.POST("/urlencode", controllers.URLEncode)
		testgin.POST("/json", controllers.Json)

		//跨域jsonp
		testgin.GET("/jsonp", controllers.Jsonp)

		//重定向
		testgin.GET("/exredirect", CORSMiddleware(), func(ctx *gin.Context) {
			// ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
			ctx.Redirect(http.StatusFound, "https://www.baidu.com")
		})
		testgin.GET("/inredirect", func(ctx *gin.Context) {
			ctx.Request.URL.Path = "/testgin/inredirect2"
			router.HandleContext(ctx)
		})
		testgin.GET("/inredirect2", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
		})

		//基础授权中间件
		var secrets = gin.H{
			"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
			"austin": gin.H{"email": "austin@example.com", "phone": "666"},
			"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
		}
		auth := testgin.Group("/baseauth", gin.BasicAuth(gin.Accounts{
			"foo":    "bar",
			"austin": "1234",
			"lena":   "hello",
			"manu":   "4321",
		}))
		{
			auth.GET("/secrets", func(ctx *gin.Context) {
				u := ctx.MustGet(gin.AuthUserKey).(string)
				if secret, ok := secrets[u]; ok {
					ctx.JSON(http.StatusOK, gin.H{"user": u, "secrets": secret})
				} else {
					ctx.JSON(http.StatusOK, gin.H{"user": u, "secret": "NO SECRET :("})
				}
			})
		}
	}

	//点赞
	testorm := router.Group("/testorm")
	{
		testorm.GET("/likeinfo", controllers.LikeInfo)
		testorm.POST("/like", controllers.Like)
	}

	return router
}
