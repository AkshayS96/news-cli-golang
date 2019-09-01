package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// CreateSideBarView this is some comment
func CreateSideBarView(g *gocui.Gui) error {
	viewName := "category"
	if sidebar, err := g.SetView(viewName, 0, maxY/2, maxX/6, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		sidebar.Title = "Category"
		sidebar.Wrap = true
		sidebar.Highlight = true
		sidebar.SelBgColor = gocui.ColorDefault
		sidebar.SelFgColor = gocui.ColorRed
		for index, category := range Categories {
			fmt.Fprintln(sidebar, fmt.Sprintf("%d: ", index), category)
		}
	}
	views = append(views, viewName)
	return nil
}
