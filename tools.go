package tools

import (
	"math/rand"	
	"io/ioutil"
	"strings"
	"errors"
	"regexp"
	"bufio"
	"bytes"
	"time"
	"fmt"
	"os"
	"io"
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

	buf := make([]byte, 1024)
	fileLines := 1

	for {
		readBytes, err := fileHandle.Read(buf)
		if err != nil {
			if readBytes == 0 && err == io.EOF {
				break
			}
			return 0, err
		}

		fileLines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}


	return fileLines, nil
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
		synonym, done, err := GetStringInBetween(byteToString, start, end)
		if done{
			break
		}
		if err != nil{
			return "", err
		}

		// Split and add make list
		synonymSplited := strings.Split(synonym, betweenWords)

		// Choose synonym
		byteToString = strings.Replace(byteToString, fmt.Sprint(start, synonym, end), synonymSplited[Random(0, cap(synonymSplited)-1)], 1)
	}

	return byteToString, nil
}

// Synonym return combine text from string
func StringMix(text string, start string, end string, betweenWords string) (string, error){
	for {
		// Too fast, too furious
		time.Sleep(time.Microsecond * 100)
		
		// Getting words between brackets
		synonym, done, err := GetStringInBetween(text, start, end)
		if done{
			break
		}
		if err != nil{
			return "", err
		}

		// Split and add make list
		synonymSplited := strings.Split(synonym, betweenWords)

		// Choose synonym
		text = strings.Replace(text, fmt.Sprint(start, synonym, end), synonymSplited[Random(0, cap(synonymSplited)-1)], 1)
	}

	return text, nil
}

// FileMixRegExp return combine text from file
func FileMixRegExp(filePath string, start string, end string, betweenWords string) (string, error) {
	fileHandle, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert bytes to strings
	byteToString := string(fileHandle)

	pattern := fmt.Sprintf("%s(.*?)%s", start, end)
	re := regexp.MustCompile(pattern)
	results := re.FindAllString(byteToString, -1)
	for _, result := range results {
		result = result[1:len(result)-1]
	    synonymSplited := strings.Split(result, betweenWords)
		byteToString = strings.Replace(byteToString, fmt.Sprint(start, result, end), synonymSplited[Random(0, cap(synonymSplited)-1)], 1)
	}
	return byteToString, nil
}

// StringMixWithRegExp works better with 
func StringMixWithRegExp(text string, start string, end string, betweenWords string) string{
	pattern := fmt.Sprintf("%s(.*?)%s", start, end)
	re := regexp.MustCompile(pattern)
	results := re.FindAllString(text, -1)
	for _, result := range results {
		result = result[1:len(result)-1]
	    synonymSplited := strings.Split(result, betweenWords)
		text = strings.Replace(text, fmt.Sprint(start, result, end), synonymSplited[Random(0, cap(synonymSplited)-1)], 1)
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
func GetStringInBetween(str string, homeChar string, endChar string) (string, bool, error) {
	home := strings.Index(str, homeChar)
	if home == -1 {
		return "", true, errors.New("No strings to find")
	}

	home += len(homeChar)
	end := strings.Index(str, endChar)
	if home > end{
		return "", false, errors.New("Syntax error with mixing text")
	}

	return str[home:end], false, nil
}

// Random return random numbers in rage
func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// RandomFileLine return random line in file
func RandomFileLine(filePath string) (string, error){
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	// Closing the file at the end of the function
	defer fileHandle.Close()

	buf := make([]byte, 1024)
	fileLines := 0

	for {
		readBytes, err := fileHandle.Read(buf)
		if err != nil {
			if readBytes == 0 && err == io.EOF {
				break
			}
			return "", err
		}

		fileLines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}


	randomLine := Random(0, fileLines)
	actualLine := 0

	// Making loop for file fileLines
	fileHandle.Seek(0, io.SeekStart)
	fileHandleScan := bufio.NewScanner(fileHandle)
	for fileHandleScan.Scan() {
		if actualLine >= randomLine {
			// Return result with removed all whitespaces and suffixes
			if fileLines == 0 && fileHandleScan.Text() == "" {
				return "", errors.New("File is empty")
			}
			return strings.TrimSpace(fileHandleScan.Text()), nil
		}

		actualLine++
	}

	return "", errors.New("Func had problem with get random line")
}