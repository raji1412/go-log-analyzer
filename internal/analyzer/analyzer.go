package analyzer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// readFile - reads the files data and return array of lines
func readFile(filePaths string) ([]string, error) {
	filePathArray := strings.Split(filePaths, ",")
	if len(filePathArray) == 0 {
		log.Println("E! [readFile] Split- empty", filePaths)
	}
	dataArray := []string{}
	for _, filePath := range filePathArray {
		fileData, err := os.Open(filePath)
		if err != nil {
			log.Println("E! [logCounter] ReadFile: ", err, filePath)
			return nil, err
		}

		scanner := bufio.NewScanner(fileData)
		if scanner == nil {
			return nil, fmt.Errorf("scanner is nil")
		}
		for scanner.Scan() {
			data := scanner.Text()
			dataArray = append(dataArray, data)
		}
		fileData.Close()
	}
	return dataArray, nil

}

// LogCounter - counts the different log types
func LogCounter(filePath string, level string) map[string]int {
	result := map[string]int{}

	readerArray, err := readFile(filePath)
	if err != nil {
		log.Println("E! [logCounter] ReadFile: ", err, filePath)
		return result
	}
	level = strings.ToLower(level)
	for _, data := range readerArray {
		if data != "" {
			arrayData := strings.Fields(data)
			if len(arrayData) > 0 {
				levelData := strings.ToLower(arrayData[0])
				if level == "" || levelData == level {
					result[levelData]++
				}

			}

		}
	}

	return result
}

// ParseResult - Print the result
func ParseResult(result map[string]int) {
	if len(result) == 0 {
		log.Println("W! [ParseResult] No logs of the selected level type")
	}
	var countResults []LogCount
	for level, count := range result {
		countResults = append(countResults, LogCount{level, count})
	}
	sort.Slice(countResults, func(i, j int) bool {
		return countResults[i].Count > countResults[j].Count
	})
	fmt.Println("LOG LEVEL   COUNT")
	fmt.Println("------------------")
	for _, v := range countResults {
		fmt.Printf("%-15s %d\n", strings.ToTitle(v.Level), v.Count)
	}

}
