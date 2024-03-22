package sqlc

import (
	"context"
	
	"testing"

	"github.com/stretchr/testify/require"
)




func TestGetGoodsByTitle(t *testing.T){


	queries := New(db)

	ctx := context.Background()

	title := "Product1"

	goods, err := queries.GetGoodsByTitle(ctx, title)
	require.NotEmpty(t, goods, "Returns empty struct")
	require.NoErrorf(t, err, "Error with get goods by title %v", err)

	

}


func TestCreateGoods(t *testing.T){

	queries := New(db)

	ctx := context.Background()


	agrs := CreateGoodsParams{
		SellerID: 1,
		Title: "new3",
		Price: 3000,
		Description: "good item",
		Image: "img.png",
		Category: "Category2",
		Rating: "4.8",
		Discount: 10,
		Status: "Active",

	}

	item, err := queries.CreateGoods(ctx, agrs)
	require.NotEmpty(t, item, "Returns empty struct")
	require.NoErrorf(t, err, "Error with create goods %v", err)


}