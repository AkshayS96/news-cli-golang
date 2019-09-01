package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

// Quit handles the ctrl + C event to close the application
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// MouseEventHandler handles the mouse click event on views
func mouseEventHandler(g *gocui.Gui, v *gocui.View) error {
	name := v.Name()
	if _, err := g.SetCurrentView(name); err != nil {
		return err
	}
	return nil
}

// HandleTabEvent handles the tab keypress event
func handleTabEvent(g *gocui.Gui, v *gocui.View) error {
	currentView = (currentView + 1) % len(views)
	if _, err := g.SetCurrentView(views[currentView]); err != nil {
		return err
	}
	return nil
}

func loadNews(g *gocui.Gui) error {

	resp, err:=
	return nil
}

func menuSelectHandler(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}
	_, cy := v.Cursor()
	if item, err := v.Line(cy); err != nil {
		item = ""
	} else {
		maxX, maxY = g.Size()
		var view *gocui.View
		var err error
		if view, err = g.View("footer"); err != nil {
			return err
		}
		view.Clear()
		if v.Name() == "country" {
			if item != "" {
				Country = item
			}

		} else if v.Name() == "category" {
			if item != "" {
				Category = item
			}
		} else if v.Name() == "sources" {
			if item != "" {
				Source = item
			}

		}
		loadNews(g)
		fmt.Fprintln(view, Country, Source, Category)
	}
	return nil
}

func clearFooter(g *gocui.Gui) error {
	var view *gocui.View
	var err error
	if view, err = g.View("footer"); err != nil {
		log.Panicln(err)
		return err
	}
	view.Clear()
	fmt.Fprintln(view, Country, Source, Category)
	return nil
}

func scrollView(v *gocui.View, dy int) error {
	if v != nil {
		v.Autoscroll = false
		ox, oy := v.Origin()
		if _, err := v.Word(ox, oy+dy); err != nil {
			return nil
		}
		if err := v.SetOrigin(ox, oy+dy); err != nil {
			return err
		}
	}
	return nil
}

func scrollEventDown(g *gocui.Gui, v *gocui.View) error {

	if err := scrollView(v, 1); err != nil {
		return err
	}
	return nil
}

func scrollEventUp(g *gocui.Gui, v *gocui.View) error {

	if err := scrollView(v, -1); err != nil {
		return err
	}
	return nil
}

func handleSubmitButton(g *gocui.Gui, v *gocui.View) error {

	var view, queryView *gocui.View
	var err error
	if queryView, err = g.View("query"); err != nil {
		log.Panicln(err)
		return err
	}
	query := queryView.ViewBuffer()
	if view, err = g.View("footer"); err != nil {
		log.Panicln(err)
		return err
	}
	view.Clear()
	fmt.Fprintln(view, query)
	return nil
}

// SetUpMouseAndKeyEvents wraps all the event in the single function
func SetUpMouseAndKeyEvents(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
		return err
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, mouseEventHandler); err != nil {
		log.Panicln(err)
		return err
	}

	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, menuSelectHandler); err != nil {
		log.Panicln(err)
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, scrollEventDown); err != nil {
		log.Panicln(err)
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, scrollEventUp); err != nil {
		log.Panicln(err)
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, scrollEventDown); err != nil {
		log.Panicln(err)
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, scrollEventUp); err != nil {
		log.Panicln(err)
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, handleTabEvent); err != nil {
		log.Panicln(err)
		return err
	}

	if err := g.SetKeybinding("submit", gocui.MouseLeft, gocui.ModNone, handleSubmitButton); err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
