package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadStats(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("file: %s\n", stats.Name())
	fmt.Printf("Time modified: %s\n", stats.ModTime().Format("13:33:22"))
}

func ReadWholeFile(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(contents))
}

func ReadByLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByWord(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByBytes(filename string, bufsize uint8) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	buf := make([]byte, bufsize)

	for {
		totalRead, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}

		fmt.Println(string(buf[:totalRead]))
	}
}

func ReadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "=")
		fmt.Printf("%s: %v\n", raw[0], raw[1])
	}
}

func main() {
	const (
		filename string = "text.txt"
		configfile string = "configfile.cfg"
	)

	ReadStats(filename)
	ReadWholeFile(filename)
	ReadByLine(filename)
	ReadByWord(filename)
	ReadByBytes(filename, 8)
	ReadConfig(configfile)
}
