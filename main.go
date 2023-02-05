package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/nihon-hitter/be/config"
	"github.com/yusianglin11010/nihon-hitter/be/handler"
	repo "github.com/yusianglin11010/nihon-hitter/be/repository"
	"github.com/yusianglin11010/nihon-hitter/be/transport"
)

func main() {

	// sheetID := "1pR0qtVUTvVx8C4O39kTIwFpN06tuqYEXqBi9r-39Mno"

	client := config.NewClient()
	config := config.NewConfig()
	srv := repo.NewSheetService(client)
	sheetRepo := transport.NewSheetRepo(config, srv)

	h := handler.NewHandler(sheetRepo)

	s := gin.Default()
	s.Use(cors.Default())

	// server.Use(cors.Default())
	s.GET("/alive", h.GetHealth)
	s.GET("/kanji-question", h.GetKanjiQuestion)
	s.GET("/imi-question", h.GetImiQuestion)
	s.Run(config.Port)

}
