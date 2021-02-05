package tools

import (
	"math/rand"	
	"io/ioutil"
	"strings"
	"errors"
	"bufio"
	"time"
	"fmt"
	"os"
)

// Flush append data to file
func Flush(filePath string, str string) error {
	fileHandle, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		// If file doesn't exist he create him
		if os.IsNotExist(err) {
			fileHandle, err = os.Create(filePath)
			if err != nil{
				return err
			}
		} else{
			return err
		}

	}

	// Closing the file at the end of the function
	defer fileHandle.Close()

	// Flush data to file
	fileHandle.WriteString(str)

	return nil
}

// FileContains check that string exist in file
func FileContains(filePath string, str string) (bool, error) {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	// Closing the file at the end of the function
	defer fileHandle.Close()

	// Making loop for file lines
	fileHandleScan := bufio.NewScanner(fileHandle)
	for fileHandleScan.Scan() {
		if strings.Contains(fileHandleScan.Text(), str) {
			return true, nil
		}
	}

	return false, nil
}

// FileLinesCounter counts lines of file
func FileLinesCounter(filePath string) (int, error) {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}

	// Closing the file at the end of the function
	defer fileHandle.Close()

	// Preparing variable for counting lines
	var lines int

	// Making loop for file lines
	fileHandleScan := bufio.NewScanner(fileHandle)
	for fileHandleScan.Scan() {
		lines++
	}

	return lines, err
}

// SpecifyLineByNumber return specific text of line in file
func SpecifyLineByNumber(filePath string, specifyLine int) (string, error) {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	// Closing the file at the end of the function
	defer fileHandle.Close()

	// Preparing variable for result
	var lines int

	// Making loop for file lines
	fileHandleScan := bufio.NewScanner(fileHandle)
	for fileHandleScan.Scan() {
		if lines <= specifyLine{
			// Return result with removed all whitespaces and suffixes
			return fmt.Sprint(strings.TrimSpace(fileHandleScan.Text())), nil
		}

		lines++
	}

	return "", errors.New("The line cannot be found")
}

// SpecifyLineByText return specific text of line in file
func SpecifyLineByText(filePath string, str string) (string, error) {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	// Closing the file at the end of the function
	defer fileHandle.Close()

	// Making loop for file lines
	fileHandleScan := bufio.NewScanner(fileHandle)
	for fileHandleScan.Scan() {
		// Return result with removed all whitespaces and suffixes
		if strings.Contains(fileHandleScan.Text(), str) {
			return fmt.Sprint(strings.TrimSpace(fileHandleScan.Text())), nil
		}
	}

	return "", errors.New("The line cannot be found")
}

// Synonym return combine text from file
func FileMix(filePath string, start string, end string, betweenWords string) (string, error) {
	fileHandle, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert bytes to strings
	byteToString := string(fileHandle)

	// Make loop for choose synonyms
	for {
		// Too fast, too furious
		time.Sleep(time.Microsecond * 100)

		// Getting words between brackets
		synonym, err := GetStringInBetween(byteToString, start, end)
		if err != nil {
			break
		}

		// Split and add make list
		synonymSplited := strings.Split(synonym, betweenWords)

		// Choose synonym
		byteToString = strings.Replace(byteToString, fmt.Sprint(start, synonym, end), synonymSplited[Random(0, cap(synonymSplited)-1)], 1)
	}

	return byteToString, nil
}

// Synonym return combine text from string
func StringMix(text string, start string, end string, betweenWords string) string{
	for {
		// Too fast, too furious
		time.Sleep(time.Microsecond * 100)
		
		// Getting words between brackets
		synonym, err := GetStringInBetween(text, start, end)
		if err != nil {
			break
		}

		// Split and add make list
		synonymSplited := strings.Split(synonym, betweenWords)

		// Choose synonym
		text = strings.Replace(text, fmt.Sprint(start, synonym, end), synonymSplited[Random(0, cap(synonymSplited)-1)], 1)
	}

	return text
}

// Exists check a file exists
func Exists(filePath string) (bool, error) {
	 _, err := os.Stat(filePath)
	 if err != nil{
	 	// Not exist
	 	if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	 }

	// Exist
	return true, nil
}

// GetStringInBetween return string between two characters
func GetStringInBetween(str string, homeChar string, endChar string) (string, error) {
	home := strings.Index(str, homeChar)
	if home == -1 {
		return "", errors.New("No strings to find")
	}
	home += len(homeChar)

	end := strings.Index(str, endChar)

	return str[home:end], nil
}

// Random return random numbers in rage
func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}