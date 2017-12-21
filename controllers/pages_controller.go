package controllers

import (
	"net/http"
	"strconv"

	"admin/models"
	"admin/helpers/sql"

	"github.com/gin-gonic/gin"

)



type PagesController struct {
	ApplicationController
}

func (ctrl PagesController) Index(c *gin.Context) {
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
	pagination.Query = ctrl.DB.Model(&models.Page{})
	query := pagination.Paginate(page)

	var pages []models.Page
	err := query.Find(&pages).Error

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "pages/index", gin.H{
			"title": "index",
			"pages": pages,
		})
	}

}

func (ctrl PagesController) Show(c *gin.Context)  {
	var page models.Page
	err := ctrl.DB.First(&page, c.Param("id")).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "pages/show",gin.H{
			"title":"show",
			"page":page,
		 })
	}

}

func (ctrl PagesController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer,http.StatusOK, "pages/new", gin.H{
		"title": "new",
	})
}

func (ctrl PagesController) Create(c *gin.Context) {
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
	page := models.Page{
						Body: body,
						IsPublished: isPublished,
						Title: title,
		}
	err = ctrl.DB.Create(&page).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}


}

func (ctrl PagesController) Edit(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	page := models.Page{ID: uint(id)}

	err = ctrl.DB.First(&page).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "pages/edit", gin.H{
			"title": "edit",
			"page":  page,
		})
	}
}

func (ctrl PagesController) Update(c *gin.Context) {
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

	page := &models.Page{
		    ID:  uint(id),
	    }
	err = ctrl.DB.Model(page).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl PagesController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	page := &models.Page{}

	err = ctrl.DB.Model(page).Delete(page, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

