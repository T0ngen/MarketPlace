package goodshandler

import (
	"context"
	"encoding/json"
	"fmt"
	"marketplace/database/sqlc"
	"marketplace/delivery"
	"net/http"
	"net/url"

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


func NewGoodsHandler(gH GoodsHandler, logger *zap.SugaredLogger)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		query := r.URL.Query()
		err := checkUnkownParams(query)
		if err != nil{
			errText := `{"message": "bad params in query"}`
			delivery.WriteResponse(logger, w, []byte(errText), http.StatusBadRequest)
			return
		}
		text := query.Get("text")
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
		
		

		delivery.WriteResponse(logger, w, goodsJSON, http.StatusOK)
		
		
		
	}}
