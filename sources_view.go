package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func createSourceSubmit(g *gocui.Gui) error {
	viewName := "sourcesSubmit"
	if sourceSubmit, err := g.SetView(viewName, maxX/6-10, maxY/4-2, maxX/6, maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
		sourceSubmit.SelBgColor = gocui.ColorDefault
		sourceSubmit.SelFgColor = gocui.ColorRed
		sourceSubmit.BgColor = gocui.ColorBlue
		fmt.Fprintln(sourceSubmit, "Submit")
	}
	return nil
}

func CreateSourcesView(g *gocui.Gui) error {
	viewName := "sources"

	if sourcesView, err := g.SetView(viewName, 0, maxY/8, maxX/6, maxY/4); err != nil {
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
	createSourceSubmit(g)
	views = append(views, viewName)
	return nil
}
