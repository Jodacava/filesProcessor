package fileProcess

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	repo        RepositoryBase
	debit       float32
	debitCount  float32
	credit      float32
	creditCount float32
}

type ServerBase interface {
	ProcessFile(fileArray [][]string, userEmail string) (CsvUploadResponse, int, error)
}

type CsvData struct {
	id          int
	month       int
	day         int
	transaction float32
}

type AdditionalData struct {
	Tittle string
	Value  string
}

func NewServer(repo RepositoryBase) ServerBase {
	return Server{repo: repo}
}

func (s Server) ProcessFile(fileArray [][]string, userEmail string) (CsvUploadResponse, int, error) {
	monthCount := map[string]int{}
	for _, rowData := range fileArray {
		val, errVal := strconv.ParseFloat(rowData[3], 32)
		if errVal != nil {
			continue
		}
		month, errMonth := strconv.Atoi(rowData[1])
		if errMonth != nil {
			continue
		}
		monthName := time.Month(month).String()
		monthCount[monthName] += 1
		if val > 0 {
			s.debit += float32(val)
			s.debitCount += 1
			continue
		}
		s.credit += float32(val)
		s.creditCount += 1
	}
	err := s.repo.EmailSender(s.HttpMessageBuilder(monthCount), userEmail)
	if err != nil {
		return CsvUploadResponse{}, http.StatusInternalServerError, err
	}
	return CsvUploadResponse{
		Message: "Email send",
	}, http.StatusBadRequest, nil
}

func (s Server) HttpMessageBuilder(monthTransactions map[string]int) []AdditionalData {
	data := []AdditionalData{}
	for month, totalPerMonth := range monthTransactions {
		data = append(data, AdditionalData{
			Tittle: fmt.Sprintf("Number of transactions in %s", month),
			Value:  fmt.Sprintf("%v", totalPerMonth)})
	}
	totalDebit := fmt.Sprintf("%v", s.debit/s.debitCount)
	data = append(data, AdditionalData{Tittle: "Average debit amount", Value: fmt.Sprintf("%v", totalDebit)})
	totalCredit := fmt.Sprintf("%v", s.credit/s.creditCount)
	data = append(data, AdditionalData{Tittle: "Average credit amount", Value: fmt.Sprintf("%v", totalCredit)})
	totalBalance := fmt.Sprintf("%v", s.credit-s.debit)
	data = append(data, AdditionalData{Tittle: "Total balance is", Value: fmt.Sprintf("%v", totalBalance)})

	return data
}
