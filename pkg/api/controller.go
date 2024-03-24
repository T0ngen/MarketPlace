package api

import (
	"context"
	"marketplace/pkg/common/database/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
)



type GoodsInterface interface{
	GetGoodsByTitle(ctx context.Context, title string) ([]sqlc.Good, error) 
}


type handler struct{
	DB GoodsInterface
	RedisDB *redis.Client
	Validator *validator.Validate
	

}





func RegisterRouter(r *gin.Engine, validate *validator.Validate, redisDB *redis.Client, db *sqlc.Queries ){
	

	h := &handler{DB: db, Validator: validate, RedisDB: redisDB}

	routes := r.Group("/api/v1/")
	
	routes.GET("/search", h.SearchGoodsByTitle)

}