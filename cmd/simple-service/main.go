package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	client = GetInstance()

	pwd, err := os.Getwd()
	check(err)
	folder := pwd + "/../../testimage/"
	files, err := getTestFiles(folder)
	check(err)

	fileName := fmt.Sprintf("%s%s", folder, files[2].Name())
	//dat, err := os.ReadFile(fileName)
	check(err)
	client.SetImage("/home/efunk/workspace/goocr/testimage/screengrab.jpg")
	fmt.Println(fileName)
	text, err := client.Text()
	check(err)
	fmt.Println(text)
}

func getTestFiles(path string) (files []fs.FileInfo, err error) {

	files, err = ioutil.ReadDir(path)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
