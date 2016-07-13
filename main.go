package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	//"./godomains"
	"github.com/bitzl/godomains/lib"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: godomains <tld> <source> | %d", len(os.Args))
	}

	log.Println("Start processing.")

	//source := "https://cgit.freedesktop.org/libreoffice/dictionaries/tree/en/en_US.dic"
	tld := os.Args[1]
	if strings.Index(tld, ".") != 0 {
		log.Fatal("<tld> must start with '.' but is '" + tld + "'.")
	}
	log.Println("Looking for " + tld + " domains.")
	source := os.Args[2]
	log.Println("Streaming dictionary from " + source + "")
	target := "availableDomains.txt"

	response, err := http.Get(source)
	if err != nil {
		log.Fatal("Could not open source file " + source + ": " + err.Error())
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	destFile, err := os.Create(target)
	if err != nil {
		log.Fatal("Could not create dest file " + target + ": " + err.Error())
	}

	defer func() {
		if err := destFile.Close(); err != nil {
			log.Fatal("Could not close dest file: " + err.Error())
		}
	}()

	wordcount := process(godomains.NewHunspellWordSource(response.Body), tld, destFile)

	log.Printf("Done checking %d domain names.\n", wordcount)
}

func process(wordsource godomains.WordSource, tld string, destFile *os.File) int {
	wordcount := 0
	for wordsource.Next() {
		if wordcount%500 == 0 {
			log.Printf("Checked %d domains.", wordcount)
		}
		domain := wordsource.Word() + tld
		if godomains.IsAvailable(domain) {
			if _, err := destFile.WriteString(domain + "\n"); err != nil {
				log.Fatal(err)
			}
		}
		wordcount = wordcount + 1
	}
	if err := wordsource.Err(); err != nil {
		log.Printf("Emergency exit after %d words.\n", wordcount)
		log.Fatal(err)
	}
	return wordcount
}
