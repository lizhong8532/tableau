package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/tealeg/xlsx"
	"fmt"
	"strconv"
  "os"
  //"io"
  "io"
)


type Column struct {
	Title string 	`json:"title"`
	Values []int `json:"values"`
}

type Sheet struct {
	Sheet string `json:"sheet"`
  Columns	[]Column	`json:"columns"`
  Labels []string `json:"labels"`
}

func main () {
	app := iris.New()
	//fmt.Println(len(a))

	app.Adapt(httprouter.New()) // <--- or gorillamux.New()
  irisDev := os.Getenv("IRIS_DEV")
  if irisDev == "IRIS_DEV" {
    app.Adapt(iris.DevLogger())
  }

  app.StaticWeb("/static", "./static")

	// HTTP Method: GET
	// PATH: http://127.0.0.1/
	// Handler(s): index
	app.Post("/api/import", index)
	app.Listen(":9111")
}

func index(ctx *iris.Context) {

  file, _, err := ctx.Request.FormFile("file")

  if err != nil {
    fmt.Println(err)
    ctx.JSON(iris.StatusBadRequest, err.Error())
    return
  }

  defer file.Close()

  f, err := os.Create("../test/tmp.xlsx")
  if err != nil {
    fmt.Println(err)
    ctx.JSON(iris.StatusBadRequest, err.Error())
    return
  }

  defer f.Close()
  io.Copy(f, file)

	xlFile, err := xlsx.OpenFile("../test/tmp.xlsx")
	//xlFile, err := xlsx.OpenFile("../test/test.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	arr := make([]Sheet, len(xlFile.Sheets))

	for sheetIndex, sheet := range xlFile.Sheets {

    maxCols := 0
    if len(sheet.Rows) >= 1 {
      maxCols = len(sheet.Rows[0].Cells) - 1
    }

    arr[sheetIndex] = Sheet{}
    arr[sheetIndex].Sheet = sheet.Name
    arr[sheetIndex].Columns = make([]Column, maxCols)
    arr[sheetIndex].Labels = make([]string, len(sheet.Rows) - 1)

		for rowIndex, row := range sheet.Rows {

			rowData := Column{}

			for colIndex, cell := range row.Cells {
				text, _ := cell.String()

				if rowIndex == 0 && colIndex != 0 {
					rowData.Values = make([]int, len(sheet.Rows) - 1)
					rowData.Title = text
					arr[sheetIndex].Columns[colIndex - 1] = rowData
				} else if colIndex != 0 {
					value, err := strconv.Atoi(text)

					if err != nil {
						value = 0
					}

					arr[sheetIndex].Columns[colIndex - 1].Values[rowIndex - 1] = value
				} else if colIndex == 0 && rowIndex != 0 {
          //fmt.Println(rowIndex, colIndex)
          arr[sheetIndex].Labels[rowIndex - 1] = text
        }
			}
		}

	}

	ctx.JSON(iris.StatusOK, arr)

}
