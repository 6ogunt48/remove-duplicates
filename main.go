package main

import (
	"bufio"
	"fmt"
	"os"
)

// ANSI color code for red
const redColor = "\033[31m"

func main() {
	inputFile := "input.txt"   //
	outputFile := "output.txt" //

	emails, err := readEmails(inputFile)
	if err != nil {
		panic(err)
	}

	uniqueEmails, duplicates := processEmails(emails)

	if err := writeEmails(outputFile, uniqueEmails); err != nil {
		panic(err)
	}

	// Log the duplicates in red
	for _, email := range duplicates {
		fmt.Printf("%sDuplicate removed: %s\033[0m\n", redColor, email)
	}
	fmt.Printf("Total duplicates removed: %d\n", len(duplicates))
}

// readEmails reads emails from a file and returns a slice of emails
func readEmails(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var emails []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, scanner.Text())
	}
	return emails, scanner.Err()
}

// processEmails removes duplicates and returns a slice of unique emails and a slice of duplicates
func processEmails(emails []string) ([]string, []string) {
	emailMap := make(map[string]bool)
	var uniqueEmails []string
	var duplicates []string

	for _, email := range emails {
		if _, exists := emailMap[email]; !exists {
			emailMap[email] = true
			uniqueEmails = append(uniqueEmails, email)
		} else {
			duplicates = append(duplicates, email)
		}
	}
	return uniqueEmails, duplicates
}

// writeEmails writes a slice of emails to a file
func writeEmails(filename string, emails []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, email := range emails {
		if _, err := writer.WriteString(email + "\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}
