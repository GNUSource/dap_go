package conf

import (
	"constant"
	"env"
	"fmt"
)

var (
	user     string
	password string
	address  string
	protocol string
	database string
)

func GetDsn() string {
	env := env.Env
	switch env {
	case constant.ENV_DEV:
		user = "dev"
		password = "bj.ai"
		protocol = "tcp"
		address = "106.15.33.15:3306"
		database = "data"
	case constant.ENV_BETA:
		user = "dev"
		password = "bj.ai"
		protocol = "tcp"
		address = "106.15.33.15:3306"
		database = "dcm_data"
	case constant.ENV_PRO:
		user = "dev"
		password = "bj.ai"
		protocol = "3306"
		address = "106.15.33.15:3306"
		database = "dcm_data_bl"

	}
	// dsn format：username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, protocol, address, database)
	return dsn
}
