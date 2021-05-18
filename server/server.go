package server

import (
	"backend/controller"
	"backend/middleware"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"backend/docs"
)

// @title Swagger FullStack API
// @version 1.0
// @description This is a sample server FullStack server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /api/v1
func NewRouter() *gin.Engine {
	docs.SwaggerInfo.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo.Description = os.Getenv("SWAGGER_DESCRIPTION")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASEPATH")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:9000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// routes
	v1 := r.Group("/api/v1")
	{
		// register
		v1.POST("user/register", controller.UserRegister)

		// login
		v1.POST("user/login", controller.UserLogin)

		// login required
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// user routing
			auth.GET("users", controller.ListUser)
			auth.DELETE("user/logout", controller.UserLogout)
			auth.GET("user/:id", controller.UserProfile)
			auth.PUT("user", controller.UpdateUser)

			// tweet
			auth.POST("tweets", controller.CreateTweet)
			auth.GET("tweets", controller.ListTweet)
			auth.DELETE("tweet/:id", controller.DeleteTweet)
			auth.GET("tweet/:id", controller.TweetByID)

			// comment
			auth.POST("comments", controller.CreateComment)
			auth.DELETE("comment/:id", controller.DeleteComment)

			// follow
			auth.POST("follow", controller.CreateFollow)
			auth.DELETE("follow/:id", controller.DeleteFollow)

			// like
			auth.POST("like", controller.CreateLike)
			auth.DELETE("like/:id", controller.DeleteLike)

			// retweet
			auth.POST("retweet", controller.CreateRetweet)
			auth.DELETE("retweet/:id", controller.DeleteRetweet)
		}

	}
	return r
}

func StartServer() {
	r := NewRouter()
	r.Run(":9000")
}
