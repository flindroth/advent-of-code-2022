package util

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseUrl    = "https://adventofcode.com"
	year       = 2022
	cookieFile = "session.txt"
)

func GetPuzzleInput(day int) ([]string, error) {
	sessionCookie, err := getSessionCookie()
	if err != nil {
		return nil, fmt.Errorf("cannot read session cookie: %v", err.Error())
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v/day/%v/input", baseUrl, year, day), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

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
