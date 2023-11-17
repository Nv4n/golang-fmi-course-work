package ghclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type UserData struct {
	Username    string `json:"login"`
	ReposApiURL string `json:"repos_url"`
	Followers   int    `json:"followers"`
}

type RepoData struct {
	Name            string    `json:"name"`
	LanguagesApiURL string    `json:"languages_url"`
	ForksCount      int       `json:"forks_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type LanguageDistribution map[string]float64

type UserFullData struct {
	UserData             UserData
	Repos                []RepoData
	LanguageDistribution LanguageDistribution
}

type UserFormatedData struct {
	Username             string
	Followers            int
	ForksCount           int
	LanguageDistribution LanguageDistribution
}

func fetchGithubData[ReturnType UserData | []RepoData | map[string]interface{}](client *http.Client, request *http.Request) ReturnType {
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

func getLanguageApiURLs(repos []RepoData, repoLimit int) []string {
	var langApiList []string
	for i, repoData := range repos {
		if i >= repoLimit {
			break
		}
		langApiList = append(langApiList, repoData.LanguagesApiURL)
	}
	return langApiList
}

func calcLangDistribution(distribution map[string]int64) LanguageDistribution {
	var totalLines int64
	langDistribution := make(map[string]float64)
	for _, lines := range distribution {
		totalLines += lines
	}
	for lang, val := range distribution {
		langDistribution[lang] = float64(val) * 100.0 / float64(totalLines)
	}
	return langDistribution
}

func calcTotalForksCount(repos []RepoData) int {
	count := 0
	for _, r := range repos {
		count = count + r.ForksCount
	}
	return count
}

func GetUserData(username string, repoLimit int) UserFormatedData {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	user := UserFullData{}
	userRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%s", username), nil)
	user.UserData = fetchGithubData[UserData](&client, userRequest)
	reposRequest, _ := http.NewRequest(http.MethodGet, user.UserData.ReposApiURL, nil)
	user.Repos = fetchGithubData[[]RepoData](&client, reposRequest)

	if repoLimit == -1 {
		repoLimit = len(user.Repos)
	}

	langApiList := getLanguageApiURLs(user.Repos, repoLimit)

	languageBitList := make(map[string]int64)
	for _, url := range langApiList {
		languageRequest, _ := http.NewRequest(http.MethodGet, url, nil)

		repoLangUsage := fetchGithubData[map[string]interface{}](&client, languageRequest)
		for lang, val := range repoLangUsage {
			if v, ok := val.(int64); ok {
				languageBitList[lang] += v
			}
		}
	}

	user.LanguageDistribution = calcLangDistribution(languageBitList)
	totalForkCount := calcTotalForksCount(user.Repos)

	return UserFormatedData{
		user.UserData.Username,
		user.UserData.Followers,
		totalForkCount,
		user.LanguageDistribution,
	}
}
