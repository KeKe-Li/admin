package controllers

import (
	"net/http"
	"strconv"

	"admin/models"
	"admin/helpers/sql"

	"github.com/gin-gonic/gin"

)

type TagsController struct {
	ApplicationController
}

func (ctrl TagsController) Index(c *gin.Context) {
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
	pagination.Query = ctrl.DB.Model(&models.Tag{})
	query := pagination.Paginate(page)

	var tags []models.Tag
	err := query.Find(&tags).Error

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "tags/index", gin.H{
			"title": "index",
			"tags": tags,
		})
	}

}

func (ctrl TagsController) Show(c *gin.Context)  {
	var tag models.Tag
	err := ctrl.DB.First(&tag, c.Param("id")).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "tags/show",gin.H{
			"title":"show",
			"tag":tag,
		 })
	}

}

func (ctrl TagsController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer,http.StatusOK, "tags/new", gin.H{
		"title": "new",
	})
}

func (ctrl TagsController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	name := c.PostForm("name")

	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}
	tag := models.Tag{
						Name: name,
		}
	err = ctrl.DB.Create(&tag).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}


}

func (ctrl TagsController) Edit(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	tag := models.Tag{ID: uint(id)}

	err = ctrl.DB.First(&tag).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "tags/edit", gin.H{
			"title": "edit",
			"tag":  tag,
		})
	}
}

func (ctrl TagsController) Update(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	id, err := strconv.ParseUint(c.Param("id"),10,0)
	name := c.PostForm("name")
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	return
	}

	updateColumns := make(map[string]interface{})
	
	updateColumns["name"]=name

	tag := &models.Tag{
		    ID:  uint(id),
	    }
	err = ctrl.DB.Model(tag).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl TagsController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	tag := &models.Tag{}

	err = ctrl.DB.Model(tag).Delete(tag, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

