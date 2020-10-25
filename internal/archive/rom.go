package archive

// Rom represents the contents a zip archive
type Rom struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Filename string `json:"filename"`
	Console  string `json:"console"`
	Version  int    `json:"version"`
}
