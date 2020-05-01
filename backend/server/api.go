package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// API accès au routeur
type API struct {
	Router *gin.Engine
	Config *viper.Viper
}
