package transport

import (
	"errors"
	"log"
	"strconv"

	"github.com/yusianglin11010/nihon-hitter/be/config"
	"google.golang.org/api/sheets/v4"
)

var (
	ErrNoDataFound = errors.New("no data found")
	ErrUnexpected  = errors.New("unexpected error")
)

type SheetRepo struct {
	Srv     *sheets.Service
	SheetID string
}

func NewSheetRepo(config config.Config, srv *sheets.Service) *SheetRepo {
	return &SheetRepo{
		Srv:     srv,
		SheetID: config.SheetID,
	}
}

func (s *SheetRepo) GetKana(readStart, readEnd int) ([]string, error) {
	readRange := "!A" + strconv.Itoa(readStart) + ":" + "A" + strconv.Itoa(readEnd)

	res := []string{}
	resp, err := s.Srv.Spreadsheets.Values.Get(s.SheetID, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
		return nil, ErrUnexpected
	}

	if len(resp.Values) == 0 {
		log.Fatalf("no data found")
		return nil, ErrNoDataFound
	} else {
		for _, row := range resp.Values {
			res = append(res, row[0].(string))
		}
	}
	return res, nil
}

func (s *SheetRepo) GetKanji(readStart, readEnd int) ([]string, error) {
	readRange := "!B" + strconv.Itoa(readStart) + ":" + "B" + strconv.Itoa(readEnd)
	res := []string{}
	resp, err := s.Srv.Spreadsheets.Values.Get(s.SheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
		return nil, ErrUnexpected
	}

	if len(resp.Values) == 0 {
		log.Fatalf("no data found")
		return nil, ErrNoDataFound
	} else {
		for _, row := range resp.Values {
			res = append(res, row[0].(string))
		}
	}
	return res, nil
}

func (s *SheetRepo) GetImi(readStart, readEnd int) ([]string, error) {
	readRange := "!C" + strconv.Itoa(readStart) + ":" + "C" + strconv.Itoa(readEnd)

	res := []string{}
	resp, err := s.Srv.Spreadsheets.Values.Get(s.SheetID, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
		return nil, ErrUnexpected
	}

	if len(resp.Values) == 0 {
		log.Fatalf("no data found")
		return nil, ErrNoDataFound
	} else {
		for _, row := range resp.Values {
			res = append(res, row[0].(string))
		}
	}
	return res, nil
}

func (s SheetRepo) GetRowNumber() int {
	resp, err := s.Srv.Spreadsheets.Values.Get(s.SheetID, "G1").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	rowNum, err := strconv.Atoi(resp.Values[0][0].(string))
	if err != nil {
		log.Fatalf("Unable to convert row number to int")
	}
	return rowNum
}
