package controllers

import (
	"net/http"
	"strconv"

	"admin/models"
	"admin/helpers/sql"

	"github.com/gin-gonic/gin"

)



type UsersController struct {
	ApplicationController
}

func (ctrl UsersController) Index(c *gin.Context) {
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
	pagination.Query = ctrl.DB.Model(&models.User{})
	query := pagination.Paginate(page)

	var users []models.User
	err := query.Find(&users).Error

	if err != nil{
		return
	}

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "users/index", gin.H{
			"title": "index",
			"users": users,
		})
	}

}

func (ctrl UsersController) Show(c *gin.Context)  {
	var user models.User
	err := ctrl.DB.First(&user, c.Param("id")).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "users/show",gin.H{
			"title":"show",
			"user":user,
		 })
	}

}

func (ctrl UsersController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer,http.StatusOK, "users/new", gin.H{
		"title": "new",
	})
}

func (ctrl UsersController) Create(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	avatarUrl := c.PostForm("avatarUrl")
	email := c.PostForm("email")
	encryptedPassword := c.PostForm("encryptedPassword")
	githubLoginId := c.PostForm("githubLoginId")
	githubUrl := c.PostForm("githubUrl")
	isAdmin, err := strconv.ParseBool(c.PostForm("isAdmin"))
	lockState, err := strconv.ParseBool(c.PostForm("lockState"))
	nickName := c.PostForm("nickName")
	secretKey := c.PostForm("secretKey")
	telephone := c.PostForm("telephone")
	verifyState := c.PostForm("verifyState")

	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}
	user := models.User{
						AvatarUrl: avatarUrl,
						Email: email,
						EncryptedPassword: encryptedPassword,
						GithubLoginId: githubLoginId,
						GithubUrl: githubUrl,
						IsAdmin: isAdmin,
						LockState: lockState,
						NickName: nickName,
						SecretKey: secretKey,
						Telephone: telephone,
						VerifyState: verifyState,
		}
	err = ctrl.DB.Create(&user).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}


}

func (ctrl UsersController) Edit(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}

	user := models.User{ID: uint(id)}

	err = ctrl.DB.First(&user).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "users/edit", gin.H{
			"title": "edit",
			"user":  user,
		})
	}
}

func (ctrl UsersController) Update(c *gin.Context) {
	resp := ctrl.NewResponse()
	
	avatarUrl := c.PostForm("avatarUrl")
	email := c.PostForm("email")
	encryptedPassword := c.PostForm("encryptedPassword")
	githubLoginId := c.PostForm("githubLoginId")
	githubUrl := c.PostForm("githubUrl")
	id, err := strconv.ParseUint(c.Param("id"),10,0)
	isAdmin, err := strconv.ParseBool(c.PostForm("isAdmin"))
	lockState, err := strconv.ParseBool(c.PostForm("lockState"))
	nickName := c.PostForm("nickName")
	secretKey := c.PostForm("secretKey")
	telephone := c.PostForm("telephone")
	verifyState := c.PostForm("verifyState")
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	return
	}

	updateColumns := make(map[string]interface{})
	
	updateColumns["avatar_url"]=avatarUrl
	updateColumns["email"]=email
	updateColumns["encrypted_password"]=encryptedPassword
	updateColumns["github_login_id"]=githubLoginId
	updateColumns["github_url"]=githubUrl
	updateColumns["is_admin"]=isAdmin
	updateColumns["lock_state"]=lockState
	updateColumns["nick_name"]=nickName
	updateColumns["secret_key"]=secretKey
	updateColumns["telephone"]=telephone
	updateColumns["verify_state"]=verifyState

	user := &models.User{
		    ID:  uint(id),
	    }
	err = ctrl.DB.Model(user).UpdateColumns(updateColumns).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

func (ctrl UsersController) Delete(c *gin.Context) {
	resp := ctrl.NewResponse()

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
		return
	}

	user := &models.User{}

	err = ctrl.DB.Model(user).Delete(user, id).Error
	if err != nil {
		resp.ErrorCode = 1
		resp.ErrorMessage = err.Error()
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	} else {
		resp.ErrorCode = 0
		ctrl.Render.JSON(c.Writer, http.StatusOK, resp)
	}
}

