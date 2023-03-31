/*
Copyright © 2023 Van GH
*/
package cmd

import (
	"fmt"
	"log"
	scraper "webscraping/scraper"

	"github.com/spf13/cobra"
)

// getIngredientsCmd represents the getIngredients command
var getIngredientsCmd = &cobra.Command{
	Use:   "getIngredients",
	Short: "Getting the Ingredients of a recipe from in a cooking website.",
	Long: `Getting the Ingredients of a recipe from a cooking website and it is saved into a json file.
	Just type:
	 ./bin/webscraping-cli getIngredients URL...`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath != "" {
			fmt.Println(filePath)
			listUrls, err := scraper.OpenFile(filePath)
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Println(listUrls)
		} else {
			var urlRecipe string
			if len(args) >= 1 && args[0] != "" {
				urlRecipe = args[0]
			}
			fmt.Println("Try to get ingredients from " + urlRecipe)
			scraper.ScrapeForRecipe(urlRecipe)

		}

	},
}

func init() {
	rootCmd.AddCommand(getIngredientsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getIngredientsCmd.PersistentFlags().String("file", "", "A file with a list of urls")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getIngredientsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
