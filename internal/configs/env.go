package configs

import (
	"os"
)

type EnvVar int

const (
	EnvVarPort EnvVar = iota
	EnvVarReverseProxyAppConfigFile
	EnvVarRoutesConfigsMasterFile
	EnvVarReverseProxyLogLevel
)

var envVarName = map[EnvVar]string{
	EnvVarPort: "PORT",
	EnvVarReverseProxyAppConfigFile: "REVERSE_PROXY_APP_CONFIG_FILE",
	EnvVarRoutesConfigsMasterFile: "ROUTES_CONFIGS_MASTER_FILE",
	EnvVarReverseProxyLogLevel: "REVERSE_PROXY_LOG_LEVEL",
}

func (ev EnvVar) String() string {
	return envVarName[ev]
}

var envVars = map[string]string{
	EnvVarPort.String(): "8080",
	EnvVarReverseProxyAppConfigFile.String(): "app.configs.json",
	EnvVarRoutesConfigsMasterFile.String(): "routes.master.json",
	EnvVarReverseProxyLogLevel.String(): "REVERSE_PROXY_LOG_LEVEL",
}

func initEnvVars() {
	for _, evName := range envVarName {
		if evVal := os.Getenv(evName); len(evVal) > 0 {
			envVars[evName] = evVal
		}
	}
}
