package archive

import (
	"github.com/PuerkitoBio/goquery"
)

// ConsolesDB contains all the information about each console
var ConsolesDB []Console

// Console struct
type Console struct {
	name           string
	URL            string
	gamesSelection []*goquery.Selection
}

// GetName returns the name of the console
func (c Console) GetName() string {
	return c.name
}

// GetURL returns the link of the console
func (c Console) GetURL() string {
	return c.URL
}

// GetGamesList returns the list of games for a given console
func (c Console) GetGamesList() []*goquery.Selection {
	return c.gamesSelection
}

// GetGame returns the a specific game
func (c Console) GetGame(i int) *goquery.Selection {
	return c.gamesSelection[i]
}

// NumberOfGames returns the number of games for a given console
func (c Console) NumberOfGames() int {
	return len(c.GetGamesList())
}

// Factory function to create new console instances
func newConsole(name, url string) Console {
	return Console{
		name:           name,
		URL:            url,
		gamesSelection: scrapeConsoleGames(url)}
}

type console struct {
	name string
	page string
}

const (
	gameboyColorURL   string = "https://ia801409.us.archive.org/view_archive.php?archive=/28/items/no-intro_romsets/no-intro%20romsets/Nintendo%20-%20Game%20Boy%20Color%20%2820201008-101210%29.zip"
	gameboyClassicURL string = "https://ia801409.us.archive.org/view_archive.php?archive=/28/items/no-intro_romsets/no-intro%20romsets/Nintendo%20-%20Game%20Boy%20%2820201007-070314%29.zip"
	gameboyAdvanceURL string = "https://ia801409.us.archive.org/view_archive.php?archive=/28/items/no-intro_romsets/no-intro%20romsets/Nintendo%20-%20Game%20Boy%20Advance%20%2820201012-021514%29.zip"
	snesURL           string = "https://ia801409.us.archive.org/view_archive.php?archive=/28/items/no-intro_romsets/no-intro%20romsets/Nintendo%20-%20Super%20Nintendo%20Entertainment%20System%20%28Combined%29%20%2820201013-141043%29.zip"
	nesURL            string = "https://ia801409.us.archive.org/view_archive.php?archive=/28/items/no-intro_romsets/no-intro%20romsets/Nintendo%20-%20Nintendo%20Entertainment%20System%20%2820200917-053714%29%20%5Bheadered_iNES2.0_NRS%282020-09-27%29%5D.zip" // headered iNES 2.0
)

// Console represents all the gaming consoles
var consolesList = []console{
	{
		name: "Nintendo Gameboy Color",
		page: gameboyColorURL,
	},
	{
		name: "Nintendo Gameboy Classic",
		page: gameboyClassicURL,
	},
	{
		name: "Nintendo Gameboy Advance",
		page: gameboyAdvanceURL,
	},
	{
		name: "Super Nintendo Entertainment System",
		page: snesURL,
	},
	{
		name: "Nintendo Entertainment System",
		page: nesURL,
	},
}

// CreateAllConsoles fetches metadata for all supported consoles
func CreateAllConsoles() {
	for i := 0; i < len(consolesList); i++ {
		ConsolesDB = append(ConsolesDB, newConsole(consolesList[i].name, consolesList[i].page))
	}
}
