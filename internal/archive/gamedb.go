package archive

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

var (
	// GamesDB is storage for all the games for all supported Consoles
	GamesDB []Rom
)

// getGameInfo fetches metadata for a given Console URL and dumps them into the Games array
// func getGameInfo(consoleName, consoleURL string) {
// 	doc := getHTML(consoleURL)
// 	doc.Find("#maincontent > div > table > tbody > tr > td > a").Each(func(i int, s *goquery.Selection) {
// 		importGameDB(s, consoleName)
// 	})
// }

func importGameDB(s *goquery.Selection, consoleName string) {
	text := s.Text()
	if isEnglish(text) {
		filename := getFilename(text)
		if filter(filename) {
			title := getTitle(filename)
			version := getVersion(filename)
			href, linkExists := s.Attr("href")
			if linkExists {
				link := ("https:" + href)
				previousFilename, gameExists := gameAlreadyExistsOnTheSameConsole(title, consoleName)
				if gameExists {
					fmt.Printf("Current : %s\nPrevious: %s\n", filename, previousFilename)
					previousGameIndex := getGameDBIndex(previousFilename)
					if isHigherVersion(version, GamesDB[previousGameIndex].Version) {
						fmt.Printf("Keeping : %s\n---------------------- \n", filename)
						// the current game has a newer version. Replace the old one with this one.
						replaceGameEntry(title, link, filename, consoleName, version, previousGameIndex)
					} else if isLowerVersion(version, GamesDB[previousGameIndex].Version) {
						// the current game has a lower version. Do nothing, keep the old one.
						fmt.Printf("Keeping : %s\n---------------------- \n", previousFilename)
					} else { // Both are the same version. They are just different Editions or languages or something else
						if hasBiggerFilename(filename, GamesDB[previousGameIndex].Filename) {
							// the current bigger filename. Replace the old one with this one.
							// Usually the bigger the filename the better edition (e.g. anniversary or special shit like that)
							replaceGameEntry(title, link, filename, consoleName, version, previousGameIndex)
							fmt.Printf("Keeping : %s\n---------------------- \n", filename)
						} else if hasSmallerFilename(filename, GamesDB[previousGameIndex].Filename) {
							// the current game has smaller filename. Do nothing, keep the old one.
							fmt.Printf("Keeping : %s\n---------------------- \n", previousFilename)
						} else { // No difference in the filename
							fmt.Printf("Keeping : Both\n---------------------- \n")
							GamesDB = append(GamesDB, Rom{Filename: filename, Title: title, Console: consoleName, Link: link, Version: version})
						}
					}
				} else { // Game doesn't exist, it's a new game. Add it to the database.
					GamesDB = append(GamesDB, Rom{Filename: filename, Title: title, Console: consoleName, Link: link, Version: version})
				}
			}
		}
	}
}

// GetAllConsoleGames runs gamePerConsole() for each console found in ConsolesDB
func GetAllConsoleGames() {
	for _, console := range ConsolesDB {
		gamePerConsole(console)
	}
}

// gamePerConsole takes a console and imports its games into GamesDB
func gamePerConsole(gamingSystem Console) {
	// loop through all the games of that pericular console
	for i := 0; i < gamingSystem.NumberOfGames(); i++ {
		game := gamingSystem.GetGame(i)
		consoleName := gamingSystem.name
		importGameDB(game, consoleName)
	}
}
