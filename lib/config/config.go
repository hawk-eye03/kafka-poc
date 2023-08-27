package config

import (
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func LoadConfig() *ConfigMap {
	// Open the YAML file
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		zap.L().Info("Error config path variable not found resetting path to default path")
		path = "resources/config.yaml"
	}
	zap.L().Info("Path Environment variable value:", zap.String("Path", path))

	file, err := os.Open(path)
	if err != nil {
		zap.L().Info("Error cannot open file at given path")
		return nil
	}

	defer file.Close()

	// Create a YAML decoder
	decoder := yaml.NewDecoder(file)

	// Create a struct to hold the parsed data
	var config ConfigMap

	// Decode the YAML into the struct
	err = decoder.Decode(&config)
	if err != nil {
		zap.L().Info("Error cannot decode yaml file")
		return nil
	}

	// Print the parsed data
	zap.L().Info("Kafka hosted on:", zap.String("host", config.Kafka.Host))
	// zap.L().Info("Kafka hosted on:", zap.String("host", config.Kafka.Host))
	return &config
}

func (db *DBConfig) ValidateDBCreds() bool {
	return (db.DBName != "" && db.Host != "" && db.Password != "" && db.Port != "" && db.Username != "")
}
