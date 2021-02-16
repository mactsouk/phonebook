/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type Entry struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Tel        string `json:"tel"`
	LastAccess string `json:"lastaccess"`
}

// JSONFILE resides in the current directory
var JSONFILE = "./data.json"

type PhoneBook []Entry

var data = PhoneBook{}
var index map[string]int

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = DeSerialize(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = Serialize(&data, f)
	if err != nil {
		return err
	}

	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func initS(N, S, T string) *Entry {
	// Both of them should have a value
	if T == "" || S == "" {
		return nil
	}

	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func setJSONFILE() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		JSONFILE = filepath
	}

	_, err := os.Stat(JSONFILE)
	if err != nil {
		fmt.Println("Creating", JSONFILE)
		f, err := os.Create(JSONFILE)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	fileInfo, err := os.Stat(JSONFILE)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", JSONFILE)
	}
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phonebook",
	Short: "A phone book application",
	Long:  `This is a Phone Book application that uses JSON records.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := setJSONFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readJSONFile(JSONFILE)
	// io.EOF is fine because it means the file is empty
	if err != nil && err != io.EOF {
		return
	}
	createIndex()

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
