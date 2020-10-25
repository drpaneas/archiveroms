package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/drpaneas/archiveroms/internal/archive"
	"github.com/spf13/cobra"
)

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		//
		// Find the homedir and create the file
		dbFileName := "db.json"
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("couldn't find the $HOME directory\nError: %s", err)

		}
		dbFile := home + "/" + dbFileName
		// If the file exists parse it
		if FileExists(dbFile) {
			// Read the file
			fileJSON, err := ioutil.ReadFile(dbFile)
			if err != nil {
				log.Fatalf("Could not read the file.\nError: %s\n", err)
			}
			json.Unmarshal(fileJSON, &archive.GamesDB)

			// Print the file
			// for k, v := range archive.GameboyColor {
			// 	gameNumber := k + 1
			// 	fmt.Printf("%3d - %s\n", gameNumber, v.Title)
			// }

			// Search for game
			name := strings.Join(args, " ")
			for _, v := range archive.GamesDB {
				if strings.Contains(v.Title, name) {
					fmt.Printf("%s (%s) - %s\n", v.Title, v.Console, v.Link)
				}
			}

		} else {
			log.Fatal("No file")
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
