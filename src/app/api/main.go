package main

import (
	cfg "echofy_backend/src/app/api/config"
	"echofy_backend/src/app/api/endpoints/routes"
	"echofy_backend/src/app/api/middlewares"
	"echofy_backend/src/infra/postgres"
	"flag"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

const defaultEnvFilePath = "./src/app/api/.env"

// NewAPI
// @title SPOTIFY RECORDS API
// @version 1.0
// @description Aplicação de artistas do spotify
// @contact.name DIT - IFAL
// @contact.email evs10@aluno.ifal.edu.br
// @BasePath /api
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	loadEnvIntoSystemFromFile()

	setupPostgres()
	Serve(
		cfg.Env().Server.Host,
		cfg.Env().Server.Port,
	)
}

func setupPostgres() {
	err := postgres.SetUpCredentials(
		cfg.Env().Postgres.User,
		cfg.Env().Postgres.Password,
		cfg.Env().Postgres.Name,
		cfg.Env().Postgres.Host,
		cfg.Env().Postgres.Port,
		cfg.Env().Postgres.SSLMode,
	)
	if err != nil {
		panic(err)
	}

	err = postgres.MigrateUp()
	if err != nil {
		panic(err)
	}
}

func loadEnvIntoSystemFromFile() {
	log.Info().Msg("loading environment variables...")

	envPathPtr := flag.String("env", "", "override environment variable path")
	flag.Parse()
	if *envPathPtr == "" {
		*envPathPtr = defaultEnvFilePath
	}

	err := godotenv.Load(*envPathPtr)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to load config from env")
	}
}

// server

func Serve(host string, port int) {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(getCORSSettings())

	apiGroup := router.Group("/api")

	rout := routes.New()
	rout.Load(apiGroup)

	address := getServerAddress(host, port)
	router.Logger.Fatal(router.Start(address))
}

func getServerAddress(host string, port int) string {
	return fmt.Sprintf("%v:%v", host, port)
}

func getCORSSettings() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:         middlewares.OriginInspectSkipper,
		AllowOriginFunc: middlewares.VerifyOrigin,
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
	})
}
