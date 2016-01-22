package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// CreatePdfFrom creates a PDF file from data
func CreatePdfFrom(data []map[string]interface{}) ([]byte, error) {

	// Check to length "[]map[string]interface{}"
	if len(data) == 0 {
		return nil, errors.New("Data is empty")
	}

	// Take a html file
	file, err := os.Create("complited.html")
	if err != nil {
		return nil, err
	}

	// Parse from template
	t := template.New("table.html")
	tempParse, err := t.ParseFiles("table.html")
	if err != nil {
		return nil, err
	}

	// Excute to html file
	err = tempParse.Execute(file, data)
	if err != nil {
		return nil, err
	}

	// converting "table.html" file to "./output/`time.Now()``.pdf"
	timeNow := time.Now()
	outputFile := "./output/" + timeNow.String() + ".pdf"
	err = exec.Command("wkhtmltopdf", "complited.html", outputFile).Run()
	if err != nil {
		return nil, err
	}

	// Take a pdf file
	outputPDF, err := os.OpenFile(outputFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Read a file
	pdf, err := ioutil.ReadAll(outputPDF)
	if err != nil {
		return nil, err
	}

	// Remove a file
	err = os.Remove(outputFile)
	if err != nil {
		return nil, err
	}

	// Return bytes
	return pdf, nil
}
