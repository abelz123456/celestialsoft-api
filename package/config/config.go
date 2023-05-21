package config

import (
	"fmt"
	"strings"

	"github.com/abelz123456/celestial-api/package/log"
	"github.com/spf13/viper"
)

func Init(path string) Config {
	logger := log.NewLog()
	var cfg Config

	fmt.Println("PATH", path)
	viper.AddConfigPath(path)
	viper.SetConfigName("base")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigName("test")
		logger.PanicOnError(viper.ReadInConfig(), "", nil)
	}

	logger.PanicOnError(viper.Unmarshal(&cfg), "", nil)
	cfg.BasePath = path

	optimizeConfig(&cfg)
	return cfg
}

func optimizeConfig(cfg *Config) {
	if cfg.AppEnv == "" {
		cfg.AppEnv = "development"
	}

	if cfg.DevelopmentPort == "" || cfg.AppEnv != "development" {
		cfg.DevelopmentPort = "3000"
	}

	if cfg.StaticFilePath == "" {
		cfg.StaticFilePath = "public"
	}

	if cfg.AppScheme == "" || (cfg.AppScheme != "http" && cfg.AppScheme != "https") {
		cfg.AppScheme = "http"
	}

	if !strings.HasPrefix(cfg.DevelopmentPort, ":") {
		cfg.DevelopmentPort = fmt.Sprintf(":%s", cfg.DevelopmentPort)
	}

	for i, host := range cfg.TrustedProxies {
		cfg.TrustedProxies[i] = strings.TrimSpace(host)
	}
}
