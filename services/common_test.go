package services

import (
	"fmt"
	"gibhub.com/raytr/simple-bank/helper/b_log"
	"strconv"
)

var logger b_log.Logger

func init() {
	{
		lg := b_log.NewLogger("testing")
		lg.Info(fmt.Sprintf("Listening at port %s", strconv.Itoa(8082)))
	}
}
