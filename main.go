package main

import "C"
import (
	"GdalProject/global"
	"GdalProject/mapper"
	"GdalProject/routers"
	"GdalProject/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strconv"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	err := mapper.ConnectDB()
	if err != nil {
		log.Printf("%+v", err)
	}
	for i := 1; i <= global.FileSetting.Num; i++ {
		service.ReadTif(strconv.Itoa(i))
	}
	defer mapper.Close()
	f, _ := os.Create("./logs/logs.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
