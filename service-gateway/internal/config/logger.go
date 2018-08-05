package config

import "go.uber.org/zap"

func GetLogger() *zap.SugaredLogger {
	var log *zap.Logger = nil
	var conf zap.Config

	if Env.IsTestMode() || Env.IsLocalMode() {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	if Env.isDebugLogLevel() {
		conf.Level.SetLevel(zap.DebugLevel)
	} else {
		conf.Level.SetLevel(zap.InfoLevel)
	}

	log, _ = conf.Build()
	logger := log.Sugar().With("service_name", Env.ServiceName, "service_id", Env.ServiceId)

	return logger
}
