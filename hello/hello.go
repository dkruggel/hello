package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func addToString(strIn []byte, nxtStr byte) (strOut []byte) {
	switch nxtStr {
	case 34:
		strOut = strIn
	case 44:
		strOut = strIn
	case 58:
		strOut = strIn
	case 123:
		strOut = strIn
	default:
		strOut = append(strIn, nxtStr)
	}
	return strOut
}

func findString(strIn []byte, strToFind string) (index int) {
	index = bytes.Index(strIn, []byte(strToFind))
	return index
}

var d map[string]string

func main() {
	resp, err := http.Get("https://www.wunderground.com/weather/us/mo/o'fallon/KMOOFALL7/")
	if err != nil {
		fmt.Println("There was an error with the request.", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// ioutil.WriteFile("C:\\users\\dkruggel\\desktop\\test_full.txt", body, 0644)
	var str []byte
	indexEnd := findString(body, "}}]}")
	indexBeg := findString(body, "{\"stationID\"")
	str = body[indexBeg+1 : indexEnd-1]
	var s []byte
	str = bytes.Replace(str, []byte("\""), s, -1)
	str = bytes.Replace(str, []byte("imperial:{"), s, -1)
	str = bytes.Replace(str, []byte(","), []byte{13, 10}, -1)

	var key []byte
	var val []byte
	var keyStr string
	var valStr string
	keyCheck := false
	valCheck := false
	d = make(map[string]string)

	for j := 0; j < len(str); j++ {
		if keyCheck == false {
			if string(str[j]) != ":" && string(str[j]) != "\n" {
				key = addToString(key, str[j])
			} else if string(str[j]) == "\n" {

			} else {
				keyStr = string(key)
				keyCheck = true
				key = nil
				valCheck = false
			}
		} else if valCheck == false {
			if string(str[j]) != "\r" && string(str[j]) != "\n" {
				val = addToString(val, str[j])
			} else if string(str[j]) == "\n" {

			} else {
				valStr = string(val)
				d[keyStr] = valStr
				valCheck = true
				val = nil
				keyCheck = false
				keyStr = ""
			}
		}
	}
	// ioutil.WriteFile("C:\\users\\dkruggel\\desktop\\test.txt", str, 0644)
	for index, element := range d {
		switch index {
		case "obsTimeLocal", "neighborhood", "temp", "heatIndex", "precipRate", "precipTotal":
			fmt.Println(index, ":", element)
		}
	}
}
