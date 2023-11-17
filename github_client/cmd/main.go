package main

import (
	"bufio"
	"fmt"
	gh "ghclient_homework/ghclient"
	"log"
	"os"
)

func readUsernames(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := scanner.Text()

		gh.GetUserData(username)
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
