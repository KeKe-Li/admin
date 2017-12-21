package controllers

import (
	"encoding/json"
	"fmt"

	"admin/config"

	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"github.com/garyburd/redigo/redis"
)

type ApplicationController struct {
	RedisPool   *redis.Pool
	DB          *gorm.DB
	Config      *config.ApplicationConfig
	Render      *render.Render
}

func (a ApplicationController) Log(v ...interface{}) {
	fmt.Println(v)
}

type ApiResponse struct {
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"response,omitempty"`
}

func (act ApplicationController) NewResponse() *ApiResponse {
	return &ApiResponse{}
}

func (a *ApiResponse) ToJSON() []byte {
	result, _ := json.Marshal(a)
	return result
}
