package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
)

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
	tableToHeadings := []string{"Date", "Invoice No.", "Company", "Address", "City", "State", "Zip", "Phone", "Email", "Amount"}
	//tableFromHeadings := []string{"Date", "Invoice No.", "Company", "Address", "City", "State", "Zip", "Phone", "Email", "Amount"}
	//contentsTo := [][]string{{"2020-01-01", "INV-0001", "ABC Services", "123 Main St.", "New York", "NY", "10001", "212-555-1212"}, {"2020-01-02", "INV-0002", "ACME", "321 Main St.", "New York", "NY", "10001", "212-555-5555"}}
	//contentFrom := [][]string{{"2020-01-01", "INV-0001", "Reticent Services LLC", "123 Anywhere St.", "New York", "NY", "10001", "555-212-5555"}, {"2020-01-02", "INV-0002", "Reticent Services LLC", "123 Anywhere St.", "New York", "NY", "10001", "555-212-5555"}}
	content := [][]string{{"2020-01-01", "INV-0001", "ABC Services", "123 Main St.", "New York", "NY", "10001", "212-555-1212", "abc@abc.com", "$100.00"}}
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

	m.TableList(tableToHeadings, content, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 1, 1, 2, 1, 1, 1, 1, 2, 1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 1, 1, 2, 1, 1, 1, 1, 2, 1},
		},
		Align:                consts.Center,
		AlternatedBackground: &lightBlue,
		HeaderContentSpace:   2,
		Line:                 false,
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

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetPageMargins(20, 10, 20)
	heading(m)
	sendToInvoice(m)

	err := m.OutputFileAndClose("internal/assets/reports/invoice.pdf")
	if err != nil {
		fmt.Println("Got error while creating pdf:", err)
		os.Exit(1)
	}
}
