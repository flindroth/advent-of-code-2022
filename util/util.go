package util

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	baseUrl    = "https://adventofcode.com"
	year       = 2022
	cookieFile = "session.txt"
	userAgent  = "github.com/flindroth/advent-of-code-2022 fredriklindroth@gmail.com"
)

func GetPuzzleInput(day int) ([]string, error) {
	lines, err := getPuzzleInputFromCache(day)
	if err != nil {
		log.Printf("Cache miss, fetching puzzle input by HTTP")
		return getPuzzleInputFromHTTP(day)
	}
	return lines, nil
}

func getPuzzleInputFromCache(day int) ([]string, error) {

	file, err := os.Open(fmt.Sprintf(".cache/days/%v", day))
	if err != nil {
		return nil, err
	}

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func getPuzzleInputFromHTTP(day int) ([]string, error) {
	sessionCookie, err := getSessionCookie()
	if err != nil {
		return nil, fmt.Errorf("cannot read session cookie: %v", err.Error())
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v/day/%v/input", baseUrl, year, day), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})
	req.Header.Add("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != 200 {
		printError(resp.Body)
		return nil, fmt.Errorf("status %v when getting puzzle input", resp.StatusCode)
	}

	lines := make([]string, 0)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := saveLinesToCache(day, lines); err != nil {
		log.Printf("WARNING: Unable to save puzzle input to cache: %v", err.Error())
	}
	return lines, nil
}

func printError(reader io.Reader) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	log.Println(string(body))
}

func getSessionCookie() (string, error) {
	bytes, err := ioutil.ReadFile(cookieFile)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func saveLinesToCache(day int, lines []string) error {
	if err := os.MkdirAll(".cache/days", fs.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf(".cache/days/%v", day))
	if err != nil {
		return err
	}

	if file != nil {
		defer file.Close()
	}

	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	return nil
}
