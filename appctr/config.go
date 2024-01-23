package appctr

import (
	"github.com/spf13/viper"
	"os"
)

const (
	envPrefix    = "TOPO"
	envVar       = "env"
	domainVar    = "domain"
	envUploadDir = "/uploads"
)

const (
	EnvLocal = "localhost"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var cfg viper.Viper
var env string
var domain string

func Cfg() *viper.Viper {
	return &cfg
}

func Env() string {
	return env
}

func Domain() string {
	return domain
}

func prepareCfg() {
	cfg = *viper.New()
	cfg.AutomaticEnv()

	cfg.SetEnvPrefix(envPrefix)
	env = cfg.GetString(envVar)
	domain = cfg.GetString(domainVar)
}

func prepareUpload() bool {
	path, _ := os.Getwd()

	if _, err := os.Stat(path + envUploadDir); os.IsNotExist(err) {
		errDir := os.Mkdir(path+envUploadDir, 0755)
		if errDir != nil {
			return false
		}
	}

	return true
}
