package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RishavSinha20/cricstream/internal/redisstore"
)

func main() {

	redisstore.Init()

	router := gin.Default()

	router.GET(
		"/matches/:id",
		func(c *gin.Context) {

			matchID := c.Param("id")

			result, err := redisstore.Client.Get(
				redisstore.Ctx,
				matchID,
			).Result()

			if err != nil {

				c.JSON(
					http.StatusNotFound,
					gin.H{
						"error": "match not found",
					},
				)

				return
			}

			c.Data(
				http.StatusOK,
				"application/json",
				[]byte(result),
			)
		},
	)

	router.Run(":8080")
}