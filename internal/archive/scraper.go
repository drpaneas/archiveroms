package archive

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

func getHTML(page string) (doc *goquery.Document) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	standardClient := retryClient.StandardClient() // *http.Client

	// Request the HTML page.
	res, err := standardClient.Get(page)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func isEnglish(text string) bool {
	if strings.Contains(text, "(USA).zip") || strings.Contains(text, "(World)") || strings.Contains(text, "(USA, ") {
		return true
	}
	return false
}

func isBeta(text string) bool {
	if strings.Contains(text, "(Beta)") {
		return true
	}
	return false
}

func isVirtualConsole(text string) bool {
	if strings.Contains(text, "(Virtual Console)") {
		return true
	}
	return false
}

func isRev2(text string) bool {
	if strings.Contains(text, "(Rev 2)") {
		return true
	}
	return false
}

func isRev1(text string) bool {
	if strings.Contains(text, "(Rev 1)") {
		return true
	}
	return false
}

func isUnlisted(text string) bool {
	if strings.Contains(text, "(Unl)") {
		return true
	}
	return false
}

func isSample(text string) bool {
	if strings.Contains(text, "(Sample)") {
		return true
	}
	return false
}

func isProto(text string) bool {
	if strings.Contains(text, "(Proto") {
		return true
	}
	return false
}

func isBIOS(text string) bool {
	if strings.Contains(text, "[BIOS]") {
		return true
	}
	return false
}

func isGameboyAdvanceVideo(text string) bool {
	if strings.Contains(text, "Game Boy Advance Video") {
		return true
	}
	return false
}

func isNickelodeonMovies(text string) bool {
	if strings.Contains(text, "(Nickelodeon Movies)") {
		return true
	}
	return false
}

func isCollection(text string) bool {
	if strings.Contains(text, "Collection") {
		return true
	}
	return false
}

func isGerman(text string) bool {
	if strings.Contains(text, "(De)") {
		return true
	}
	return false
}

func isSpanish(text string) bool {
	if strings.Contains(text, "(Es)") {
		return true
	}
	return false
}

func isFrench(text string) bool {
	if strings.Contains(text, "(Fr)") {
		return true
	}
	return false
}

func gameAlreadyExistsOnTheSameConsole(gameTitle, gameConsole string) (previousFilename string, exists bool) {
	for _, previousGame := range GamesDB {
		if gameTitle == previousGame.Title {
			if gameConsole == previousGame.Console {
				return previousGame.Filename, true
			}
		}
	}
	return "", false
}

func isHigherVersion(currentGameVersion, previousGameVersion int) bool {
	if currentGameVersion > previousGameVersion {
		return true
	}
	return false
}

func isLowerVersion(currentGameVersion, previousGameVersion int) bool {
	if currentGameVersion < previousGameVersion {
		return true
	}
	return false
}

func isSameVersion(currentGameVersion, previousGameVersion int) bool {
	if currentGameVersion == previousGameVersion {
		return true
	}
	return false
}

func getFilename(text string) string {
	parts := strings.Split(text, "/")
	return parts[1]
}

func getTitle(text string) string {
	if isEnglish(text) {
		parts := strings.Split(text, " (")
		return parts[0]
	}
	log.Fatalf("Cannot get the title for %s. Parsing non-USA titles is not supported.\n", text)
	return ""
}

func getVersion(gameFilename string) (version int) {
	if isRev2(gameFilename) {
		version = 2
	} else if isRev1(gameFilename) {
		version = 1
	} else {
		version = 0
	}
	return version
}

func getGameDBIndex(gameFilename string) (index int) {
	for index, value := range GamesDB {
		if value.Filename == gameFilename {
			return index
		}
	}
	log.Fatalf("Couldn't find the index of %s in the database", gameFilename)
	return 0
}

func hasBiggerFilename(currentGameFilename, previousGameFilename string) bool {
	lenCurrent := len(currentGameFilename)
	lenPrevious := len(previousGameFilename)
	if lenCurrent > lenPrevious {
		return true
	}
	return false
}

func hasSmallerFilename(currentGameFilename, previousGameFilename string) bool {
	lenCurrent := len(currentGameFilename)
	lenPrevious := len(previousGameFilename)
	if lenCurrent < lenPrevious {
		return true
	}
	return false
}

func replaceGameEntry(title, link, filename, console string, version, index int) {
	GamesDB[index].Title = title
	GamesDB[index].Link = link
	GamesDB[index].Filename = filename
	GamesDB[index].Console = console
	GamesDB[index].Version = version
}
