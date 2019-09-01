package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func CreateSourcesView(g *gocui.Gui) error {
	viewName := "sources"

	if sourcesView, err := g.SetView(viewName, 0, maxY/8, maxX/6, maxY/8+maxY/8); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
		sourcesView.Highlight = true
		sourcesView.Title = "Sources"
		sourcesView.SelBgColor = gocui.ColorDefault
		sourcesView.SelFgColor = gocui.ColorRed
		for index, source := range Sources {
			fmt.Fprintln(sourcesView, fmt.Sprintf("%d: ", index), source)
		}

	}
	views = append(views, viewName)
	return nil
}
