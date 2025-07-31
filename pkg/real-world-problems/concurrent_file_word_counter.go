package real_world_problems

import (
"bufio"
"fmt"
"io/fs"
"os"
"path/filepath"
"strings"
"sync"
)

// SafeCounter to store word counts safely across goroutines
type SafeCounter struct {
	mu    sync.Mutex
	count map[string]int
}

func (sc *SafeCounter) Add(word string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count[word]++
}

func processFile(path string, sc *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", path, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		word = strings.Trim(word, ".,!?:;\"'()[]{}") // clean punctuation
		if word != "" {
			sc.Add(word)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", path, err)
	}
}

func main() {
	dirPath := "./texts" // Change this to your directory path
	sc := &SafeCounter{count: make(map[string]int)}
	var wg sync.WaitGroup

	// Walk through directory to read all .txt files
	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".txt") {
			wg.Add(1)
			go processFile(path, sc, &wg)
		}
		return nil
	})

	wg.Wait()

	// Print top 5 frequent words
	fmt.Println("\n--- Word Frequency ---")
	type kv struct {
		key   string
		value int
	}
	var wordList []kv
	for k, v := range sc.count {
		wordList = append(wordList, kv{k, v})
	}

	// Sort by frequency (simple bubble sort for interview clarity)
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if wordList[j].value > wordList[i].value {
				wordList[i], wordList[j] = wordList[j], wordList[i]
			}
		}
	}

	top := 5
	if len(wordList) < 5 {
		top = len(wordList)
	}
	for i := 0; i < top; i++ {
		fmt.Printf("%s: %d\n", wordList[i].key, wordList[i].value)
	}
}

