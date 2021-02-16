/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an entry",
	Long:  `delete an entry from the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get key
		key, _ := cmd.Flags().GetString("key")
		if key == "" {
			fmt.Println("Not a valid key:", key)
			return
		}

		// Remove data
		err := deleteEntry(key)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().String("key", "", "Key to delete")
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}
	data = append(data[:i], data[i+1:]...)

	// Update the index - key does not exist any more
	// delete(index, key)

	err := saveJSONFile(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}
