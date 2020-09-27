package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/thomas-tacquet/go-vue-starterkit/backend/helpers"
	"github.com/thomas-tacquet/go-vue-starterkit/backend/server"
	"github.com/thomas-tacquet/go-vue-starterkit/backend/store"
)

func main() {
	api := &server.API{
		Router: gin.New(),
		Config: viper.New(),
	}

	// setting up viper to read .env file
	if err := api.SetupViper(); err != nil {
		panic(err)
	}
	// setting up log
	var logs helpers.Logger
	if err := logs.InitWithViper(api.Config); err != nil {
		panic(err)
	}

	db := store.InitAndGetDB(true, helpers.InitWithViper(api.Config), logs.Logs)
	defer func() {
		if err := db.Close(); err != nil {
			logs.Logs.Errorf("couldn't close DB : %s", err.Error())
		}
	}()

	api.InitializeRoutes(false, db)

	srv := &http.Server{
		Addr:    ":" + api.Config.GetString("PORT"),
		Handler: api.Router,
	}

	go func() {
		err := srv.ListenAndServeTLS(
			api.Config.GetString(helpers.EnvPublicRSA),
			api.Config.GetString(helpers.EnvPrivateRSA),
		)

		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen :%s\n", err)
		}
	}()

	// ctrl+c catch
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGABRT)
	<-quit
	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown: %s \n", err)
	}
	select {
	case <-ctx.Done():
		fmt.Println("timeout of 2 seconds")
	}
	fmt.Println("Server exiting")

}
