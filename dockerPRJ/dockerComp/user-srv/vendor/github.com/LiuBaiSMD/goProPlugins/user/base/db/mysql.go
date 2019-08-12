package db

import (
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"github.com/LiuBaiSMD/goProPlugins/user/base/config"
	"os"
)

func initMysql() {

	var err error

	// 创建连接
	var mysqlUrl string
	if 	dockerMode := os.Getenv("RUN_DOCKER_MODE");dockerMode == "on"{
		mysqlUrl = config.GetMysqlConfig().GetDockerURL()
	}else{
		mysqlUrl = config.GetMysqlConfig().GetURL()
	}
	mysqlDB, err = sql.Open("mysql", mysqlUrl)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
