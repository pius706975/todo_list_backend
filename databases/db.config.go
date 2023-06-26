package databases

import "os"

type Configuration struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func GetConfig() Configuration {
	
	conf := Configuration{
		DB_HOST: os.Getenv("MYSQL_HOST"),
		DB_PORT: os.Getenv("MYSQL_PORT"),
		DB_USER: os.Getenv("MYSQL_USER"),
		DB_PASS: os.Getenv("MYSQL_PASSWORD"),
		DB_NAME: os.Getenv("MYSQL_DBNAME"),
	}

	return conf
}
