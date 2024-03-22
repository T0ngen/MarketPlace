package main

import (
	"marketplace/pkg/database/redis"
	"marketplace/pkg/database/sqlc"
	// "marketplace/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)



func main(){
	// logger.InitLogger()
	logrus.Info("The logger is connected to the server!")
	
	db, err := sqlc.OpenPsgtreConnection()
	if err != nil {
		logrus.Fatalf("error in connection to mysql: %s", err)
		return
	}
	
	redisCon := redis.OpenRedis()
	
	if err != nil{
		logrus.Fatalf("error in connection to redis: %s", err)
	}
	
	_= redisCon
	logrus.Info("The Redis is connected to the server!")


	client := sqlc.New(db)
	_ = client
	logrus.Info("The Postgres is connected to the server!")
	// newGoodsHanlder := goodshandler.NewGoodsHandler(queries, logger, *redisCon)


	router := mux.NewRouter()
	// router.HandleFunc("/search", newGoodsHanlder).Methods("GET")
	addr := ":8080"
	logrus.Info("starting server",
		"type", "START",
		"addr", addr,
	)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		logrus.Fatalf("errror in server start")
	}
}