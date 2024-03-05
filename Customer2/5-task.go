package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Info struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func jsonniolish() ([]Info, error) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/info")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var info []Info
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func writeToJSONFile(fileName string, info []Info, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("xato fayle yaratishda:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		fmt.Println("xato marshiling qilishda:", err)
		return
	}

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("xato faylga yozishda:", err)
		return
	}

	fmt.Println("ma'lumot yozildi: ", fileName)
}

func main() {
	info, err := jsonniolish()
	if err != nil {
		fmt.Println("xato jsonni olishda:", err)
		return
	}

	numFiles := 5
	uzunligi := (len(info) + numFiles - 1) / numFiles

	var wg sync.WaitGroup
	for i := 0; i < numFiles; i++ {
		start := i * uzunligi
		end := (i + 1) * uzunligi
		if end > len(info) {
			end = len(info)
		}

		wg.Add(1)
		go writeToJSONFile(fmt.Sprintf("fayl_%d.json", i+1), info[start:end], &wg)
	}

	wg.Wait()
}
