package testing

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/helper/database"
	"gibhub.com/raytr/simple-bank/initialization"
	"github.com/go-kit/kit/log"
)

var mux *http.ServeMux
var logger log.Logger

func TestMain(m *testing.M) {
	cfg := config.Init("config_testing", "yml")

	db, err := database.InitDatabase(&cfg.DBConfig)
	if err != nil {
		panic(err.Error())
	}

	mux = initialization.InitRouting(db, cfg, logger)
	// http.Handle("/", accessControl(mux))

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}
