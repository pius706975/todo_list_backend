package main

import (
	"log"
	"os"
	commandline "pius/commandLine"
	_ "github.com/joho/godotenv/autoload"
	"github.com/asaskevich/govalidator"
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