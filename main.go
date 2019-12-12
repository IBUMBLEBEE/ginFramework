// Package REST API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
// 	   Contact: IBUMBLEBEE<ibumblebeet@gmail.com>
//     Site: https://neutroncourse.xyz
//
//     Consumes:
//     - application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ginFramework/conf"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ginConfig     *conf.Config
	ginConfigData []byte
)

var (
	// the variables will be set at compile time from golang build ldflags
	gService   = "service"
	gVersion   = "version"
	gBuildData = "build_data"
	gCommmitID = "commit_id"
)

func initConfig() (*conf.Config, []byte, error) {
	var (
		cfgFile  string
		etcdAddr string
	)
	flag.StringVar(&cfgFile, "c", "", "config file")
	flag.StringVar(&etcdAddr, "etcd", "", "seperated with [,] if you have multiple addresses")
	flag.Parse()

	cfg := &conf.Config{}
	data, err := config.GetCfgFromEtcdOrFile(etcdAddr, cfgFile, gVersion, cfg)
	return cfg, data, err
}

// startHttpServer 始于http
func startHTTPServer(cfgData []byte, addr string) (*gin.Engine, error) {
	mode := viper.GetString("mode")
	gin.SetMode(mode)

	// begin Http server
	httpServer, err := gin.Default()
	if err != nil {
		fmt.Printf("httpServer error: %v", err)
	}
	return nil, httpServer
}

func main() {
	var err error
	ginConfig, ginConfigData, err = initConfig()
	if err != nil {
		fmt.Printf("initConfig error:%v", err)
		return
	}
	conf.SetConfig(ginConfig)

	gengine, err := startHTTPServer(ginConfigData, ginConfig.ListenAddr)
	if err != nil {
		log.Error("startHttpServer error:", err)
		return
	}
	// graceful exit
	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch. so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdowm Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := gengine.Shutdowm(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	log.Println("Server exiting")
}
