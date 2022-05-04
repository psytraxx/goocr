package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/erixfunktxgroup/goocr/internal"
)

func main() {

	folder := "/home/efunk/workspace/goocr/testimage/"
	files, err := getTestFiles(folder)
	check(err)

	for _, fileName := range files {
		fileName := fmt.Sprintf("%s%s", folder, fileName.Name())
		fmt.Println(fileName)

		file, err := os.Open(fileName)
		check(err)
		defer file.Close()

		imageData, _, err := image.Decode(file)

		if err == nil {
			if !internal.Opaque(imageData) {
				fmt.Println("File should be stored as PNG because it has an alpha channel")
				continue
			}
		}

		internal.ExtractHistogram(imageData)

		nr := internal.CountColors(imageData)

		if nr < 1024 {
			fmt.Println("File should be stored as PNG because it has less than 1024 colors")
			continue
		}

		out, err := internal.ExtractText(fileName)
		check(err)
		if len(out) > 0 {
			fmt.Println("File should be stored as PNG because it contains text")
			continue
		}

	}

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
