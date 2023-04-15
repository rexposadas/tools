package cmd

import (
	"github.com/rexposadas/tools/models"
	"github.com/spf13/cobra"
	"log"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "does everything to process a given file",
	Long:  `Takes a file and does all processing`,
	Run: func(cmd *cobra.Command, args []string) {

		d := models.NewData(file)
		log.Printf("%+v", d.Genesis)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
