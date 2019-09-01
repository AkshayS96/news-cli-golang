package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// CreateContentView creates a view to show the content
func CreateContentView(g *gocui.Gui) error {
	viewName := "content"
	if content, err := g.SetView(viewName, maxX/6+1, 0, maxX-1, maxY-maxY/4-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		content.Title = "Content"
		content.Autoscroll = true
		content.Wrap = true
		content.SelBgColor = gocui.ColorGreen
		content.SelFgColor = gocui.ColorRed
		fmt.Fprintln(content, "hey this is the content view of my application")
	}
	views = append(views, viewName)
	return nil
}
