package main

import (
	"log"
	"online-storage/lib"
	"online-storage/model/mysql"
	"online-storage/router"
)

func main() {
	serverConfig := lib.LoadServerConfig()
	mysql.InitDB(serverConfig)
	defer mysql.DB.Close()

	r := router.SetupRoute()

	r.LoadHTMLGlob("view/*")
	r.Static("/static", "./static")

	if err := r.Run(":80"); err != nil {
		log.Fatal("服务器启动失败...")
	}
}
