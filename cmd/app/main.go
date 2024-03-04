package main

import (
	"log"
	"marketplace/database/sqlc"
	"marketplace/delivery/handlers/goodshandler"
	"net/http"

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
	
	queries := sqlc.New(db)

	router := mux.NewRouter()
	router.HandleFunc("/users", goodshandler.NewGoodsHandler(queries)).Methods("GET")
	addr := ":8080"
	logger.Infow("starting server",
		"type", "START",
		"addr", addr,
	)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		logger.Fatalf("errror in server start")
	}
}