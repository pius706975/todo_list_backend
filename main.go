package main

import (
	"log"
	"os"
	commandline "pius/commandLine"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
)

func init()  {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main()  {
	
	err := commandline.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}