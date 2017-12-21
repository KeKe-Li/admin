package controllers

import (
	"net/http"
	"strconv"

	"admin/models"
	"admin/helpers/sql"

	"github.com/gin-gonic/gin"

)



type PostsController struct {
	ApplicationController
}

func (ctrl PostsController) Index(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	pagenub := c.Query("pagenub")
	ipageNub, err := strconv.Atoi(pagenub)
	if err != nil {
		ipageNub = 20
	}
	pagination := sql.Pagination{}
	pagination.PerPage = ipageNub
	pagination.UrlQuery = c.Request.URL.Query()
	pagination.Path = c.Request.RequestURI
	pagination.Query = ctrl.DB.Model(&models.Post{})
	query := pagination.Paginate(page)

	var posts []models.Post
	err := query.Find(&posts).Error

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "posts/index", gin.H{
			"title": "index",
			"posts": posts,
		})
	}

}

func (ctrl PostsController) Show(c *gin.Context)  {
	var post models.Post
	err := ctrl.DB.First(&post, c.Param("id")).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "posts/show",gin.H{
			"title":"show",
			"post":post,
		 })
	}

}

func (ctrl PostsController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer,http.StatusOK, "posts/new", gin.H{
		"title": "new",
	})
}

func (ctrl PostsController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	body := c.PostForm("body")
	isPublished, err := strconv.ParseBool(c.PostForm("isPublished"))
	title := c.PostForm("title")

	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}
	post := models.Post{
						Body: body,
						IsPublished: isPublished,
						Title: title,
		}
	err = ctrl.DB.Create(&post).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}


}

func (ctrl PostsController) Edit(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	post := models.Post{ID: uint(id)}

	err = ctrl.DB.First(&post).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "posts/edit", gin.H{
			"title": "edit",
			"post":  post,
		})
	}
}

func (ctrl PostsController) Update(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	body := c.PostForm("body")
	id, err := strconv.ParseUint(c.Param("id"),10,0)
	isPublished, err := strconv.ParseBool(c.PostForm("isPublished"))
	title := c.PostForm("title")
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	return
	}

	updateColumns := make(map[string]interface{})
	
	updateColumns["body"]=body
	updateColumns["is_published"]=isPublished
	updateColumns["title"]=title

	post := &models.Post{
		    ID:  uint(id),
	    }
	err = ctrl.DB.Model(post).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl PostsController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	post := &models.Post{}

	err = ctrl.DB.Model(post).Delete(post, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

