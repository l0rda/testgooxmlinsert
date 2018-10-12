package main // import "github.com/l0rda/testgooxmlinsert"

import (
	"log"

	"baliance.com/gooxml/spreadsheet"
)

func test1() {
	wb, err := spreadsheet.Open("test.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer wb.Close()
	sheet, err := wb.GetSheet("test")
	if err != nil {
		log.Fatal(err)
	}
	sheet.InsertRow(2)
	wb.SaveToFile("out1.xlsx")
}

func main() {
	test1()
}
