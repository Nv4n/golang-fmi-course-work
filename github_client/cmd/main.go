package main

import (
	"bufio"
	"fmt"
	gh "ghclient_homework/ghclient"
	"log"
	"os"

	"github.com/rodaine/table"
)

func getUsernames(file *os.File) []string {
	var usernames []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := scanner.Text()
		usernames = append(usernames, username)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during reading username: %v\n", err)
	}
	return usernames
}

func presentGhData(users []gh.UserFormattedData) {
	languages := make(map[string]struct{})
	for _, u := range users {
		for l, _ := range u.LanguageDistribution {
			languages[l] = struct{}{}
		}
	}
	var languageColumns []string
	for l, _ := range languages {
		languageColumns = append(languageColumns, l)
	}

	columns := []string{"Username", "Followers", "Forks"}
	columns = append(columns, languageColumns...)
	tbl := table.New(columns)

	for _, u := range users {
		var languageDistribution []interface{}
		for _, dist := range u.LanguageDistribution {
			languageDistribution = append(languageDistribution, fmt.Sprintf("%.2f", dist))
		}
		data := []interface{}{u.Username, fmt.Sprintf("%d", u.Followers), fmt.Sprintf("%d", u.ForksCount)}
		//fmt.Println(data)
		data = append(data, languageDistribution...)
		tbl.AddRow(data...)
	}

	tbl.Print()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No filenames in the cli arguments")
	}
	filename := os.Args[1]
	pwd, _ := os.Getwd()
	open, err := os.Open(fmt.Sprintf("%s\\cmd\\%s", pwd, filename))
	if err != nil {
		_ = fmt.Errorf("error opening filename %s: %v", filename, err)
		return
	}
	usernames := getUsernames(open)
	var users []gh.UserFormattedData
	for _, u := range usernames {
		data := gh.GetUserData(u, 10)
		users = append(users, data)
	}
	presentGhData(users)

	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(open)

}
