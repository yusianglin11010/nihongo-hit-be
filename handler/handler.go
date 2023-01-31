package handler

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/nihon-hitter/be/transport"
)

type sheetHandler struct {
	sheetRepo *transport.SheetRepo
}

func NewHandler(sheetRepo *transport.SheetRepo) sheetHandler {
	return sheetHandler{
		sheetRepo: sheetRepo,
	}
}

func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h sheetHandler) GetKanjiQuestion(c *gin.Context) {
	question := map[string]string{}

	maxRowNum := h.sheetRepo.GetRowNumber()
	readStart := rand.Intn(maxRowNum-50) + 1
	kanaList, _ := h.sheetRepo.GetKana(readStart, readStart+50)
	kanjiList, _ := h.sheetRepo.GetKanji(readStart, readStart+50)
	for idx, value := range kanjiList {
		if value == "null" {
			continue
		}
		question[value] = kanaList[idx]
	}
	c.JSON(200, gin.H{
		"question": question,
	})

}

func (h sheetHandler) GetImiQuestion(c *gin.Context) {
	question := map[string]string{}
	maxRowNum := h.sheetRepo.GetRowNumber()
	readStart := rand.Intn(maxRowNum-50) + 1
	kanjiList, _ := h.sheetRepo.GetKanji(readStart, readStart+50)
	imiList, _ := h.sheetRepo.GetImi(readStart, readStart+50)
	for idx, value := range kanjiList {
		if value == "null" {
			continue
		}
		question[value] = imiList[idx]
	}

	c.JSON(200, gin.H{
		"question": question,
	})
}
