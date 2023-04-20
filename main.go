package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type PageData struct {
	Title  string
	Body   string
	Images []string
}

type GithubTreeResponse struct {
	Tree []struct {
		Path string `json:"path"`
	} `json:"tree"`
}

func main() {
	file, err := os.Open("./public/index.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	htmlBytes, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	htmlString := string(htmlBytes)

	url := "https://api.github.com/repos/keidarcy/images-host/git/trees/master?recursive=1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while calling GitHub API:", err)
		return
	}
	defer resp.Body.Close()

	var tree GithubTreeResponse
	err = json.NewDecoder(resp.Body).Decode(&tree)
	if err != nil {
		fmt.Println("Error while decoding response:", err)
		return
	}

	var jpgPaths []string
	for _, node := range tree.Tree {
		if strings.HasSuffix(node.Path, ".jpg") {
			jpgPaths = append(jpgPaths, "https://keidarcy.github.io/images-host/"+node.Path)
		}
	}

	title := "My Page"
	data := PageData{
		Title:  title,
		Body:   "<p>Hello, world!</p>",
		Images: jpgPaths,
	}
	tmpl := template.Must(template.New("html").Parse(htmlString))

	newFile, _ := os.Create("output.html")

	err = tmpl.Execute(newFile, data)

	if err != nil {
		panic(err)
	}

}
