package database

import "fmt"

//Config to maintain DB configuration properties
type Config struct {
	ServerName string "127.0.0.1:3306"
	User       string "root"
	Password   string "11111111"

	DB         string "app"
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}
