package config

import (
	"github.com/BeGeos/go-echo/pkg/logger"
	"github.com/labstack/echo/v4/middleware"
)

var LoggerConfigs = middleware.RequestLoggerConfig{
	LogStatus:     true,
	LogURI:        true,
	LogMethod:     true,
	LogError:      true,
	HandleError:   true,
	LogValuesFunc: logger.LogValuesFunc, // custom logging function
}
