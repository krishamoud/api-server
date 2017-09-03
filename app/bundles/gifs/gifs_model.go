package gifs

import (
	_ "bytes"
	"fmt"
	_ "io"
	"net/http"
	_ "strings"
)

func NewGif(r *http.Request) bool {
	// fmt.Println("here")
	gif := loadFile(r)
	fmt.Println(gif)
	return true
}

func loadFile(r *http.Request) error {
	file, _, err := r.FormFile("gif")
	if err != nil {
		return err
	}
	defer file.Close()
	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	if _, err = file.Read(buff); err != nil {
		return err
	}

	fmt.Println(http.DetectContentType(buff))
	return nil
}
