package archive

import (
	"github.com/PuerkitoBio/goquery"
)

func filter(filename string) bool {
	valid := !isBeta(filename) &&
		!isVirtualConsole(filename) &&
		!isUnlisted(filename) &&
		!isSample(filename) &&
		!isProto(filename) &&
		!isBIOS(filename) &&
		!isGameboyAdvanceVideo(filename) &&
		!isNickelodeonMovies(filename) &&
		!isGerman(filename) &&
		!isFrench(filename) &&
		!isSpanish(filename)
	if valid {
		return true
	}
	return false
}

func scrapeConsoleGames(consoleURL string) []*goquery.Selection {
	doc := getHTML(consoleURL)
	var sel []*goquery.Selection
	doc.Find("#maincontent > div > table > tbody > tr > td > a").Each(func(i int, s *goquery.Selection) {
		sel = append(sel, s)
	})
	return sel
}

func importAllGamesDB() {

}
