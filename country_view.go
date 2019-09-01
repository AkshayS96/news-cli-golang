package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func CreateCountriesView(g *gocui.Gui) error {
	viewName := "country"

	if countryView, err := g.SetView(viewName, 0, 0, maxX/6, maxY/8-1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
		countryView.Highlight = true
		countryView.Title = "Countries"
		countryView.SelFgColor = gocui.ColorRed
		countryView.SelBgColor = gocui.ColorDefault
		for index, country := range Countries {
			fmt.Fprintln(countryView, fmt.Sprintf("%d: ", index), country)
		}
	}
	views = append(views, viewName)
	return nil
}
