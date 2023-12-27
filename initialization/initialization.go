package initialization

import (
	"gibhub.com/raytr/simple-bank/helper/b_log"
	"net/http"

	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/helper/token"
	"gibhub.com/raytr/simple-bank/middleware"
	"gibhub.com/raytr/simple-bank/repository"
	"gibhub.com/raytr/simple-bank/services"
	"gibhub.com/raytr/simple-bank/transport"
	"gorm.io/gorm"
)

const (
	TokenTypeJwt    = "jwt"
	TokenTypePesato = "pesato"
)

func InitRouting(db *gorm.DB, cfg *config.Config, logger b_log.Logger) *http.ServeMux {
	var tokenMaker token.Maker
	var err error

	if cfg.SecConfig.Type == TokenTypePesato {
		tokenMaker, err = token.NewPasetoMaker(cfg.SecConfig.PasetoTokenSymmetricKey)
		if err != nil {
			panic(err)
		}
	} else {
		tokenMaker = token.NewJWTMaker(cfg.SecConfig.PasswordPepper)
	}

	baseRepo := repository.NewBaseRepository(db)
	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	authSvc := services.NewAuthService(baseRepo, userRepo, sessionRepo, cfg.SecConfig, tokenMaker, logger)
	userSvc := services.NewUserService(userRepo, cfg.SecConfig, logger)

	authHttp := transport.AuthHttpHandler(authSvc, logger)
	userHttp := transport.UserHttpHandler(userSvc, logger)

	auth := middleware.Auth{
		Pepper:     cfg.SecConfig.PasswordPepper,
		Skip:       cfg.SecConfig.AuthSkip,
		TokenMaker: tokenMaker,
	}

	server := http.NewServeMux()
	server.Handle("/user/register", userHttp)
	server.Handle("/user/", auth.Authenticate(userHttp))
	server.Handle("/users", auth.Authenticate(userHttp))
	server.Handle("/login", authHttp)
	server.Handle("/renew-access-token", authHttp)

	return server
}
