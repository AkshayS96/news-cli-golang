package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// CreateFooterView creates a footer area for the application
func CreateFooterView(g *gocui.Gui) error {
	viewName := "footer"
	if footer, err := g.SetView(viewName, maxX/6+1, maxY-maxY/4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		footer.Wrap = true
		footer.Title = "Top HeadLines"
		footer.SelBgColor = gocui.ColorGreen
		footer.SelFgColor = gocui.ColorRed
		fmt.Fprintln(footer, Country, Source, Category)
	}
	views = append(views, viewName)
	return nil
}
