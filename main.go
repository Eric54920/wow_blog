package main

import (
	"log"
	"wow_blog/config"
	"wow_blog/routes"
	"wow_blog/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库连接
	err := utils.InitDB()
	if err != nil {
		log.Fatalf("init mysql failed, err:%v\n", err)
		return
	}
	defer utils.Close()

	r := gin.Default()

	routes.SetupRouter(r)

	r.Run(":8000")
}
