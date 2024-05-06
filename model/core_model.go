package model

type Config struct {
	Port               int    `yaml:"port"`
	DBConnectionString string `yaml:"db_connection_string"`
	DBName             string `yaml:"db_name"`
}
