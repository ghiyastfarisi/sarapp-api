package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	stdResp struct {
		Header stdRespHeader `json:"header"`
		Data   stdRespData   `json:"data"`
	}
	stdRespHeader struct {
		StatusCode   int64  `json:"status_code"`
		ResponseTime string `json:"response_time"`
	}
	stdRespData struct {
		Feeds interface{} `json:"feeds,omitempty"`
	}
	feeds struct {
		ImageURL string `json:"image_url"`
		Caption  string `json:"caption"`
		Location string `json:"location"`
		Date     string `json:"date"`
		User     user   `json:"user"`
	}
	user struct {
		ID           int64  `json:"id"`
		Username     string `json:"username"`
		ProfileImage string `json:"profile_image"`
	}
)

func main() {
	handler()
}

func handler() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerGroupAPI := r.Group("/api")
	{
		routerGroupAPI.GET("/feeds", APIFeedsHandler)
	}

	r.Run(":1992")
}

// APIFeedsHandler ...
func APIFeedsHandler(c *gin.Context) {
	start := time.Now()
	f := []feeds{
		feeds{
			Location: "Jatikramat Garden",
			ImageURL: "http://animalsbirds.com/wp-content/uploads/2016/07/Cat-Animal-Facts-with-Pictures-768x576.jpg",
			Caption:  "Caption event pertama",
			Date:     "01-12-2019",
			User: user{
				ID:           1,
				Username:     "moeghifar",
				ProfileImage: "https://static.pokemonpets.com/images/monsters-images-300-300/1-Bulbasaur.png",
			},
		},
		feeds{
			Location: "Kemang Pratama",
			ImageURL: "http://animalsbirds.com/wp-content/uploads/2016/07/Cat-Animal-Photos-Gallery-1-768x512.jpg",
			Caption:  "Caption event kedua",
			Date:     "01-12-2019",
			User: user{
				ID:           2,
				Username:     "moefaris",
				ProfileImage: "https://static.pokemonpets.com/images/monsters-images-300-300/12-Butterfree.png",
			},
		},
		feeds{
			Location: "Kranggan Permai",
			ImageURL: "http://animalsbirds.com/wp-content/uploads/2016/07/Cats-Animal-768x576.jpg",
			Caption:  "Caption event ketiga",
			Date:     "01-12-2019",
			User: user{
				ID:           3,
				Username:     "sarah",
				ProfileImage: "https://static.pokemonpets.com/images/monsters-images-300-300/25-Pikachu.png",
			},
		},
	}
	r := stdResp{
		Header: stdRespHeader{
			ResponseTime: fmt.Sprintf("%fms", time.Since(start).Seconds()*1000),
			StatusCode:   200,
		},
		Data: stdRespData{
			Feeds: f,
		},
	}
	c.JSON(200, r)
}
