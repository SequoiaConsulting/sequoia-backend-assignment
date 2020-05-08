package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/model"
	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	listenPort int
	debugLogs  bool
)

func init() {
	pflag.BoolVar(&debugLogs, "debug", false, "Enable verbose debug logs with pretty printing")
	pflag.IntVar(&listenPort, "port", 3000, "HTTP server port")
	pflag.Parse()
}

func main() {
	initLoggerConfig()
	log.Info().Msg("initializing database connection")
	db, err := initDbConn()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to open the database connection")
	}

	defer db.Close()
	if err = model.AutoMigrate(db); err != nil {
		log.Fatal().Err(err).Msg("unable to migrate database schemas")
	}

	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// start http server on the main thread
	log.Info().Int("port", listenPort).Msg("listening for HTTP requests")
	err = http.ListenAndServe(fmt.Sprintf(":%d", listenPort), router)
	log.Fatal().Err(err).Msg("HTTP server exited with error")
}

// initLoggerConfig initializes global config for zerolog
func initLoggerConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debugLogs { // enable debug logging level and disable structured logging
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

// initDbConn initializes and configures the gorm.DB instance.
func initDbConn() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", getDBConnStrFromEnv())
	if err != nil {
		return nil, err
	}

	db.SetLogger(&utils.GORMLogger{})                 // only outputs logs if zerolog is set to debug level
	db = db.Set("gorm:association_autocreate", false) // disable association auto create
	db = db.Set("gorm:association_autoupdate", false) // disbale association auto update
	db = db.LogMode(true)
	return db, nil
}

// getDBConnStrFromEnv generates postgres connection string from env vars
func getDBConnStrFromEnv() string {
	host := requireEnvVar("PG_HOST")
	port := requireEnvVar("PG_PORT")
	user := requireEnvVar("PG_USER")
	password := requireEnvVar("PG_PASSWORD")
	dbname := requireEnvVar("PG_DBNAME")
	sslmode := requireEnvVar("PG_SSLMODE")
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
}

// requireEnvVar either returns value for env variable with given key, or panics
func requireEnvVar(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().
			Str("key", key).
			Msg("unable to find the environment variable")
	}

	return val
}
