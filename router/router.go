package router

import (
	v1 "ginchat/api/v1"
	"ginchat/docs"
	"ginchat/pkg/global/log"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	r.Use(Recovery)
	// 静态资源
	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")
	// 首页
	r.GET("/", v1.Index.GetIndex)
	r.GET("/index", v1.Index.GetIndex)
	r.GET("/toRegister", v1.Index.ToRegister)
	r.GET("/toChat", v1.Index.ToChat)
	r.GET("/chat", v1.Index.Chat)
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiV1 := r.Group("/api/v1")
	{
		// 用户
		eg := apiV1.Group("/user")
		{
			eg.POST("/createUser", v1.User.CreateUser)
			eg.GET("/deleteUser", v1.User.DeleteUser)
			eg.POST("/updateUser", v1.User.UpdateUser)
			eg.POST("/login", v1.User.Login)
			// 查询好友
			eg.POST("/searchFriend", v1.User.SearchFriend)
			eg.POST("/findByUserId", v1.User.FindByUserId)

			eg.POST("/getMsgByRedis", v1.User.GetMsgByRedis)
		}
		// 文件上传
		attach := apiV1.Group("/attach")
		{
			attach.POST("/upload", v1.Attach.Upload)
		}
		// 用户关系
		contact := apiV1.Group("/contact")
		{
			contact.POST("/addFriend", v1.Contact.AddFriend)
			contact.POST("/createGroup", v1.Contact.CreateGroup)
			contact.POST("/findGroup", v1.Contact.FindGroup)
			contact.POST("/addGroup", v1.Contact.AddGroup)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			var m any = nil
			if err := recover(); err != m {
				log.Logger.Error("HttpError", zap.Any("HttpError", err))
			}
		}()

		c.Next()
	}
}

func Recovery(c *gin.Context) {
	defer func() {
		var m any = nil
		if r := recover(); r != m {
			log.Logger.Error("gin catch error: ", log.Any("gin catch error: ", r))
			utils.RespFail(c.Writer, "系统内部错误")
		}
	}()
	c.Next()
}
