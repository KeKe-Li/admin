package controllers

import (
	"net/http"
	"strconv"

	"admin/models"
	"admin/helpers/sql"

	"github.com/gin-gonic/gin"

)

type CommentsController struct {
	ApplicationController
}

func (ctrl CommentsController) Index(c *gin.Context) {
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
	pagination.Query = ctrl.DB.Model(&models.Comment{})
	query := pagination.Paginate(page)

	var comments []models.Comment
	err := query.Find(&comments).Error

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "comments/index", gin.H{
			"title": "index",
			"comments": comments,
		})
	}

}

func (ctrl CommentsController) Show(c *gin.Context)  {
	var comment models.Comment
	err := ctrl.DB.First(&comment, c.Param("id")).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "comments/show",gin.H{
			"title":"show",
			"comment":comment,
		 })
	}

}

func (ctrl CommentsController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer,http.StatusOK, "comments/new", gin.H{
		"title": "new",
	})
}

func (ctrl CommentsController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	avatarUrl := c.PostForm("avatarUrl")
	content := c.PostForm("content")
	githubUrl := c.PostForm("githubUrl")
	nickName := c.PostForm("nickName")
	postId, err := strconv.ParseUint(c.PostForm("postId"),10,0)
	readState, err := strconv.ParseBool(c.PostForm("readState"))
	userId, err := strconv.ParseUint(c.PostForm("userId"),10,0)

	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}
	comment := models.Comment{
						AvatarUrl: avatarUrl,
						Content: content,
						GithubUrl: githubUrl,
						NickName: nickName,
						PostID: uint(postId),
						ReadState: readState,
						UserID: uint(userId),
		}
	err = ctrl.DB.Create(&comment).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}


}

func (ctrl CommentsController) Edit(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	comment := models.Comment{ID: uint(id)}

	err = ctrl.DB.First(&comment).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "comments/edit", gin.H{
			"title": "edit",
			"comment":  comment,
		})
	}
}

func (ctrl CommentsController) Update(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	avatarUrl := c.PostForm("avatarUrl")
	content := c.PostForm("content")
	githubUrl := c.PostForm("githubUrl")
	id, err := strconv.ParseUint(c.Param("id"),10,0)
	nickName := c.PostForm("nickName")
	postId, err := strconv.ParseUint(c.PostForm("postId"),10,0)
	readState, err := strconv.ParseBool(c.PostForm("readState"))
	userId, err := strconv.ParseUint(c.PostForm("userId"),10,0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	return
	}

	updateColumns := make(map[string]interface{})
	
	updateColumns["avatar_url"]=avatarUrl
	updateColumns["content"]=content
	updateColumns["github_url"]=githubUrl
	updateColumns["nick_name"]=nickName
	updateColumns["post_id"]=postId
	updateColumns["read_state"]=readState
	updateColumns["user_id"]=userId

	comment := &models.Comment{
		    ID:  uint(id),
	    }
	err = ctrl.DB.Model(comment).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl CommentsController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	comment := &models.Comment{}

	err = ctrl.DB.Model(comment).Delete(comment, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

