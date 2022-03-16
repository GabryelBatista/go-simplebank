package config


type Database struct {
	Username string
	Password string
	Database string
	Host string
}

func MongoDB () Database {
	mongo := Database{Username: "root", Password: "", Database: "simple-bank", Host: "mongodb://127.0.0.1:28017/"}
	return mongo;
}