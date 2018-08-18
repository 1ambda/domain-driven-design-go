package main

import (
	"context"
	"os"

	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/config"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/product"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver"
	"github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/jessevdk/go-flags"
		"github.com/1ambda/domain-driven-design-go/service-gateway/internal/rest"
	"github.com/rs/cors"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/test"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/cart"
)

func main() {
	env := config.Env
	logger := config.GetLogger()

	logger.Infow("Build Manifest",
		"build_date", env.BuildDate,
		"git_commit", env.GitCommit,
		"git_branch", env.GitBranch,
		"git_state", env.GitState,
		"version", env.Version,
		"service_host", env.Host,
		"service_port", env.RestPort,
		"mode", env.Mode,
	)

	swaggerSpec, err := loads.Analyzed(swagserver.FlatSwaggerJSON, "")
	if err != nil {
		logger.Fatalw("Failed to configure REST server", "error", err)
	}
	api := swagapi.NewGatewayAPI(swaggerSpec)
	server := swagserver.NewServer(api)
	server.Port = env.RestPort

	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			logger.Fatalw("Failed to parse command-line option for REST server", "error", err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.Logger = logger.Infof

	// configure session storage
	sessionStore := user.NewSessionStore()

	// configure database
	db := config.GetDatabase(test.MigrateCallback)

	// configure REST server handlers, middlewares
	logger.Info("Configure REST API handlers")

	userRepo := user.NewRepository(db)
	productRepo := product.NewRepository(db)
	cartRepo := cart.NewRepository(db)
	encryptor := user.NewEncryptor(0)

	cartHandler := cart.NewCartHandler(sessionStore, db, cartRepo, userRepo, productRepo)
	cartHandler.Configure(api)

	authHandler := user.NewAuthHandler(userRepo, encryptor, sessionStore)
	authHandler.Configure(api)

	productHandler := product.NewProductHandler(productRepo)
	productHandler.Configure(api)

	logger.Info("Configure REST API middleware")

	handler := api.Serve(nil)
	handler = rest.InjectAuthMiddleware(sessionStore, handler)
	handler = cors.New(cors.Options{
		AllowedOrigins:   env.CorsAllowURLs,
		AllowCredentials: true,
		Debug:            env.EnableDebugCors,
	}).Handler(handler)
	handler = rest.InjectHttpLoggingMiddleware(handler)
	handler = rest.InjectHealthCheckMiddleware(handler)
	server.SetHandler(handler)

	_, cancel := context.WithCancel(context.Background())

	api.ServerShutdown = func() {
		cancel()
	}

	if err := server.Serve(); err != nil {
		logger.Fatalw("Failed to start REST server", "error", err)
	}
}
