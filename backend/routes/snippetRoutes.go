package routes

import (
	snippetControllers "vault-dev/controllers/Snippet_Controller"

	"github.com/gin-gonic/gin"
)

func Snippet_Routes(r *gin.Engine) {
	/*
		 api := r.Group("/api/auth")
		 api.POST("/login", controllers.Login)
		 api.POST("/register", controllers.Register)

		Adding thses later for login/register

	*/

	snippet := r.Group("api")

	snippet.POST("/snippet", snippetControllers.PostSnippet)
	snippet.GET("/snippet", snippetControllers.GetSnippets)
	snippet.GET("/snippet/favorites", snippetControllers.GetFavSnippets)
	snippet.PUT("snippet/:id", snippetControllers.UpdateSnippet)
	snippet.DELETE("/snippet/:id", snippetControllers.DeleteSnippet)

}
