package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/nihon-hitter/be/handler"
	"github.com/yusianglin11010/nihon-hitter/be/transport"
)

func main() {
	sheetID := "1pR0qtVUTvVx8C4O39kTIwFpN06tuqYEXqBi9r-39Mno"

	sheetRepo := transport.NewSheetRepo(sheetID)
	sheetRepo.GetRowNumber()
	h := handler.NewHandler(sheetRepo)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/alive", handler.GetHealth)
	r.GET("/kanji-question", h.GetKanjiQuestion)
	r.GET("/imi-question", h.GetImiQuestion)
	r.Run(":80")

}
