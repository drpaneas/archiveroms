package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/drpaneas/archiveroms/internal/archive"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		archive.CreateAllConsoles()  // Initialize all consoles
		archive.GetAllConsoleGames() // Import games of each console into GamesDB

		// Save to file
		dbFileName := "db.json"
		home, err := os.UserHomeDir()
		if err != nil {
			log.Printf("couldn't find the $HOME directory\nError: %s", err)

		}
		dbFile := home + "/" + dbFileName

		fileJSON, err := json.Marshal(archive.GamesDB)
		if err != nil {
			log.Print("Couldn't encode to JSON")
		}
		err = ioutil.WriteFile(dbFile, fileJSON, 0644)
		if err != nil {
			log.Printf("Couldn't update the db file %s\nError: %s", dbFile, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
