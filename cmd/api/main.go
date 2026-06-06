package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RishavSinha20/cricstream/internal/redisstore"
	"time"

    "github.com/gin-contrib/cors"
)

func main() {

	redisstore.Init()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5500",
		},
	
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
	
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
		},
	
		MaxAge: 12 * time.Hour,
	}))

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