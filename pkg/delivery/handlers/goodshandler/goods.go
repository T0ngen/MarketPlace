package goodshandler

import (
	"context"
	"encoding/json"
	"fmt"
	"marketplace/pkg/database/sqlc"
	"marketplace/pkg/delivery"

	"net/http"
	"net/url"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)



type GoodsHandler interface{
	GetGoodsByTitle(ctx context.Context, title string) ([]sqlc.Good, error) 
}

func checkUnkownParams(query url.Values) error{
	for key := range query{
		if key != "text"{
			return fmt.Errorf("unkown query parametr")
		}
	}
	return nil
}

//TODO РАЗОБРАТЬ ХЭНДЛЕР
func NewGoodsHandler(gH GoodsHandler, logger *zap.SugaredLogger, rdb redis.Client)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		query := r.URL.Query()
		err := checkUnkownParams(query)
		if err != nil{
			errText := `{"message": "bad params in query"}`
			delivery.WriteResponse(logger, w, []byte(errText), http.StatusBadRequest)
			return
		}
		text := query.Get("text")
		ctx := context.Background()
		val, err:= rdb.Get(ctx, text).Result()
		if err == redis.Nil{
			goods, err :=gH.GetGoodsByTitle(r.Context(), text)
		
			if err != nil{
				errText := `{"message": "nothing found"}`
				delivery.WriteResponse(logger, w, []byte(errText), http.StatusBadRequest)
				return
			}
			if len(goods) ==0{
				errText := `{"message": "nothing found"}`
				delivery.WriteResponse(logger, w, []byte(errText), http.StatusBadRequest)
				return
			}
			goodsJSON, err := json.Marshal(goods)
			if err != nil{
				errText := fmt.Sprintf(`{"message": "error in coding films: %s"}`, err)
				delivery.WriteResponse(logger, w, []byte(errText), http.StatusInternalServerError)
				return
			}
			
			

			
			err = rdb.Set(ctx, text, goodsJSON, 10*time.Second).Err()
			if err != nil{
				logger.Errorf("error with redis")
			}
			delivery.WriteResponse(logger, w, goodsJSON, http.StatusOK)
		}else{
			delivery.WriteResponse(logger, w, []byte(val), http.StatusOK)
		}
		
		
			
		
		
		
		
	}}
