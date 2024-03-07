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
	// r.Run(":8000")

// 	tr := router.TrCollectRoute()
	// tr.Run(":8002")

	server := http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: r,
		// TLSConfig: GetServerTlsConfig(),
	}

// 	twoWayServer := http.Server{
// 		Addr:      ":8002",
// 		Handler:   tr,
// 		TLSConfig: common.GetServerTlsConfig(),
// 	}

	wg.Add(1)
	go func() {
		// server.ListenAndServeTLS("config/certificate/servera/localhost.pem", "config/certificate/servera/localhost.key")
		server.ListenAndServe()
		wg.Done()
	}()

// 	wg.Add(1)
// 	go func() {
// 		twoWayServer.ListenAndServeTLS("config/certificate/servera/localhost.pem", "config/certificate/servera/localhost.key")
// 		wg.Done()
// 	}()

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
