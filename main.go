package main

import (
	"flag"
	"fmt"
	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/helper/b_log"
	"gibhub.com/raytr/simple-bank/helper/database"
	"gibhub.com/raytr/simple-bank/initialization"
	"gibhub.com/raytr/simple-bank/middleware"
	"net/http"
	"strconv"
)

func main() {
	cfg := config.Init("config", "yml")

	logger := b_log.NewLogger(cfg.Server.Name)

	db, err := database.InitDatabase(&cfg.DBConfig)
	if err != nil {
		panic(err.Error())
	}

	logger.Info(fmt.Sprintf("Listening at port %s", strconv.Itoa(cfg.Server.Port)))
	mux := initialization.InitRouting(db, cfg, logger)
	httpServer := http.Server{
		Addr:    *flag.String("listen", ":"+strconv.Itoa(cfg.Server.Port), "Listen address."),
		Handler: middleware.AddMiddleware(mux),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
