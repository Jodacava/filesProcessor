package fileProcess

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ErrorBuffer struct {
	Message string `json:"message"`
}

type CsvUploadResponse struct {
	Message        string `json:"message"`
	FsDocumentPath string `json:"fsDocumentPath"`
}

var ExpectedHeader = []string{"ID", "MONTH", "DAY", "TRANSACTION"}

func ProcessHandler(service ServerBase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userEmail := ctx.Request.FormValue("user-email")
		file, fileHeaders, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, CsvUploadResponse{Message: err.Error()})
			return
		}

		fileNameSplit := strings.Split(fileHeaders.Filename, ".")
		if len(fileNameSplit) == 1 {
			ctx.JSON(http.StatusBadRequest, CsvUploadResponse{Message: "File must have a file extension."})
			return
		}

		fileExtension := fileNameSplit[len(fileNameSplit)-1]
		if fileExtension != "csv" {
			ctx.JSON(http.StatusBadRequest, CsvUploadResponse{Message: "File must have a .CSV file extension."})
			return
		}

		csvReader := csv.NewReader(file)
		if errHeader := headerValidations(csvReader); errHeader != nil {
			ctx.JSON(
				http.StatusBadRequest, CsvUploadResponse{Message: errHeader.Error()},
			)
			return
		}

		csvDocumentArrayToValidate, errorBodyParsing := validateBodyParsing(csvReader)
		if errorBodyParsing != nil {
			ctx.JSON(
				http.StatusBadRequest, CsvUploadResponse{Message: errorBodyParsing.Error()},
			)
			return
		}
		defer file.Close()

		response, code, errSender := service.ProcessFile(csvDocumentArrayToValidate, userEmail)
		if errSender != nil {
			ctx.JSON(code, ErrorBuffer{Message: errSender.Error()})
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func validateBodyParsing(csvReader *csv.Reader) ([][]string, error) {
	csvDocumentArrayToValidate, errReader := csvReader.ReadAll()
	if errReader != nil {
		messageError := "Error while parsing csv file, make sure you have a valid csv file"
		csvError, ok := errReader.(*csv.ParseError)
		if ok {
			messageError = fmt.Sprintf(
				"Error while parsing csv file. Error: %s", csvError.Error(),
			)
		}
		return nil, errors.New(messageError)
	}
	return csvDocumentArrayToValidate, nil
}

func headerValidations(csvReader *csv.Reader) error {
	csvHeader, errHeaderReader := csvReader.Read()
	if errHeaderReader != nil {
		messageError := fmt.Sprintf(
			"Can't parsing csv HEADER file. Error: %s", errHeaderReader.Error(),
		)
		return errors.New(messageError)
	}
	return validateHeader(csvHeader)
}

func validateHeader(header []string) error {
	headerErrors := []string{}
	if len(header) != len(ExpectedHeader) {
		headerErrors = append(
			headerErrors,
			fmt.Sprintf("Header: %v, must have this separated comma format: %v", header, ExpectedHeader),
		)
	}
	for colIndex, colData := range header {
		expData := ExpectedHeader[colIndex]
		if colData != expData {
			headerErrors = append(
				headerErrors,
				fmt.Sprintf(fmt.Sprintf("column: %v, must be equal to: %v", colData, expData)),
			)
		}
	}
	if len(headerErrors) > 0 {
		return errors.New(strings.Join(headerErrors, ". "))
	}
	return nil
}
