package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

func loadNewsFromQuery(g *gocui.Gui, query string) error {
	var resp *http.Response
	var contentView *gocui.View
	var err error
	query = strings.TrimSpace(query)
	queries := strings.Split(query, "\n")
	query = strings.Join(queries, " ")

	base, err := url.Parse("https://newsapi.org/v2/everything")
	if err != nil {
		return err
	}
	// url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?q=%v&apiKey=%v", query, Apikey)
	params := url.Values{}
	params.Add("q", query)
	params.Add("apiKey", Apikey)
	base.RawQuery = params.Encode()
	resp, err = http.Get(base.String())
	if err != nil {
		log.Panicln(err)
		return err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["articles"] == nil {
		return nil
	}
	Articles = result["articles"].([]interface{})

	if contentView, err = g.View("content"); err != nil {
		log.Panicln(err)
		return err
	}
	contentView.Clear()
	if len(Articles) == 0 {
		fmt.Fprintln(contentView, "No articles are present for this query")
		return nil
	}
	for index, article := range Articles {
		title := article.(map[string]interface{})["title"].(string)
		author := article.(map[string]interface{})["author"]
		url := article.(map[string]interface{})["url"]
		fmt.Fprintln(contentView, fmt.Sprintf("%v:", index), title, author, "\n", url, "\n")
	}
	return nil
}

func deleteStoryView(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView("content"); err != nil {
		return err

	}
	if err := g.DeleteKeybinding("story", gocui.KeyCtrlD, gocui.ModNone); err != nil {
		return err
	}
	deleteError := g.DeleteView("story")
	if deleteError != nil {
		return deleteError
	}

	return nil
}

func openWholeStory(g *gocui.Gui) error {
	if view, err := g.SetView("story", maxX/4, maxY/6, maxX-maxX/4, maxY-maxY/6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		view.Title = Articles[CurrentArticle].(map[string]interface{})["title"].(string)
		description := Articles[CurrentArticle].(map[string]interface{})["description"].(string)
		content := Articles[CurrentArticle].(map[string]interface{})["content"].(string)
		fmt.Fprintln(view, "\nDescription: ", description, "\nContent: ", content)
	}

	if err := g.SetKeybinding("story", gocui.KeyCtrlD, gocui.ModNone, deleteStoryView); err != nil {
		return err
	}
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
		splits := strings.Split(item, ":")
		if v.Name() == "country" {
			if item != "" {
				if len(splits) > 1 {
					Country = strings.TrimSpace(splits[1])
				}
			}

		} else if v.Name() == "category" {
			if item != "" {
				if len(splits) > 1 {
					Category = strings.TrimSpace(splits[1])
				}
			}
		} else if v.Name() == "sources" {
			if item != "" {
				if len(splits) > 1 {
					Source = strings.TrimSpace(splits[1])
				}
			}

		} else if v.Name() == "content" {
			if item != "" {
				CurrentArticle, err = strconv.ParseInt(strings.Split(item, ":")[0], 10, 64)
				openWholeStory(g)
			}
		}
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
	err = loadNewsFromQuery(g, query)
	if err != nil {
		return err
	}
	fmt.Fprintln(view, query)
	return nil
}

func sourceSubmitHandler(g *gocui.Gui, v *gocui.View) error {
	// Call the API using sources
	var contentView *gocui.View
	resp, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/everything?sources=%v&apiKey=%v", Source, Apikey))
	if err != nil {
		log.Panicln(err)
		return err
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["articles"] == nil {
		fmt.Println(fmt.Sprintf("https://newsapi.org/v2/everything?sources=%v&apiKey=%v", Source, Apikey))
		return nil
	}
	Articles = result["articles"].([]interface{})

	if contentView, err = g.View("content"); err != nil {
		log.Panicln(err)
		return err
	}
	contentView.Clear()
	if len(Articles) == 0 {
		fmt.Fprintln(contentView, "No articles are present for this query")
		return nil
	}
	for index, article := range Articles {
		title := article.(map[string]interface{})["title"].(string)
		// url := article.(map[string]interface{})["url"].(string)
		description := article.(map[string]interface{})["description"].(string)
		fmt.Fprintln(contentView, fmt.Sprintf("%v:", index), title, "\n\t", description, "\n")
	}
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

	if err := g.SetKeybinding("search", gocui.MouseLeft, gocui.ModNone, handleSubmitButton); err != nil {
		log.Panicln(err)
		return err
	}

	if err := g.SetKeybinding("sourcesSubmit", gocui.MouseLeft, gocui.ModNone, sourceSubmitHandler); err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
