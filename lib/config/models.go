package config

type ConfigMap struct {
	App    AppConfig   `yaml:"app"`
	Kafka  KafkaConfig `yaml:"kafka"`
	MainDB DBConfig    `yaml:"main-db"`
}

type AppConfig struct {
	Mode string `yaml:"mode"`
}

type KafkaConfig struct {
	Host    string `yaml:"host"`
	GroupID string `yaml:"groupID"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"name"`
}
