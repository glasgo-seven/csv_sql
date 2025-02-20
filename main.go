package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"

	// "fyne.io/fyne/v2/layout"
	"bufio"
	"fmt"
	"os"
)

/* PLAN
FILE = Read csv file lines
N_OF_COLS = len( FILE[0].split(',') )
N_OF_ROWS = len(FILE)

Make a table
[
	[ col, col, ... ],
	[ col, col, ... ],
	...
]
*/

// var data = [][]string{
// 	[]string{
// 		"1","Finland","7.769","1.340","1.587","0.986","0.596","0.153","0.393",
// 	},
// 	[]string{
// 		"2","Denmark","7.600","1.383","1.573","0.996","0.592","0.252","0.410",
// 	},
// }

func readFile(_fileName string) []string {
	file, err := os.Open(_fileName)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		text := scanner.Text()
		content = append(content, text)
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}

	return content
}

func getFileContent(_fileName string, _delimiter string) [][]string {
	content := readFile(_fileName)
	var parsed_content [][]string

	for _, line := range content {
		parsed_content = append(parsed_content, strings.Split(line, _delimiter) )
	}

	return parsed_content
}

func populateTable(_data [][]string) *widget.Table {
	return widget.NewTable(
		func() (int, int) {
			return len(_data), len(_data[0])
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(150, 40), widget.NewEntry())
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			entry := container.Objects[0].(*widget.Entry)
			entry.SetText(_data[i.Row][i.Col])
		},
	)
}


func main() {
	app := app.New()
	window := app.NewWindow("csv_viewer")
	window.Resize(fyne.NewSize(1900, 1080))

	var data [][]string
	var table *widget.Table

	scroll := container.NewScroll(widget.NewLabel(" ^ Open File"))

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.FileIcon(),
			func() {
				dialogWindow := dialog.NewFileOpen(
					func(reader fyne.URIReadCloser, err error) {
						if err != nil {
							dialog.ShowError(err, window)
							return
						}
						if reader == nil {
							return
						}

						fmt.Println(reader.URI().Path())

						data = getFileContent(reader.URI().Path(), ",")
						table = populateTable(data)
						table.Refresh()

						scroll.Content = table
						scroll.Refresh()
					},
					window,
				)
				dialogWindow.Resize(fyne.NewSize(800, 600))

				dialogWindow.Show()
			},

		),
		widget.NewToolbarAction(
			theme.DocumentSaveIcon(),
			func() {
				dialogWindow := dialog.NewFileSave(
					func(reader fyne.URIWriteCloser, err error) {
						if err != nil {
							dialog.ShowError(err, window)
							return
						}
						if reader == nil {
							return
						}

						fmt.Println(reader.URI().Path())

						data = getFileContent(reader.URI().Path(), ",")
						table = populateTable(data)
						table.Refresh()

						scroll.Content = table
						scroll.Refresh()
					},
					window,
				)

				dialogWindow.Resize(fyne.NewSize(800, 600))

				dialogWindow.Show()
			},
		),
	)


	// data := getFileContent("./happiness.csv", ",")

	// table := widget.NewTable(
	// 	func() (int, int) {
	// 		return len(data), len(data[0])
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return container.NewGridWrap(fyne.NewSize(150, 40), widget.NewEntry())
	// 	},
	// 	func(i widget.TableCellID, o fyne.CanvasObject) {
	// 		container := o.(*fyne.Container)
	// 		entry := container.Objects[0].(*widget.Entry)
	// 		entry.SetText(data[i.Row][i.Col])
	// 	},
	// )

	// table := container.NewVBox()
	// for _, line := range data {
	// 	row := container.NewHBox(
	// 		widget.NewButton(
	// 			"X",
	// 			func() {},
	// 		),
	// 	)
	// 	for _, cell := range line {
	// 		entry := widget.NewEntry()
	// 		entry.SetText(cell)
	// 		row.Add(entry)
	// 	}
	// 	table.Add(row)
	// }

	// scroll := container.NewScroll()

	content := container.NewBorder(
		toolbar,
		nil,
		nil,
		nil,
		scroll,
	)

	window.SetContent(content)
	// window.SetContent(scroll)

	window.ShowAndRun()
}