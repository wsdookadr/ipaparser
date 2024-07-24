package main

import (
	"archive/zip"
	"errors"
	//"github.com/DHowett/go-plist"
	"howett.net/plist"
	"io"
	"log"
	"flag"
	"fmt"
)

//ParseIpa : It parses the given ipa and returns a map from the contents of Info.plist in it
func ParseIpa(name string) (map[string]interface{}, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		log.Println("Error opening ipa/zip ", err.Error())
		return nil, err
	}
	defer r.Close()

	for _, file := range r.File {
		if file.FileInfo().Name() == "Info.plist" {
			rc, err := file.Open()
			if err != nil {
				log.Println("Error opening Info.plist in zip", err.Error())
				return nil, err
			}
			buf := make([]byte, file.FileInfo().Size())
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				log.Println("Error reading Info.plist", err.Error())
				return nil, err
			}
			var info_map map[string]interface{}
			_, err = plist.Unmarshal(buf, &info_map)
			if err != nil {
				log.Println("Error reading Info.plist", err.Error())
				return nil, err
			}
			return info_map, nil
		}
	}
	return nil, errors.New("Info.plist not found")
}

func main() {
	ipa := flag.String("ipa", "", "Provide an IPA file")
	flag.Parse()
	if *ipa == "" {
		fmt.Println("Error: --ipa flag is required")
		return
	}
	parsedData, err := ParseIpa(*ipa)
	if err != nil {
		fmt.Printf("Error parsing IPA: %v\n", err)
		return
	}

	fmt.Println("Parsed IPA data:")
	for key, value := range parsedData {
		fmt.Printf("%s: %v\n", key, value)
	}
}

