package api

import (
	"context"
	
	"marketplace/pkg/common/database/sqlc"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DBQuerier interface {
    GetGoodsByTitle(ctx context.Context, title string) ([]sqlc.Good, error)
}


type MockDB struct {
    mock.Mock
}

func (m *MockDB) GetGoodsByTitle(ctx context.Context, title string) ([]sqlc.Good, error) {
    args := m.Called(ctx, title)
    return args.Get(0).([]sqlc.Good), args.Error(1)
}




func TestSearchGoodsByTitle(t *testing.T) {
 
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer mr.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockDB:= new(MockDB)
	
	h := handler{
	RedisDB: rdb,
	DB:    mockDB,
	Validator: validator.New(),
	}


	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/?text=someGoods", nil)

	
	expectedGoods := []sqlc.Good{{Title: "Good1"}, {Title: "Good2"}}
	mockDB.On("GetGoodsByTitle", c, "someGoods").Return(expectedGoods, nil)

	
	h.SearchGoodsByTitle(c)


	assert.Equal(t, http.StatusOK, w.Code)
	

	mockDB.AssertCalled(t, "GetGoodsByTitle", c, "someGoods")
}
