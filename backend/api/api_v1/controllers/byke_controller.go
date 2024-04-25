package controllers

import (
	"api/api_v1/configs"
	"api/api_v1/models"
	"api/api_v1/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var bykeCollection *mongo.Collection = configs.GetCollection(configs.DB, "general_info_bikes")

func GetInfoBikes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var bikes []models.BykeGeneralInfo
		defer cancel()

		results, err := bykeCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BikeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var byke models.BykeGeneralInfo

			if err = results.Decode(&byke); err != nil {
				c.JSON(http.StatusInternalServerError, responses.BikeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.BikeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": bikes}})
	}
}
