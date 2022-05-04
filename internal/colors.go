package internal

import (
	"fmt"
	"image"

	"github.com/emirpasic/gods/sets/hashset"
)

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
