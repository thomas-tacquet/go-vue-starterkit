package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thinkerou/favicon"

	"github.com/thomas-tacquet/go-vue-starterkit/backend/helpers"
)

var lastModTime time.Time
var indexFile []byte

// InitializeRoutes main route initializer
func (a *API) InitializeRoutes(testMode bool, db *gorm.DB) {
	r := a.Router
	r.Use(gin.Recovery())
	r.NoRoute(NoRoute)

	if !testMode {
		r.Use(static.Serve("/assets/", static.LocalFile("./dist/", false)))
		r.Use(favicon.New("./frontend/assets/favicon/favicon.ico"))
	}

	api := r.Group("/api")
	{
		v1NoAuth := api.Group("/v1")
		{
			v1NoAuth.GET("/", index)
		}
	}
}

// NoRoute is called when no route found
func NoRoute(c *gin.Context) {

	if c.Request.URL.Path == "/favicon.ico" {
		return
	}

	path := strings.Split(c.Request.URL.Path, "/")

	if path[1] == "api" {
		helpers.SendError(c, helpers.ErrNotFound(), nil, errors.New("URL not found : "+c.Request.URL.Path))
	} else {
		var info os.FileInfo
		var err error
		if info, err = os.Stat("./dist/index.html"); err != nil {
			helpers.SendError(c, helpers.ErrInternalServerError(), nil, err)
			return
		}

		if indexFile == nil || !lastModTime.Equal(info.ModTime()) {
			lastModTime = info.ModTime()
			fmt.Println("Reloading index.html")
			indexFile, err = ioutil.ReadFile("./dist/index.html")
			if err != nil {
				helpers.SendError(c, helpers.ErrInternalServerError(), helpers.MakeEasyDetail(helpers.URL, strings.Join(path, "/")), err)
				return
			}
		}
		c.Data(http.StatusOK, gin.MIMEHTML, indexFile)
	}
}

func index(c *gin.Context) {
	helpers.SendData(c, http.StatusOK, gin.H{"status": "success", "title": "You Successfully reached the API !"})
}
