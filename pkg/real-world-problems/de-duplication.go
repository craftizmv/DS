package real_world_problems

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

const chunkSize = 1000000 // 1M lines read per chunk.

func main() {
	inputFile := "large_image_ids.txt"
	outputFile := "deduplicated_ids.txt"

	// 1️⃣ Step 1: Read file in chunks and write sorted unique IDs into temp files
	tempFiles := processChunks(inputFile, chunkSize)

	// 2️⃣ Step 2: Merge all sorted chunks into a single deduplicated output file
	mergeChunks(tempFiles, outputFile)

	// 3️⃣ Step 3: Clean up temporary chunk files to free disk space
	for _, f := range tempFiles {
		os.Remove(f)
	}

	fmt.Println("✅ Deduplication complete:", outputFile)
}

// processChunks reads the input file in fixed-size chunks, deduplicates each chunk,
// sorts it, and writes it to a temporary file.
// returns the list of paths to those chunked files.
func processChunks(inputFile string, chunkSize int) []string {
	var tempFiles []string

	// 1. Open file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		e := file.Close()
		if e != nil {
			log.Fatal(e)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, chunkSize)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		if len(lines) >= chunkSize {
			tempFiles = append(tempFiles, writeUniqueChunk(lines))
			// reset the slice to use the same underlying array.
			lines = lines[:0]
		}
	}

	// cover the condition of partial read if data read is lesser than chunk size.
	if len(lines) > 0 {
		tempFiles = append(tempFiles, writeUniqueChunk(lines))
	}

	return tempFiles
}

func writeUniqueChunk(lines []string) string {
	set := make(map[string]struct{})

	for _, l := range lines {
		// TODO : Check if we could have created a hash of it.
		set[strings.TrimSpace(l)] = struct{}{}
	}

	unique := make([]string, 0, len(set))
	for key := range set {
		unique = append(unique, key)
	}

	sort.Strings(unique)

	tempFile := fmt.Sprintf("chunk_%d.txt", time.Now().UnixNano())
	f, _ := os.Create(tempFile)
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, id := range unique {
		writer.WriteString(id + "\n")
	}
	writer.Flush()
	return tempFile
}

// mergeChunks merges all temporary sorted files while removing duplicates
func mergeChunks(tempFilePaths []string, outputFilePath string)
