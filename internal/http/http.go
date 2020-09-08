package http

import (
	"go-scaffold/internal/http/handler"
	"go-scaffold/internal/http/middleware"
	"go-scaffold/pkg/conf"
	"go-scaffold/pkg/xgin"
)

// Serve ..
func Serve() {
	xgin.RawConfig("app.http", conf.GetHandler()).Setup(
		// middleware
		middleware.UseLogger,
		middleware.UseRecovery,
		middleware.UseAuthentication,
		middleware.UseSwagger,

		// register handler
		handler.RoutePing,
		handler.RouteAuthentication,
	).Build().ListenAndServe()
}
