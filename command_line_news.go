package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

var maxX, maxY int

var (
	views       = []string{}
	currentView = -1
	Sources     = []string{"hey", "ejrjes"}
	Categories  = []string{"General", "Technology", "Business", "Entertainment", "Health", "Science", "Sports"}
	Countries   = []string{"us", "in", "ae", "ar", "at", "au", "be", "bg", "br", "ca", "ch", "cn", "co"}
	Country     = "1: us"
	Source      = ""
	Category    = "1. Technology"
	// "ae", "ar", "at", "au", "be", "bg", "br", "ca", "ch", "cn", "co", "cu", "cz", "de", "eg", "fr", "gb", "gr", "hk", "hu", "id", "ie", "il", "in", "it", "jp", "kr", "lt", "lv", "ma", "mx", "my", "ng", "nl", "no", "nz", "ph", "pl", "pt", "ro", "rs", "ru", "sa", "se", "sg", "si", "sk", "th", "tr", "tw", "ua", "us", "ve", "za"}
)

func loadDataFromAPI() {

}

func layout(g *gocui.Gui) error {
	maxX, maxY = g.Size()
	// creating the menu view
	CreateSideBarView(g)
	//creating the content view
	CreateContentView(g)
	//creating the footer view
	CreateFooterView(g)
	//creating the sources view
	CreateSourcesView(g)
	// creating the country view
	CreateCountriesView(g)
	//creating a query search field
	CreateQueryView(g)
	//Getting data from google Api
	loadDataFromAPI()
	return nil
}

func startApplication() error {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.SelBgColor = gocui.ColorGreen
	g.Cursor = true
	g.Mouse = true
	g.SetManagerFunc(layout)
	if err := SetUpMouseAndKeyEvents(g); err != nil {
		log.Panicln(err)
		return err
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
		return err
	}
	return nil
}

func main() {
	startApplication()
}
