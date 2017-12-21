package controllers

import (
	"net/http"
	"strconv"

	"admin/models"
	"admin/helpers/sql"

	"github.com/gin-gonic/gin"

)


type SubscribersController struct {
	ApplicationController
}

func (ctrl SubscribersController) Index(c *gin.Context) {
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
	pagination.Query = ctrl.DB.Model(&models.Subscriber{})
	query := pagination.Paginate(page)

	var subscribers []models.Subscriber
	err := query.Find(&subscribers).Error

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "subscribers/index", gin.H{
			"title": "index",
			"subscribers": subscribers,
		})
	}

}

func (ctrl SubscribersController) Show(c *gin.Context)  {
	var subscriber models.Subscriber
	err := ctrl.DB.First(&subscriber, c.Param("id")).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "subscribers/show",gin.H{
			"title":"show",
			"subscriber":subscriber,
		 })
	}

}

func (ctrl SubscribersController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer,http.StatusOK, "subscribers/new", gin.H{
		"title": "new",
	})
}

func (ctrl SubscribersController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	email := c.PostForm("email")
	secretkey := c.PostForm("secretkey")
	signature := c.PostForm("signature")
	subscribeState, err := strconv.ParseBool(c.PostForm("subscribeState"))
	verifyState, err := strconv.ParseBool(c.PostForm("verifyState"))

	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}
	subscriber := models.Subscriber{
						Email: email,
						Secretkey: secretkey,
						Signature: signature,
						SubscribeState: subscribeState,
						VerifyState: verifyState,
		}
	err = ctrl.DB.Create(&subscriber).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}


}

func (ctrl SubscribersController) Edit(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	subscriber := models.Subscriber{ID: uint(id)}

	err = ctrl.DB.First(&subscriber).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "subscribers/edit", gin.H{
			"title": "edit",
			"subscriber":  subscriber,
		})
	}
}

func (ctrl SubscribersController) Update(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	email := c.PostForm("email")
	id, err := strconv.ParseUint(c.Param("id"),10,0)
	secretkey := c.PostForm("secretkey")
	signature := c.PostForm("signature")
	subscribeState, err := strconv.ParseBool(c.PostForm("subscribeState"))
	verifyState, err := strconv.ParseBool(c.PostForm("verifyState"))
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	return
	}

	updateColumns := make(map[string]interface{})
	
	updateColumns["email"]=email
	updateColumns["secretkey"]=secretkey
	updateColumns["signature"]=signature
	updateColumns["subscribe_state"]=subscribeState
	updateColumns["verify_state"]=verifyState

	subscriber := &models.Subscriber{
		    ID:  uint(id),
	    }
	err = ctrl.DB.Model(subscriber).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl SubscribersController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	subscriber := &models.Subscriber{}

	err = ctrl.DB.Model(subscriber).Delete(subscriber, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

