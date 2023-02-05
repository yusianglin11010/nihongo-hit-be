package handler

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/nihon-hitter/be/transport"
)

type SheetHandler struct {
	sheetRepo *transport.SheetRepo
}

func NewHandler(sheetRepo *transport.SheetRepo) *SheetHandler {
	return &SheetHandler{
		sheetRepo: sheetRepo,
	}
}

func (h *SheetHandler) GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *SheetHandler) GetKanjiQuestion(c *gin.Context) {
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

func (h *SheetHandler) GetImiQuestion(c *gin.Context) {
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
