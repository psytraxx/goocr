package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func main() {

	pwd, err := os.Getwd()
	check(err)
	folder := pwd + "/../../testimage/"
	files, err := getTestFiles(folder)
	check(err)

	for _, fileName := range files {
		fileName := fmt.Sprintf("%s%s", folder, fileName.Name())
		dat, err := os.ReadFile(fileName)
		check(err)
		fmt.Println(fileName)
		extractText(dat)

		//img, t, err := image.Decode(bytes.NewReader(dat))
		//check(err)
		//fmt.Println(t)
		//fmt.Println(img)
	}

}

func extractText(dat []byte) {
	client := getClient()
	defer client.Close()

	client.SetImageFromBytes(dat)

	boxes, err := client.GetBoundingBoxes(gosseract.RIL_TEXTLINE)
	check(err)
	for _, v := range boxes {
		if v.Confidence > 75 {
			fmt.Println(v.Word)
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

func getClient() (client *gosseract.Client) {

	client = gosseract.NewClient()
	client.SetPageSegMode(gosseract.PSM_AUTO_ONLY)
	client.SetVariable("hocr_char_boxes", "1")
	client.Languages = append(client.Languages, "deu")
	client.Languages = append(client.Languages, "fra")
	client.Languages = append(client.Languages, "ita")
	client.Languages = append(client.Languages, "eng")
	return
}

/*
func saveFrames(imgByte []byte) {
	img, _, _ := image.Decode(bytes.NewReader(imgByte))
	out, err := os.Create("./img.jpeg")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = jpeg.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
*/
