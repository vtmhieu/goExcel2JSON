package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FerdinaKusumah/excel2json"
)



func main() {
	var (
		result []*map[string]interface{}
		err    error
		url    = "https://www.wisdomaxis.com/technology/software/data/for-reports/Data%20Refresh%20Sample%20Data.xlsx"
		// select sheet name
		sheetName = "Sheet1"
		// select only selected field
		// if you want to show all headers just passing nil or empty list
		headers = []string{"Profit", "Shipping Cost", "Unit Price"}
	)
	if result, err = excel2json.GetExcelFileUrl(url, sheetName, headers); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	for _, val := range result {
		result, _ := json.Marshal(val)
		fmt.Println(string(result))
	}
}
