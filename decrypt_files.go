package main

import (
	"./cipherValue"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	// VerboseMode It's used to control verbose mode
	VerboseMode = false
)

func main() {
	// https://gobyexample.com/command-line-arguments
	// `os.Args` provides access to raw command-line

	if len(os.Args) == 1 {
		fmt.Println("Execute with rsa archive(s) name(s) as param(s):")
		execFilename := os.Args[0]
		fmt.Println(execFilename, "example_config.php.rsa")
		readPrompt()
		return
	}

	inputfiles := os.Args[1:]

	for i := 0; i < len(inputfiles); i++ {
		if inputfiles[i] == "--verbose" || inputfiles[i] == "-v" {
			inputfiles = removeIndex(inputfiles, i)
			VerboseMode = true
			continue
		} else if inputfiles[i] == "--help" || inputfiles[i] == "-h" {
			showHelp()
			os.Exit(0)
		} else if inputfiles[i] == "--version" {
			showVersion()
			os.Exit(0)
		}
	}

	cipherValue.VerboseMode = VerboseMode

	for key := range inputfiles {
		inputfile := inputfiles[key]
		var data = readFile(inputfile)

		output, err := updateText(string(data))
		if err != nil {
			panic("Error from UpdateText:" + err.Error())
		}

		if VerboseMode {
			fmt.Println("File output:")
			fmt.Println(output)
		}
		newfilename := getFilename(inputfile)
		writeFile(newfilename, output)
		if VerboseMode {
			fmt.Println("A file has been created:", newfilename)
		}
	}
}

//https://play.golang.org/p/gtU4HjbYoZ
func updateText(input string) (string, error) {
	re, err := regexp.Compile(`{{%rsa:(.*?)%}}`)
	if err != nil {
		return "", err
	}
	indexes := re.FindAllStringSubmatchIndex(input, -1)
	output := input
	for _, match := range indexes {
		contentStart := match[2]
		tagStart := contentStart - len("{{%rsa:")
		contentEnd := match[3]
		tagEnd := contentEnd + len("%}}")
		allParam := input[tagStart:tagEnd]
		content := input[contentStart:contentEnd]
		unciphertext, _ := cipherValue.DecryptValue(content)
		replaceTagText := unciphertext
		output = strings.Replace(output, allParam, replaceTagText, -1)
	}
	return output, nil
}

func readFile(filename string) []byte {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return fileData
}

func writeFile(filename string, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func getFilename(filename string) string {
	re := regexp.MustCompile(`(.*)\.`)
	result := re.FindString(filename)
	if len(result) == 0 {
		return filename
	}
	result = result[:len(result)-1]
	return string(result)
}

func readPrompt() {
	fmt.Print("Enter text to encrypt:\n")
	var input string
	fmt.Scanln(&input)
	ciphertext, _ := cipherValue.EncryptValue(input)
	fmt.Print("The encrypted text is:\n")
	ciphertext = "{{%rsa:" + ciphertext + "%}}\n"
	fmt.Print(ciphertext)
}

func removeIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func showHelp() {
	fmt.Println(`
Usage: rsaconfigcipher [OPTION]... [FILE]...

	-v, --verbose              explain what is being done
  -h, --help                 display this help and exit
	    --version              output version information and exit
	`)
}

func showVersion() {
	fmt.Println(`
rsaconfigcipher (juusechec Tools) 1.1.0
Copyright (C) 2017 Jorge Ulises Useche Cuellar.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by juusechec.
	`)
}
