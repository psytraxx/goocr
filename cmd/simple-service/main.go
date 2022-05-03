package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	client.Languages = append(client.Languages, "deu")
	client.Languages = append(client.Languages, "eng")
	client.Languages = append(client.Languages, "fra")
	client.Languages = append(client.Languages, "ita")
	defer client.Close()
	client.SetImage("/home/eric/workspace/goocr/testimage/screengrab.jpg")
	text, err := client.Text()
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println(text)
}
