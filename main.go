package main

import (
	"fmt"
	"github.com/LaiJunBin/gin-api/global"
	"github.com/LaiJunBin/gin-api/internal/model"
	"github.com/LaiJunBin/gin-api/internal/routers"
	"github.com/LaiJunBin/gin-api/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


func init() {
	err := setupSetting()
	if err != nil {
		fmt.Printf("init.setupSetting error: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		fmt.Printf("init.setupDBEngine error: %v", err)
	}
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
	}

	s.ListenAndServe()
}
