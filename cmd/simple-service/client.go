package main

import (
	"fmt"
	"sync"

	"github.com/otiai10/gosseract/v2"
)

var lock = &sync.Mutex{}

var client *gosseract.Client

func GetInstance() *gosseract.Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			fmt.Println("Creating single instance now.")
			client := gosseract.NewClient()
			client.Languages = append(client.Languages, "deu")
			client.Languages = append(client.Languages, "eng")
			client.Languages = append(client.Languages, "fra")
			client.Languages = append(client.Languages, "ita")
			defer client.Close()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return client
}
