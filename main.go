package main

import (
	"github.com/bryanium-templates/go-clean-architecture/infrastructure"
)

func init() {
	infrastructure.LoadEnvVariables()
	infrastructure.ConnectDb()
	infrastructure.SyncDatabase()
}


func main(){
	infrastructure.RunGin()
}