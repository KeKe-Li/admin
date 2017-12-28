package config

import (
	"fmt"

	"payment/helpers/configurations"

	"payment/helpers/utils/qiniu"
)

type ApplicationConfig struct {
	ENV          string
	Port         string
	SessionStore map[string]string
	Qiniu        map[string]*qiniu.Config
	Redis        *configurations.Redis
	WechatPay    *configurations.WechatPay
	//Mysql        *configurations.Mysql
	Sqlite       string
}

// 通过BUCKET获取对应七牛配置
func (conf *ApplicationConfig) GetQiniu(bucket string) (*qiniu.Config, error) {
	if qiniu := conf.Qiniu[bucket]; qiniu == nil {
		return nil, fmt.Errorf("qiniu bucket missing")
	} else {
		return qiniu, nil
	}
}

