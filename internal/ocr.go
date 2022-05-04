package internal

import "github.com/otiai10/gosseract/v2"

func ExtractText(fileName string) (out []string, err error) {
	client := getClient()
	defer client.Close()

	client.SetImage(fileName)

	boxes, err := client.GetBoundingBoxes(gosseract.RIL_TEXTLINE)
	if err != nil {
		return
	}
	for _, v := range boxes {
		if v.Confidence > 75 {
			out = append(out, v.Word)
		}
	}
	return
}

func getClient() (client *gosseract.Client) {

	client = gosseract.NewClient()
	client.Languages = append(client.Languages, "deu")
	client.Languages = append(client.Languages, "fra")
	client.Languages = append(client.Languages, "ita")
	client.Languages = append(client.Languages, "eng")
	return
}
