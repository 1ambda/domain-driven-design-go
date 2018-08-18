// This file is safe to edit. Once it exists it will not be overwritten

package swagserver

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi/cart"
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi/product"
)

//go:generate swagger generate server --target ../pkg/generated/swagger --name  --spec ../../schema-swagger/gateway-rest.yaml --api-package swagapi --model-package swagmodel --server-package swagserver --exclude-main

func configureFlags(api *swagapi.GatewayAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *swagapi.GatewayAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AuthLoginHandler = auth.LoginHandlerFunc(func(params auth.LoginParams) middleware.Responder {
		return middleware.NotImplemented("operation auth.Login has not yet been implemented")
	})
	api.AuthLogoutHandler = auth.LogoutHandlerFunc(func(params auth.LogoutParams) middleware.Responder {
		return middleware.NotImplemented("operation auth.Logout has not yet been implemented")
	})
	api.AuthRegisterHandler = auth.RegisterHandlerFunc(func(params auth.RegisterParams) middleware.Responder {
		return middleware.NotImplemented("operation auth.Register has not yet been implemented")
	})
	api.AuthWhoamiHandler = auth.WhoamiHandlerFunc(func(params auth.WhoamiParams) middleware.Responder {
		return middleware.NotImplemented("operation auth.Whoami has not yet been implemented")
	})
	api.ProductFindAllHandler = product.FindAllHandlerFunc(func(params product.FindAllParams) middleware.Responder {
		return middleware.NotImplemented("operation product.FindAll has not yet been implemented")
	})
	api.ProductFindOneWithOptionsHandler = product.FindOneWithOptionsHandlerFunc(func(params product.FindOneWithOptionsParams) middleware.Responder {
		return middleware.NotImplemented("operation product.FindOneWithOptions has not yet been implemented")
	})
	api.CartGetUserCartHandler = cart.GetUserCartHandlerFunc(func(params cart.GetUserCartParams) middleware.Responder {
		return middleware.NotImplemented("operation cart.GetUserCart has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
