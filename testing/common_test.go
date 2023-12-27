package testing

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/helper/b_log"
	"gibhub.com/raytr/simple-bank/helper/database"
	"gibhub.com/raytr/simple-bank/initialization"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	cfg := config.Init("config_testing", "yml")
	lg := b_log.NewLogger(cfg.Server.Name)
	db, err := database.InitDatabase(&cfg.DBConfig)
	if err != nil {
		panic(err.Error())
	}

	mux = initialization.InitRouting(db, cfg, lg)

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}
