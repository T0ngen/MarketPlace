package goodshandler

import (
	"context"
	"fmt"
	"marketplace/database/sqlc"
	"net/http"
)



type GoodsHandler interface{
	GetGoodsByTitle(ctx context.Context, title string) (sqlc.Good, error)
}


func NewGoodsHandler(gH GoodsHandler)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		good, err :=gH.GetGoodsByTitle(r.Context(), "Product9")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(good)
	}}
