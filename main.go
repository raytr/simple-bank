package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/helper/database"
	"gibhub.com/raytr/simple-bank/initialization"
	"gibhub.com/raytr/simple-bank/middleware"
	"github.com/go-kit/kit/log"
)

func main() {
	cfg := config.Init("config", "yml")

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, cfg.Server.Port, "caller", log.DefaultCaller)

	db, err := database.InitDatabase(&cfg.DBConfig)
	if err != nil {
		panic(err.Error())
	}

	_ = logger.Log("msg", fmt.Sprintf("Listening at port %s", strconv.Itoa(cfg.Server.Port)))
	mux := initialization.InitRouting(db, cfg, logger)
	httpServer := http.Server{
		Addr:    *flag.String("listen", ":"+strconv.Itoa(cfg.Server.Port), "Listen address."),
		Handler: middleware.AddMiddleware(mux),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func readFlags() map[string]interface{} {
	flags := make(map[string]interface{})
	args := os.Args[1:]
	params := make(map[string]string)

	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			params[key] = value
		}
	}

	// Create a map to store key-value pairs
	dbhost := params["db_host"]
	if dbhost != "" {
		flags["db_host"] = dbhost
	}

	dbportStr := params["db_port"]
	if dbportStr != "" {
		dbport, err := strconv.Atoi(dbportStr)
		if err != nil {
			panic(err.Error())
		}
		flags["db_host"] = dbport
	}

	isTest := params["test"]
	if isTest != "" && isTest == "true" {
		flags["test"] = true
	}

	return flags
}
