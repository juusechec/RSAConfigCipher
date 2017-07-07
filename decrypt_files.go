package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
  "./cipher_value"
)

func main() {
	// https://gobyexample.com/command-line-arguments
	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	if len(os.Args) == 1 {
		fmt.Println("Execute with rsa archive(s) name(s) as param(s):")
		execFilename := os.Args[0]
		fmt.Println(execFilename, "example_config.php.rsa")
		os.Exit(1)
	}

	inputfiles := os.Args[1:]
	for i := 0; i < len(inputfiles); i++ {
		inputfile := inputfiles[i]
		var data = ReadFile(inputfile)

		output, err := UpdateText(string(data))
		if err != nil {
			log.Fatal("Error from UpdateText: %s\n", err)
			return
		}

	  fmt.Println("File output:")
		fmt.Println(output)
	  newfilename := GetFilename(inputfile)
	  WriteFile(newfilename, output)

	  fmt.Println("A file has been created:", newfilename)
	}

	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
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
    unciphertext,_ := cipher_value.DecryptValue(content)
		replaceTagText := unciphertext
		output = strings.Replace(output, allParam, replaceTagText, -1)
	}
	return output, nil
}

func ReadFile(filename string) []byte {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fileData
}

func WriteFile(filename string, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetFilename(filename string) string {
  re := regexp.MustCompile(`(.*)\.`)
	result := re.FindString(filename)
  result = result[:len(result)-1]
  return string(result)
}
