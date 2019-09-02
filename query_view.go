package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// CreateQueryView creates a search field for doing query

func submitButton(g *gocui.Gui) error {
	viewName := "search"
	if submitButton, err := g.SetView(viewName, maxX/15, maxY/3+6, maxX/10, maxY/2-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		submitButton.BgColor = gocui.ColorBlue
		submitButton.SelBgColor = gocui.ColorBlue
		submitButton.SelFgColor = gocui.ColorBlue
		fmt.Fprintln(submitButton, "Search")
	}
	views = append(views, viewName)
	return nil
}

func CreateQueryView(g *gocui.Gui) error {
	viewName := "query"
	if queryView, err := g.SetView(viewName, 0, maxY/4+2, maxX/6, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		queryView.Title = "Search with your query"
		queryView.Editable = true
		queryView.FgColor = gocui.ColorRed
		queryView.BgColor = gocui.ColorDefault
		queryView.SetCursor(0, 0)
		queryView.SetOrigin(0, 0)
	}
	views = append(views, viewName)
	submitButton(g)
	return nil
}
