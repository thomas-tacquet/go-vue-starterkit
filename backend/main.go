package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/thomas-tacquet/go-vue-starterkit/backend/server"
)

func main() {
	api := &server.API{
		Router: gin.New(),
		Config: viper.New(),
	}

	if err := api.SetupViper(); err != nil {
		panic(err)
	}
}
