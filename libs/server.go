package libs

import (
	"log"
	"net/http"
	"os"
	"pius/router"
	"time"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use: "serve",
	Short: "for running API",
	RunE: serve,
}

func corsHandler() *cors.Cors {
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},

		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
	})

	return c
}

func serve(cmd *cobra.Command, args []string) error {
	
	mainRoute := router.RouterApp()

	var address string = "0.0.0.0:" + os.Getenv("PORT")
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "0.0.0.0:" + PORT
	}

	cors := corsHandler()

	serve := &http.Server{
		Addr: address,
		WriteTimeout: time.Minute * 2,
		ReadTimeout: time.Minute * 2,
		IdleTimeout: time.Minute,
		Handler: cors.Handler(mainRoute),
	}

	log.Println("App is running on PORT", os.Getenv("PORT"))

	return serve.ListenAndServe()
}