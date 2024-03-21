package main

import (
	"net/http"
	"os"
	"server/internal/common"
	"server/internal/router"
	"sync"

	"github.com/spf13/viper"
)

var wg sync.WaitGroup

func main() {
	InitConfig()
	_ = common.InitDB()

	r := router.CollectRoute()

	server := http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: r,
	}

	wg.Add(1)
	go func() {
		server.ListenAndServe()
		wg.Done()
	}()

	wg.Wait()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
