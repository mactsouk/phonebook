/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert new data",
	Long:  `This command inserts new data into the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the data
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Not a valid name:", name)
			return
		}

		surname, _ := cmd.Flags().GetString("surname")
		if surname == "" {
			fmt.Println("Not a valid surname:", surname)
			return
		}

		tel, _ := cmd.Flags().GetString("telephone")
		if tel == "" {
			fmt.Println("Not a valid telephone:", tel)
			return
		}

		t := strings.ReplaceAll(tel, "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", tel)
			return
		}

		temp := initS(name, surname, t)
		if temp == nil {
			fmt.Println("Not a valid record:", temp)
			return
		}

		// Insert data
		err := insert(temp)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("name", "n", "", "name value")
	insertCmd.Flags().StringP("surname", "s", "", "surname value")
	insertCmd.Flags().StringP("telephone", "t", "", "telephone value")
}

func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)

	// Save the data
	err := saveJSONFile(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}
