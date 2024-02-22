package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	// yaml:"env" - название конфига в yaml файле
	// env-default:"local" - значение по умолчанию, если не указано
	Env string `yaml:"env" env-default:"local"`
	// env-required:"true" - должно быть обязательно указано
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// Must - предупреждаем, что пункция будет паниковать при ошибке
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	// проверяем обязательно, что файл реально существует
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// --config="path/to/file.yaml"
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		// Если не передали флаг при запуске, то читаем переменную из конфига
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
