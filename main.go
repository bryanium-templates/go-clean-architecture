package main

import (
	"flag"
	"log"
	"os"

	"github.com/bryanium-templates/go-clean-architecture/infrastructure"
)

func main() {
	infrastructure.LoadEnvVariables()

	infrastructure.ConnectDb()

	migrateNow := flag.Bool("migrate", false, "run DB migration & exit")
	flag.Parse()

	if *migrateNow {
		infrastructure.SyncDatabase()
		log.Println("Migration complete... exiting...")
		os.Exit(0)
	}

	infrastructure.RunGin()
}