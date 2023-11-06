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

type UnknownJSON struct {
	Data map[string]interface{}
}

func fetchUserData(client *http.Client, request *http.Request) UserData {
	res, err := client.Do(request)
	if err != nil {
		_ = fmt.Errorf("error fetching for user: %v\n", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	userdata := UserData{}
	body, _ := io.ReadAll(res.Body)
	jsonErr := json.Unmarshal(body, &userdata)
	if jsonErr != nil {
		_ = fmt.Errorf("error json parsing: %v\n", jsonErr)
	}
	return userdata
}
func fetchRepoData(client *http.Client, request *http.Request) []RepoData {
	res, err := client.Do(request)
	if err != nil {
		_ = fmt.Errorf("error fetching repos: %v\n", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	var repos []RepoData
	body, _ := io.ReadAll(res.Body)
	jsonErr := json.Unmarshal(body, &repos)
	if jsonErr != nil {
		_ = fmt.Errorf("error json parsing: %v\n", jsonErr)
	}
	return repos
}

func fetchRepoLanguageUsage(client *http.Client, request *http.Request) map[string]interface{} {
	res, err := client.Do(request)
	if err != nil {
		_ = fmt.Errorf("error fetching languages: %v\n", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	languages := make(map[string]interface{})
	body, _ := io.ReadAll(res.Body)
	jsonErr := json.Unmarshal(body, &languages)
	if jsonErr != nil {
		_ = fmt.Errorf("error json parsing: %v\n", jsonErr)
	}
	fmt.Println(languages)
	return languages
}

func getUserData(username string) {
	client := http.Client{
		Timeout: time.Second * 2,
	}
	user := User{}
	userRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%s", username), nil)
	reposRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%s/repos", username), nil)
	user.UserData = fetchUserData(&client, userRequest)
	user.Repos = fetchRepoData(&client, reposRequest)
	var repoNamesList []string
	for _, repoData := range user.Repos {
		repoNamesList = append(repoNamesList, repoData.Name)
	}
	//langDistribution := make(map[string]interface{})
	for _, repoName := range repoNamesList {
		languageRequest, _ := http.NewRequest(http.MethodGet,
			fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", username, repoName),
			nil)

		repoLangUsage := fetchRepoLanguageUsage(&client, languageRequest)
		fmt.Println(repoLangUsage)
		//for lang, val := range repoLangUsage {
		//	langDistribution[lang] += rune(val)
		//}
	}
	//var totalLines rune
	//for _, lines := range langDistribution {
	//	totalLines += lines
	//}
	//for lang, val := range langDistribution {
	//	langDistribution[lang] = val * 100.0 / totalLines
	//}
	//user.LanguageDistribution = langDistribution
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

	open, err := os.Open(fmt.Sprintf("%s\\hw_github_client\\cmd\\github_client\\%s", pwd, filename))
	if err != nil {
		_ = fmt.Errorf("error opening filename %s: %v", filename, err)
		return
	}
	readUsernames(open)

	defer open.Close()

}
