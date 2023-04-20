package main

import (
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

	fmt.Printf("%+v\n", searchResults)

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

	// if scanned > 0 {
	// 	fmt.Println()
	// }
	fmt.Printf(
		color.New(color.FgGreen).Add(color.Bold).Sprintf("Scanned files: ") + strconv.Itoa(len(searchResults)) +
			color.New(color.FgGreen).Add(color.Bold).Sprintf("\nValid documents: ") + strconv.Itoa(valid) +
			// color.New(color.FgYellow).Add(color.Bold).Sprintf("\nSuspicious files (%v): ", len(suspiciousFiles)) + strings.Join(suspiciousFiles, ", ") +
			color.New(color.FgRed).Add(color.Bold).Sprintf("\nInfected files (%v): ", len(infectedFiles)) + strings.Join(infectedFiles, ", ") + "\n",
	)
}
