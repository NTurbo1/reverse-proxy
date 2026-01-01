package server

import (
	"fmt"
	"net/http"

	"github.com/nturbo1/apigtw/internal/configs"
	"github.com/nturbo1/apigtw/internal/middleware"
	"github.com/nturbo1/apigtw/internal/routing"
	"github.com/nturbo1/apigtw/internal/log"
)

func NewServer(
	appConfigs *configs.AppConfigs, env *configs.Environment,
) (*http.Server, error) {

	mux := http.NewServeMux()
	serverHandler := NewServerHandler(mux)
	log.Debug("Setting up the routes...")
	err := routing.SetUpRouteHandlers(appConfigs, env, mux)
	if err != nil {
		log.Error("Failed to set up route handlers due to: %s", err)
		return nil, err
	}

	return &http.Server{
		Addr:           fmt.Sprintf(":%d", appConfigs.Server.Port),
		Handler:        serverHandler,
		ReadTimeout:    appConfigs.Server.Timeout,
		WriteTimeout:   appConfigs.Server.Timeout,
		MaxHeaderBytes: 1 << 20,
	}, nil
}

func NewServerHandler(mux *http.ServeMux) http.Handler {

	return middleware.PrependMiddlewareChain(
		mux,
		middleware.RateLimitMiddleware,
		middleware.LogMiddleware,
		middleware.AuthMiddleware,
	)
}
