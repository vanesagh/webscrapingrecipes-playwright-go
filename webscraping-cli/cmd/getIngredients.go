/*
Copyright © 2023 Van GH
*/
package cmd

import (
	"fmt"
	scraper "webscraping/scraper"

	"github.com/spf13/cobra"
)

// getIngredientsCmd represents the getIngredients command
var getIngredientsCmd = &cobra.Command{
	Use:   "getIngredients",
	Short: "Get the Ingredients of a recipe from in a cooking website.",
	Long: `Get the Ingredients of a recipe from a cooking website and it is saved into a json file.
	Just type:
	 ./bin/webscraping-cli getIngredients URL...`,
	Run: func(cmd *cobra.Command, args []string) {
		var urlRecipe string
		if len(args) >= 1 && args[0] != "" {
			urlRecipe = args[0]
		}
		fmt.Println("Try to get ingredients from " + urlRecipe)
		scraper.ScrapForRecipe(urlRecipe)
	},
}

func init() {
	rootCmd.AddCommand(getIngredientsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getIngredientsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getIngredientsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
