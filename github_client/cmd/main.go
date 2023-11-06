package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type UserData struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type RepoData struct {
	Name      string `json:"name"`
	ForkCount int    `json:"fork_count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	UserData             UserData
	Repos                []RepoData
	LanguageDistribution map[string]rune
}

func FetchGithubData[ReturnType UserData | []RepoData | map[string]interface{}](client *http.Client, request *http.Request) ReturnType {
	res, err := client.Do(request)
	if err != nil {
		_ = fmt.Errorf("error fetching: %v\n", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	var data ReturnType
	body, _ := io.ReadAll(res.Body)
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		_ = fmt.Errorf("error json parsing: %v\n", jsonErr)
	}
	return data
}

func getUserData(username string) {
	client := http.Client{
		Timeout: time.Second * 2,
	}
	user := User{}
	userRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%s", username), nil)
	reposRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%s/repos", username), nil)
	user.UserData = FetchGithubData[UserData](&client, userRequest)
	user.Repos = FetchGithubData[[]RepoData](&client, reposRequest)

	var repoNamesList []string
	for _, repoData := range user.Repos {
		repoNamesList = append(repoNamesList, repoData.Name)
	}

	langDistribution := make(map[string]rune)
	for _, repoName := range repoNamesList {
		languageRequest, _ := http.NewRequest(http.MethodGet,
			fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", username, repoName),
			nil)

		repoLangUsage := FetchGithubData[map[string]interface{}](&client, languageRequest)
		fmt.Println(repoLangUsage)
		for lang, val := range repoLangUsage {
			if v, ok := val.(rune); ok {
				langDistribution[lang] += v
			}
		}
	}
	var totalLines rune
	for _, lines := range langDistribution {
		totalLines += lines
	}
	for lang, val := range langDistribution {
		langDistribution[lang] = val * 100.0 / totalLines
	}
	user.LanguageDistribution = langDistribution
	fmt.Println(user)
}

func readUsernames(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := scanner.Text()

		getUserData(username)
	}
}
func main() {
	if len(os.Args) < 2 {
		log.Fatal("No filenames in the cli arguments")
	}
	filename := os.Args[1]
	pwd, _ := os.Getwd()
	open, err := os.Open(fmt.Sprintf("%s\\%s", pwd, filename))
	if err != nil {
		_ = fmt.Errorf("error opening filename %s: %v", filename, err)
		return
	}
	readUsernames(open)

	defer open.Close()

}
