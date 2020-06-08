package model

var Config AppConfig

type AppConfig struct {
	App      App
	Postgres Database
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type App struct {
	Profile string
	Port    int
}
