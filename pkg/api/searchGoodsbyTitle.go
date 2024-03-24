package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"marketplace/pkg/api/responses"
	// "marketplace/pkg/common/database/sqlc"

	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)





func genLog(err error, funcName, file string ) logrus.Fields {
  
	return logrus.Fields{
	  "Error":            err,
	  "File":             file,
	  "FromFunction":     funcName,
	}
  
  }

func (h *handler) SearchGoodsByTitle(c *gin.Context){
	text := c.Query("text")
	if text == "" {
		err := errors.New("missing 'text' parameter")
		logrus.WithFields(genLog(err, "SearchGoodsByTitle",
		"searchGoodsByTitle")).Errorf("missing 'text' parameter")
        c.JSON(http.StatusBadRequest, gin.H{
				"error": responses.ErrorResponse{
					Error: "missing 'text' parameter",
					Description: responses.ErrorDescription{
						ErrorCode:           http.StatusBadRequest,
						TagError:            "query_error",
						DetailedDescription: "missing 'text' parameter, try again",
					},
				},
			})
			return
    }
	
	val, err:= h.RedisDB.Get(c, text).Result()
	if err != redis.Nil{
		logrus.WithFields(genLog(err, "SearchGoodsByTitle",
		"searchGoodsByTitle")).Infof("successfully found in redis")
		responses.WriteResponse(c.Writer, []byte(val), http.StatusOK)
		return
		
	}
	
	

	goods, err :=h.DB.GetGoodsByTitle(c, text)
	
	if err != nil{
		logrus.WithFields(genLog(err, "SearchGoodsByTitle",
		"searchGoodsByTitle")).Errorf("error while GetGoodsByTitle ")
        c.JSON(http.StatusInternalServerError, gin.H{
				"error": responses.ErrorResponse{
					Error: "nothing found",
					Description: responses.ErrorDescription{
						ErrorCode:           http.StatusInternalServerError,
						TagError:            "inner_error",
						DetailedDescription: "error while searching goods by title, try again",
					},
				},
			})
			return
	}

	if len(goods) ==0{
		
		err := errors.New("nothing found")
		logrus.WithFields(genLog(err, "SearchGoodsByTitle",
		"searchGoodsByTitle")).Errorf("Nothing found while searching by title")
        c.JSON(http.StatusBadRequest, gin.H{
				"error": responses.ErrorResponse{
					Error: "Nothing found while searching by title",
					Description: responses.ErrorDescription{
						ErrorCode:           http.StatusBadRequest,
						TagError:            "query_error",
						DetailedDescription: "Nothing found, try with another request",
					},
				},
			})
			return
		
	}
	

   
	goodsJSON, err := json.Marshal(goods)
	if err != nil{
		logrus.WithFields(genLog(err, "SearchGoodsByTitle",
		"searchGoodsByTitle")).Errorf("error while marshal json")
        c.JSON(http.StatusInternalServerError, gin.H{
				"error": responses.ErrorResponse{
					Error: "error while data processing with json",
					Description: responses.ErrorDescription{
						ErrorCode:           http.StatusInternalServerError,
						TagError:            "inner_error",
						DetailedDescription: "error while data processing with json, try again",
					},
				},
			})
			return
	}
	err = h.RedisDB.Set(c, text, goodsJSON, 10*time.Second).Err()
	if err != nil{
		fmt.Printf("error with redis %v", err)
		return
	}
	logrus.WithFields(genLog(err, "SearchGoodsByTitle",
		"searchGoodsByTitle")).Infof("successfully found in DB")
	
	responses.WriteResponse(c.Writer, []byte(goodsJSON), http.StatusOK)
	

}
