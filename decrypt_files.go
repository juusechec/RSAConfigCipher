package main

import (
	"./cipher_value"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	VerboseMode = false
)

func main() {
	// https://gobyexample.com/command-line-arguments
	// `os.Args` provides access to raw command-line

	if len(os.Args) == 1 {
		fmt.Println("Execute with rsa archive(s) name(s) as param(s):")
		execFilename := os.Args[0]
		fmt.Println(execFilename, "example_config.php.rsa")
		ReadPrompt()
		return
	}

	inputfiles := os.Args[1:]

	for i := 0; i < len(inputfiles); i++ {
		if inputfiles[i] == "--verbose" || inputfiles[i] == "-v" {
			inputfiles = RemoveIndex(inputfiles, i)
			VerboseMode = true
			continue
		} else if inputfiles[i] == "--help" || inputfiles[i] == "-h" {
			inputfiles = RemoveIndex(inputfiles, i)
			showHelp()
			os.Exit(0)
		}
	}

	cipher_value.VerboseMode = VerboseMode

	for key := range inputfiles {
		inputfile := inputfiles[key]
		var data = ReadFile(inputfile)

		output, err := UpdateText(string(data))
		if err != nil {
			panic("Error from UpdateText:" + err.Error())
			return
		}

		if VerboseMode {
			fmt.Println("File output:")
			fmt.Println(output)
		}
		newfilename := GetFilename(inputfile)
		WriteFile(newfilename, output)
		if VerboseMode {
			fmt.Println("A file has been created:", newfilename)
		}
	}
}

//https://play.golang.org/p/gtU4HjbYoZ
func UpdateText(input string) (string, error) {
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
		unciphertext, _ := cipher_value.DecryptValue(content)
		replaceTagText := unciphertext
		output = strings.Replace(output, allParam, replaceTagText, -1)
	}
	return output, nil
}

func ReadFile(filename string) []byte {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return fileData
}

func WriteFile(filename string, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func GetFilename(filename string) string {
	re := regexp.MustCompile(`(.*)\.`)
	result := re.FindString(filename)
	if len(result) == 0 {
		return filename
	}
	result = result[:len(result)-1]
	return string(result)
}

func ReadPrompt() {
	fmt.Print("Enter text to encrypt:\n")
	var input string
	fmt.Scanln(&input)
	ciphertext, _ := cipher_value.EncryptValue(input)
	fmt.Print("The encrypted text is:\n")
	ciphertext = "{{%rsa:" + ciphertext + "%}}\n"
	fmt.Print(ciphertext)
}

func RemoveIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func showHelp() {
	fmt.Println(`
Usage: rsaconfigcipher [OPTION]... [FILE]...

  -h, --help                 show help
  -v, --verbose              verbose mode
	`)
}
