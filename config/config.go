package config

import (
	"fmt"
	mdwModel "gb-ui-core/internal/pkg/middleware/model"
	dconfig "gb-ui-core/pkg/damqp/config"
	"gb-ui-core/pkg/tconfig"
	thttpConfig "gb-ui-core/pkg/thttp/config"
	tloggerConfig "gb-ui-core/pkg/tlogger/config"
	tsecureConfig "gb-ui-core/pkg/tsecure/config"
	tstorageConfig "gb-ui-core/pkg/tstorage/config"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	BaseConfig       tconfig.BaseConfig
	HttpConfig       thttpConfig.ThttpConfig
	LoggerConfig     tloggerConfig.TLoggerConfig
	SecureConfig     tsecureConfig.TSecureConfig
	StorageConfig    tstorageConfig.TStorageConfig
	AmqpConfig       dconfig.BrokerConfig
	MiddlewareConfig mdwModel.MiddlewareConfig
}

func NewConfig() *Config {
	v, err := loadConfig()
	if err != nil {
		panic(fmt.Errorf("can't parse config: %s", err.Error()))
	}
	config, err := parseConfig(v)
	if err != nil {
		panic(fmt.Errorf("can't parse config: %s", err.Error()))
	}

	return config
}

func BuildConfig(ctn di.Container) (interface{}, error) {
	return NewConfig(), nil
}

func loadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
