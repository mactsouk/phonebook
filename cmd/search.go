/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for the number",
	Long: `search whether a telephone number exists in the
	phone book application or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get key
		searchKey, _ := cmd.Flags().GetString("key")
		if searchKey == "" {
			fmt.Println("Not a valid key:", searchKey)
			return
		}
		t := strings.ReplaceAll(searchKey, "-", "")

		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		// Search for it
		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*temp)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("key", "k", "", "Key to search")
}

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	return &data[i]
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}
