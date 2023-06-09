package main

import (
	"awesomeInvoice/internal/repository"
	"awesomeInvoice/internal/repository/dbrepo"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"log"
	"net/http"
	"os"
)

func csvRead() [][]string {
	fd, err := os.Open("templates/players.html")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
	}

	fmt.Println("File opened successfully!")
	defer fd.Close()

	fileReader := csv.NewReader(fd)

	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println("Got error while reading file:", err)
	}
	fmt.Println(records)
	return records
}

func loadCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Got error while reading file:", err)
		os.Exit(1)
	}
	return records
}

func heading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				err := m.FileImage("internal/assets/images/letter-r.jpg", props.Rect{
					Center:  true,
					Percent: 100,
				})
				if err != nil {
					fmt.Println("Got error while opening file:", err)
				}
				fmt.Println("PDF created successfully!")
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Services provided by Reticent Services LLC.", props.Text{
				Top:   5,
				Size:  16,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getBlueColor(),
			})
		})
	})
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   33,
		Green: 91,
		Blue:  191,
	}
}

func sendToInvoice(m pdf.Maroto) {
	tableHeadings := []string{"Player Name", "Position", "Team", "Age", "Jersey", "Retired"}
	content := csvRead()
	lightBlue := getLightBlueColor()
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice", props.Text{
				Top:    2,
				Size:   12,
				Color:  color.NewWhite(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.TableList(tableHeadings, content, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 1, 1, 2, 1, 1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 1, 1, 2, 1, 1},
		},
		Align:                consts.Center,
		AlternatedBackground: &lightBlue,
		HeaderContentSpace:   2,
		Line:                 true,
	})
}

func getTealColor() color.Color {
	return color.Color{
		Red:   7,
		Green: 117,
		Blue:  117,
	}
}

func getLightBlueColor() color.Color {
	return color.Color{
		Red:   33,
		Green: 91,
		Blue:  191,
	}
}

type application struct {
	DSN string
	DB  repository.DatabaseRepo
}

const port = ":8080"

func main() {
	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=players sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")

	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		fmt.Println("Got error while connecting to database:", err)
		os.Exit(1)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetPageMargins(10, 10, 10)
	heading(m)
	sendToInvoice(m)

	err = m.OutputFileAndClose("internal/assets/reports/invoice.pdf")
	if err != nil {
		fmt.Println("Got error while creating pdf:", err)
		os.Exit(1)
	}
	log.Println("Starting server on port", port)

	err = http.ListenAndServe(port, app.routes())
}
