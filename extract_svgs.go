package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	inputFile := "DEFAULT_THEMES.md"
	outputFile := "DEFAULT_THEMES_updated.md"
	outputDir := "themes"
	// Matches: ![label](data:image/svg+xml,...) (the first markdown image tag in the table cell)
	svgPattern := regexp.MustCompile(`!\[([^\]]+)\]\(data:image/svg\+xml,[^)]+\)`)

	// Ensure output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		// Find SVG markdown image
		matches := svgPattern.FindStringSubmatch(line)
		if len(matches) == 2 {
			fullMatch := svgPattern.FindString(line) // the whole ![...]() tag
			name := matches[1]

			// Extract SVG data URI
			// Capture inside parens
			uriPattern := regexp.MustCompile(`\((data:image/svg\+xml,[^)]+)\)`)
			uriMatch := uriPattern.FindStringSubmatch(fullMatch)
			if len(uriMatch) == 2 {
				dataURI := uriMatch[1]
				// Remove prefix:
				svgRaw := strings.TrimPrefix(dataURI, "data:image/svg+xml,")
				// Decode percent escapes
				svgDecoded, err := url.QueryUnescape(svgRaw)
				if err != nil {
					fmt.Printf("Error decoding SVG for %s: %v\n", name, err)
					outFile.WriteString(line + "\n")
					continue
				}
				// Sanitize file name
				filename := strings.ToLower(name)
				filename = strings.ReplaceAll(filename, " ", "_")
				filename = regexp.MustCompile(`[^\w\-_]`).ReplaceAllString(filename, "")
				svgPath := fmt.Sprintf("%s/%s.svg", outputDir, filename)

				// Write the SVG file
				err = ioutil.WriteFile(svgPath, []byte(svgDecoded), 0644)
				if err != nil {
					fmt.Printf("Error writing SVG for %s: %v\n", name, err)
					outFile.WriteString(line + "\n")
					continue
				}
				fmt.Printf("Wrote %s\n", svgPath)

				// Replace the SVG cell with a clean image link
				// Find start/end index of the cell (first | ... |)
				firstPipe := strings.Index(line, "|")
				secondPipe := strings.Index(line[firstPipe+1:], "|")
				if secondPipe >= 0 {
					// Compose a clean cell with our new svg file link
					newCell := fmt.Sprintf(" ![%s](%s) ", name, svgPath)
					// Replace between first and second pipe
					newLine := line[:firstPipe+1] + newCell + line[firstPipe+1+secondPipe:]
					outFile.WriteString(newLine + "\n")
				} else {
					// fallback: just swap image tag for new link
					newLine := strings.Replace(line, fullMatch, fmt.Sprintf("![%s](%s)", name, svgPath), 1)
					outFile.WriteString(newLine + "\n")
				}
			} else {
				// If no data URI, just output line as-is
				outFile.WriteString(line + "\n")
			}
		} else {
			outFile.WriteString(line + "\n")
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("Finished! Updated markdown written to %s\n", outputFile)
}
