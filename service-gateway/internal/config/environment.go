package config

import (
	"github.com/kelseyhightower/envconfig"
)

var (
	// These fields are populated by govvv
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
)

type Environment struct {
	// general
	Host     string `envconfig:"HOST" default:"localhost"` // parsed by go-swagger
	RestPort int    `envconfig:"PORT" default:"30001"`     // parsed by go-swagger

	Mode        string `envconfig:"SERVICE_MODE" default:"LOCAL"` // `LOCAL`, `TEST`, `DEV`, `PROD`
	ServiceName string `envconfig:"SERVICE_NAME" default:"service-gateway"`
	ServiceId   string `envconfig:"SERVICE_ID" default:"0"`

	// deployment
	NodeName string `envconfig:"NODE_NAME" default:""`
	PodName string `envconfig:"POD_NAME" default:""`
	PodNamespace string `envconfig:"POD_NAMESPACE" default:""`
	PodIP string `envconfig:"POD_IP" default:""`

	// storage
	MysqlHost     string `envconfig:"MYSQL_HOST" default:"localhost"`
	MysqlPort     string `envconfig:"MYSQL_PORT" default:"3306"`
	MysqlDatabase string `envconfig:"MYSQL_DATABASE" default:"application"`
	MysqlUserName string `envconfig:"MYSQL_USERNAME" default:"root"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD" default:"root"`
	SchemaAssetDir string `envconfig:"SCHEMA_ASSET_DIR" default:"../../asset/sql"`

	// server
	CorsAllowURLs []string `envconfig:"GATEWAY_COR_URLS" default:"http://localhost:8080,http://127.0.0.1:8080,http://0.0.0.0:8080"`

	// debugging
	EnableDebugSQL      bool   `envconfig:"ENABLE_DEBUG_SQL" default:"true"`
	EnableDebugHTTP     bool   `envconfig:"ENABLE_DEBUG_HTTP" default:"true"`
	EnableSwaggerUI     bool   `envconfig:"ENABLE_SWAGGER_UI" default:"true"`
	EnableDebugCors     bool   `envconfig:"ENABLE_DEBUG_CORS" default:"false"`
	DisableSessionCheck bool   `envconfig:"DISABLE_SESSION_CHECK" default:"false"`
	LogLevel            string `envconfig:"LOG_LEVEL" default:"INFO"` // `DEBUG`, `INFO`

	// copied from govvv injected values
	BuildDate string
	GitCommit string
	GitBranch string
	GitState  string
	Version   string
}

func (e *Environment) DebugSQLEnabled() bool {
	return e.EnableDebugSQL
}

func (e *Environment) DebugHTTPEnabled() bool {
	return e.EnableDebugHTTP
}

func (e *Environment) SwaggerUIEnabled() bool {
	return e.EnableSwaggerUI
}

func (e *Environment) SessionCheckDisabled() bool {
	return e.DisableSessionCheck
}

func (e *Environment) IsInfoLogLevel() bool {
	return e.LogLevel == "INFO"
}

func (e *Environment) isDebugLogLevel() bool {
	return e.LogLevel == "DEBUG"
}

func (e *Environment) IsTestMode() bool {
	return e.Mode == "TEST"
}

func (e *Environment) IsLocalMode() bool {
	return e.Mode == "LOCAL"
}

var Env Environment

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic("Failed to get specification")
	}

	Env.BuildDate = BuildDate
	Env.GitCommit = GitCommit
	Env.GitBranch = GitBranch
	Env.GitState = GitState
	Env.GitState = GitState
	Env.Version = Version
}
