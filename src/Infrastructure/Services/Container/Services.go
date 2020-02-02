package container

import (
	// "fmt"
	"log"
	"github.com/sarulabs/di"

	// Logger
	domainlogger "github.com/gnemes/golang-ws-skeleton/Domain/Services/Logger"
	"github.com/gnemes/golang-ws-skeleton/Infrastructure/Services/Logger"

	"github.com/gnemes/golang-ws-skeleton/Infrastructure/Services/Config"
	//"github.com/gnemes/golang-ws-skeleton/Infrastructure/Services/Database"
	
)

// Services contains the definitions of the application services.
var Services = []di.Def{
	{
		Name:  "Logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			// Logger configuration
			loggerConfig := domainlogger.LoggerConfiguration{
				EnableConsole:     config.Get().LoggerConsoleEnable,
				ConsoleLevel:      config.Get().LoggerConsoleLevel,
				ConsoleJSONFormat: true,
				EnableFile:        config.Get().LoggerFileEnable,
				FileLevel:         config.Get().LoggerFileLevel,
				FileJSONFormat:    true,
				FileLocation:      config.Get().LoggerFileLocation,
			}

			// Create logger instance
			l, err := logger.NewLogger(loggerConfig, logger.InstanceZapLogger)
			if err != nil {
				log.Fatalf("Could not instantiate log %s", err.Error())
			}
			return l.(domainlogger.Logger), nil
		},
	},
	/*
	{
		Name:  "MySQLClient",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {

			mysqlConfig := database.MySQLConfiguration{
				Host:     config.Get().MySQLHost,
				Port:     config.Get().MySQLPort,
				Username: config.Get().MySQLUser,
				Password: config.Get().MySQLPassword,
				Database: config.Get().MySQLDatabase,
			}

			mysqlInstance, errDb := database.New(mysqlConfig)
			if errDb != nil {
				logger.Fatalf("Could not instantiate database %s", errDb.Error())
			}
			return mysqlInstance, nil
		},
	},
	*/
}
