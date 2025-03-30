package csv_viewer

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/storage"
)


func populateTable(_data []map[string]interface{} ) *widget.Table {
	// for id, line := range _data {
	// 	println(id, " ")
	// 	for key, value := range line {
	// 		println("\t", key, " ", value.(string), " ")
	// 	}
	// }

	return widget.NewTable(
		func() (int, int) {
			return len(_data), len(_data[0])
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(
				fyne.NewSize(200, 40),
				// widget.NewEntry(),
				widget.NewRichText(),
			)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			richText := container.Objects[0].(*widget.RichText)
			text := _data[i.Row][header[i.Col]].(string)
			
			if i.Row == 0 {
				richText.ParseMarkdown("**" + text + "**")
			} else {
				richText.ParseMarkdown(text)
			}

			// TODO - In place cell change
			// entry := container.Objects[0].(*widget.Entry)
			// entry.SetText(_data[i.Row][header[i.Col]].(string))
			// if i.Row == 0 {
			// 	entry.Disable()
			// }

			// ! BUG - on window size change cell are messing up
			// entry.OnChanged = func(s string) {
			// 	fileContent[i.Row][header[i.Col]] = s
			// }
		},
	)
}


var query_editor_input *widget.Entry

var window_title	string	= "Opal - CSV Viewer"
var window_icon		string	= "./opal.png"
var window_width	float32 = 1900.0/2
var window_height	float32 = 1080.0/2


func MainGUI() {
	app := app.New()
	window := app.NewWindow(window_title)
	icon, _ := fyne.LoadResourceFromPath(window_icon)
	window.SetIcon(icon)
	window.Resize(fyne.NewSize(window_width, window_height))

	// var data [][]string
	var table *widget.Table
	var currentFile fyne.URI


	csv_scroll := container.NewScroll(widget.NewLabel("To begin, open a CSV file above"))


	csv_toolbar := widget.NewToolbar(
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

						data := readCSV(reader.URI().Path(), ",")
						table = populateTable(data)
						table.Refresh()

						csv_scroll.Content = table
						csv_scroll.Refresh()

						currentFile = reader.URI()
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
				if currentFile == nil {
					// dialogWindow := nil

					// dialogWindow.Resize(fyne.NewSize(800, 600))
					// dialogWindow.Show()
				} else {
					writer, err := storage.Writer(currentFile)
					if err != nil {
						dialog.ShowError(err, window)
						return
					}
					defer writer.Close()

					for _, row := range fileContent {
						for _, title := range header {
							_, err = writer.Write(([]byte)(row[title].(string)))
							if err != nil {
								dialog.ShowError(err, window)
								return
							}
						}
						_, err = writer.Write(([]byte)("\n"))
						if err != nil {
							dialog.ShowError(err, window)
							return
						}
					}
					
				}
			},
		),
	)



	sql_runQueryButton := widget.NewButton(
		"",
		func() {
			queryResult := executeQuery()
			if queryResult.err != "SUCCESS" {
				dialogWindow := dialog.NewInformation(
					"SQL Error",
					queryResult.err,
					window,
				)
				dialogWindow.Resize(fyne.NewSize(200, 90))
				dialogWindow.Show()
				return
			}
			table = populateTable(queryResult.result)
			table.Refresh()
		},
	)
	sql_runQueryButton.SetIcon(theme.MediaPlayIcon())
	sql_toolbar_run := container.NewHBox(
		sql_runQueryButton,
	)

	sql_clearQueryButton := widget.NewButton(
		"",
		func() {
			query_editor_input.SetText("")
		},
	)
	sql_clearQueryButton.SetIcon(theme.ContentClearIcon())
	sql_toolbar_clear := container.NewHBox(
		sql_clearQueryButton,
	)

	query_editor_input = widget.NewEntry()
	query_editor_input.Resize(fyne.NewSize(500, 40))
	query_editor_input.Refresh()
	query_editor := container.NewScroll(query_editor_input)
	query_editor.Resize(fyne.NewSize(500, 40))
	query_editor.Refresh()


	sql_editor := container.NewBorder(
		nil, nil,
		sql_toolbar_run, sql_toolbar_clear,
		query_editor,
	)

	csv_viewer := container.NewBorder(
		sql_editor,
		nil, nil, nil,
		csv_scroll,
	)


	content := container.NewBorder(
		csv_toolbar,
		nil, nil, nil,
		csv_viewer,
	)

	window.SetContent(content)

	window.ShowAndRun()
}