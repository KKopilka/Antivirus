package main

import (
	"encoding/json"
	"fmt"
	"kkopilka/AV/database"
	"kkopilka/AV/internal/avs"
	searchtree "kkopilka/AV/internal/search-tree"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func init() {
	// так пошли нахуй в stdout
	log.SetOutput(os.Stdout)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func loadDotEnv() error {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return godotenv.Load(exPath + "/.env")
}
func run() error {
	err := loadDotEnv()
	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Println("Programm started")
	if err := database.Open(); err != nil {
		return err
	}
	log.Println("Database connection opened")
	defer database.Close()

	if err := searchtree.BuildSearchTree(); err != nil {
		return err
	}

	ptree := searchtree.GetSearchTree()
	log.Println("Prefix tree loaded")

	if len(os.Args) > 1 {
		for index, scanDir := range os.Args {
			if index == 0 {
				continue
			}

			log.Println("Start search in location:", scanDir)
			if err := avs.FindSignatures(scanDir, ptree); err != nil {
				return err
			}

			printResults()
		}
	}

	return nil
}

func printResults() {
	searchResults := avs.SearchResults()

	if b, err := json.MarshalIndent(searchResults, "", "\t"); err == nil {
		log.Printf("Scan verbose results: %+v\n", string(b))
	}

	infectedFilesScan := map[string]struct{}{}
	for file, signStats := range searchResults {
		if len(signStats) > 0 {
			infectedFilesScan[file] = struct{}{}
		}
	}
	infectedFiles := []string{}
	for file, _ := range infectedFilesScan {
		infectedFiles = append(infectedFiles, file)
	}
	valid := len(searchResults) - len(infectedFiles)

	fmt.Printf(
		color.New(color.FgGreen).Add(color.Bold).Sprintf("Scanned files: ") + strconv.Itoa(len(searchResults)) +
			color.New(color.FgGreen).Add(color.Bold).Sprintf("\nValid documents: ") + strconv.Itoa(valid) +
			color.New(color.FgRed).Add(color.Bold).Sprintf("\nInfected files (%v): ", len(infectedFiles)) + strings.Join(infectedFiles, ", ") + "\n",
	)
}
