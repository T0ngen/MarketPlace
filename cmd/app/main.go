package main

import (
	"log"
	"marketplace/database/sqlc"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)



func main(){
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Printf("error in logger start")
		return
	}
	logger := zapLogger.Sugar()
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Printf("error in logger sync")
		}
	}()

	db, err := sqlc.OpenPsgtreConnection()
	if err != nil {
		logger.Errorf("error in connection to mysql: %s", err)
		return
	}
	_ = db
	router := mux.NewRouter()
}