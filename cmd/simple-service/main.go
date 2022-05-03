package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/otiai10/gosseract/v2"
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
			if !Opaque(imageData) {
				fmt.Println("File should be stored as PNG because it has an alpha channel")
				continue
			}
		}

		ExtractHistogram(imageData)

		nr := CountColors(imageData)

		if nr < 1024 {
			fmt.Println("File should be stored as PNG because it has less than 1024 colors")
			continue
		}

		out := ExtractText(fileName)
		if len(out) > 0 {
			fmt.Println("File should be stored as PNG because it contains text")
			continue
		}

	}

}

func CountColors(imageData image.Image) int {
	bounds := imageData.Bounds()

	set := hashset.New()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			set.Add(imageData.At(x, y))
		}
	}

	return len(set.Values())
}

func ExtractHistogram(imageData image.Image) {
	bounds := imageData.Bounds()
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := imageData.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}

func ExtractText(fileName string) (out []string) {
	client := getClient()
	defer client.Close()

	client.SetImage(fileName)

	boxes, err := client.GetBoundingBoxes(gosseract.RIL_TEXTLINE)
	check(err)
	for _, v := range boxes {
		if v.Confidence > 75 {
			out = append(out, v.Word)
		}
	}
	return
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
	client.Languages = append(client.Languages, "deu")
	client.Languages = append(client.Languages, "fra")
	client.Languages = append(client.Languages, "ita")
	client.Languages = append(client.Languages, "eng")
	return
}

func Opaque(im image.Image) bool {
	// Check if image has Opaque() method:
	if oim, ok := im.(interface {
		Opaque() bool
	}); ok {
		return oim.Opaque() // It does, call it and return its result!
	}

	// No Opaque() method, we need to loop through all pixels and check manually:
	rect := im.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if _, _, _, a := im.At(x, y).RGBA(); a != 0xffff {
				return false // Found a non-opaque pixel: image is non-opaque
			}
		}

	}
	return true // All pixels are opaque, so is the image
}
