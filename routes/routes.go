package routes

import (
	"github.com/gin-gonic/gin"
	"admin/controllers"
)

func InitRoutes(engine *gin.RouterGroup,appCtrl controllers.ApplicationController)  {
	router := engine.Group("")
	{
		//Comment
		commentsCtrl := controllers.CommentsController{appCtrl}
		router.GET("/comments", commentsCtrl.Index)
		router.GET("/comments-show/:id", commentsCtrl.Show)
		router.GET("/comments/new",commentsCtrl.New)
		router.GET("/comments-edit/:id", commentsCtrl.Edit)
		router.POST("/comments", commentsCtrl.Create)
		router.PUT("/comments/:id", commentsCtrl.Update)
		router.DELETE("/comments/:id", commentsCtrl.Delete)

		//Page
		pagesCtrl := controllers.PagesController{appCtrl}
		router.GET("/pages", pagesCtrl.Index)
		router.GET("/pages-show/:id", pagesCtrl.Show)
		router.GET("/pages/new",pagesCtrl.New)
		router.GET("/pages-edit/:id", pagesCtrl.Edit)
		router.POST("/pages", pagesCtrl.Create)
		router.PUT("/pages/:id", pagesCtrl.Update)
		router.DELETE("/pages/:id", pagesCtrl.Delete)

		//Post
		postsCtrl := controllers.PostsController{appCtrl}
		router.GET("/posts", postsCtrl.Index)
		router.GET("/posts-show/:id", postsCtrl.Show)
		router.GET("/posts/new",postsCtrl.New)
		router.GET("/posts-edit/:id", postsCtrl.Edit)
		router.POST("/posts", postsCtrl.Create)
		router.PUT("/posts/:id", postsCtrl.Update)
		router.DELETE("/posts/:id", postsCtrl.Delete)

		//Subscriber
		subscribersCtrl := controllers.SubscribersController{appCtrl}
		router.GET("/subscribers", subscribersCtrl.Index)
		router.GET("/subscribers-show/:id", subscribersCtrl.Show)
		router.GET("/subscribers/new",subscribersCtrl.New)
		router.GET("/subscribers-edit/:id", subscribersCtrl.Edit)
		router.POST("/subscribers", subscribersCtrl.Create)
		router.PUT("/subscribers/:id", subscribersCtrl.Update)
		router.DELETE("/subscribers/:id", subscribersCtrl.Delete)


		//Tag
		tagsCtrl := controllers.TagsController{appCtrl}
		router.GET("/tags", tagsCtrl.Index)
		router.GET("/tags-show/:id", tagsCtrl.Show)
		router.GET("/tags/new",tagsCtrl.New)
		router.GET("/tags-edit/:id", tagsCtrl.Edit)
		router.POST("/tags", tagsCtrl.Create)
		router.PUT("/tags/:id", tagsCtrl.Update)
		router.DELETE("/tags/:id", tagsCtrl.Delete)

		//User
		usersCtrl := controllers.UsersController{appCtrl}
		router.GET("/users", usersCtrl.Index)
		router.GET("/users-show/:id", usersCtrl.Show)
		router.GET("/users/new",usersCtrl.New)
		router.GET("/users-edit/:id", usersCtrl.Edit)
		router.POST("/users", usersCtrl.Create)
		router.PUT("/users/:id", usersCtrl.Update)
		router.DELETE("/users/:id", usersCtrl.Delete)

	}

}
