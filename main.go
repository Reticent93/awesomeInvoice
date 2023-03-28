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
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice", props.Text{
				Top:   5,
				Size:  16,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.NewWhite(),
			})
		})
	})
}

func getTealColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 128,
		Blue:  128,
	}
}

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetPageMargins(10, 10, 10)
	heading(m)
	sendToInvoice(m)

	err := m.OutputFileAndClose("internal/assets/reports/invoice.pdf")
	if err != nil {
		fmt.Println("Got error while creating pdf:", err)
		os.Exit(1)
	}
}
