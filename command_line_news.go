package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
)

var maxX, maxY int

var (
	views       = []string{}
	currentView = -1
	Sources     = []string{}
	Categories  = []string{"General", "Technology", "Business", "Entertainment", "Health", "Science", "Sports"}
	Countries   = []string{"us", "in", "ae", "ar", "at", "au", "be", "bg", "br", "ca", "ch", "cn", "co"}
	Country     = "1: us"
	Source      = ""
	Category    = "1: Technology"
	apikey      = os.Getenv("NEWAPIKEY")
	// "ae", "ar", "at", "au", "be", "bg", "br", "ca", "ch", "cn", "co", "cu", "cz", "de", "eg", "fr", "gb", "gr", "hk", "hu", "id", "ie", "il", "in", "it", "jp", "kr", "lt", "lv", "ma", "mx", "my", "ng", "nl", "no", "nz", "ph", "pl", "pt", "ro", "rs", "ru", "sa", "se", "sg", "si", "sk", "th", "tr", "tw", "ua", "us", "ve", "za"}
)

type SourceStructure struct {
	id          string
	name        string
	description string
	url         string
	category    string
	language    string
	country     string
}

func LoadDataFromAPI() error {
	//load sources and top headlines with defaul options
	var country, category, source string

	countryOption := strings.Split(Country, ": ")
	if len(countryOption) > 1 {
		country = countryOption[1]
	}
	categoryOption := strings.Split(Category, ": ")
	if len(categoryOption) > 1 {
		category = categoryOption[1]
	}
	sourceOption := strings.Split(Source, ": ")
	if len(sourceOption) > 1 {
		source = sourceOption[1]
	}
	//first load sources
	resp, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/sources?apiKey=%v", apikey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// var rawSources []SourceStructure
	fmt.Println(category, country, source)
	var result map[string]interface{}
	if err != nil {
		log.Panicln(err)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	sourcesList := result["sources"].([]interface{})
	for _, source := range sourcesList {
		source = source.(map[string]interface{})["id"]
		Sources = append(Sources, source.(string))
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY = g.Size()
	//Getting data from google Api
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
	return nil
}

func startApplication() error {
	if err := LoadDataFromAPI(); err != nil {
		return err
	}
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
