package main

/*
TODO: Add DB Connection
TODO: Import the normalized data into the database
TODO: Score the CVSS report and assign the values to score
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Run the prep script to populate data directory
	if err := runPrepScript(); err != nil {
		log.Fatalf("Error running prep script: %v", err)
	}

	// Read and process the data files
	jsonFiles, err := os.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}
	var count int = 0 // count the number of files processed
	for _, file := range jsonFiles {
		if err := processJsonFiles(file); err != nil {
			log.Fatal(err)
		} else {
			count++
			fmt.Println("Processed ", count, " files")
			//TODO: Add the normalized data to the database (see db.go)

		}
	}
	//fmt.Println("Processed ", count, " files")
}

func processJsonFiles(file os.DirEntry) error {
	if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
		return fmt.Errorf("file is a directory or not a JSON file")
	}

	filePath := filepath.Join("data", file.Name())
	fmt.Println("Processing file: ", filePath)

	// Read JSON file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	// Unmarshal JSON into Findings struct
	var finding Findings
	if err := json.Unmarshal(jsonData, &finding); err != nil {
		log.Printf("Error unmarshaling JSON from %s: %v", filePath, err)
		return fmt.Errorf("error unmarshaling JSON from %s: %v", filePath, err)
	}

	// Normalize the data
	normalized := normalizeFinding(finding)
	for _, nv := range normalized {
		fmt.Printf("Normalized Data Validation: %+v\n", nv)

	}
	//TODO: Add the normalized data to the database

	return nil
}

func normalizeFinding(finding Findings) []NormalizedVuln {
	var normalized []NormalizedVuln

	// Get alias (use first alias if available, otherwise use ID)
	alias := finding.ID
	if len(finding.Aliases) > 0 {
		alias = finding.Aliases[0]
	}
	// sometimes summary is empty, and details contains the summary values
	summary := finding.Summary
	if summary == "" {
		summary = finding.Details
	}
	// Process each affected package
	for _, affected := range finding.Affected {
		// Only process supported ecosystems
		if !verifySupportedEcosystem(affected.Package.Ecosystem) {
			continue
		}

		// Extract introduced and fixed versions from ranges
		introduced := ""
		fixed := ""

		for _, r := range affected.Ranges {
			for _, event := range r.Events {
				if event.Introduced != "" {
					introduced = event.Introduced
				}
				if event.Fixed != "" {
					fixed = event.Fixed
				}
			}
		}

		// Extract severity score (simplified - you may want to parse CVSS scores)
		severity := 0.0
		if len(finding.Severity) > 0 {
			// For now, just set a placeholder - you'll need to parse the CVSS score
			severity = 0.0 // TODO: Parse CVSS score from finding.Severity[0].Score
		}

		normalized = append(normalized, NormalizedVuln{
			VulnID:     finding.ID,
			Alias:      alias,
			Ecosystem:  affected.Package.Ecosystem,
			Package:    affected.Package.Name,
			Introduced: introduced,
			Fixed:      fixed,
			Severity:   severity,
			Summary:    summary,
		})
	}

	return normalized
}

func verifySupportedEcosystem(ecosystem string) bool {
	// Case-insensitive check for supported ecosystems
	ecosystemLower := strings.ToLower(ecosystem)
	switch ecosystemLower {
	case "npm":
		return true
	case "pypi":
		return true
	default:
		return false
	}
}
