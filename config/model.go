package config

type Config struct {
	Server     Server     `yaml:"server"`
	Connection Connection `yaml:"connection"`
	Log        Log        `yaml:"log"`
}

type Server struct {
	Port    string `yaml:"port"`
	XApiKey string `yaml:"x-api-key"`
}

type Connection struct {
	MySQL MySQL `yaml:"mysql"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type Log struct {
	Level string `yaml:"level"`
}
