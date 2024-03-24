package main

import (
	"fmt"
	"marketplace/pkg/api"
	"marketplace/pkg/common/config"
	"marketplace/pkg/common/database/redis"
	"marketplace/pkg/common/database/sqlc"

	// "marketplace/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)



func main(){
	// logger.InitLogger()
	logrus.Info("The logger is connected to the server!")
	

	var conf config.Config

	db, err := sqlc.OpenPostgresConnection(conf)
	fmt.Println(err)
	if err != nil {
		logrus.Fatalf("error in connection to psql: %s", err)
		return
	}
	
	redisClient := redis.OpenRedis(conf)
	logrus.Info("The Redis is connected to the server!")

	validate := validator.New()
	

	psgreClient := sqlc.New(db)
	logrus.Info("The Postgres is connected to the server!")
	// newGoodsHanlder := goodshandler.NewGoodsHandler(queries, logger, *redisCon)

	r := gin.Default()

	
	
	api.RegisterRouter(r,validate, redisClient, psgreClient)

	// router.HandleFunc("/search", newGoodsHanlder).Methods("GET")
	addr := ":8080"
	logrus.Info("starting server",
		"type", "START",
		"addr", addr,
	)

	err = r.Run()
	
	if err != nil {
		logrus.Fatal("Can't start the server on the port: 8080")
		return
	}

	logrus.Info("The server is up!")
}