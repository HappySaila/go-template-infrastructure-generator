package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Data struct {
	Camel     string
	CapsCamel string
	Dashed    string
	Letter    string
	Namespace string
	Sentence  string
	Snake     string
}

func main() {
	// Define data to pass to template (you can replace this with your own data)
	// data := struct {
	// 	Camel     string
	// 	CapsCamel string
	// 	Dashed    string
	// 	Letter    string
	// 	Namespace string
	// 	Sentence  string
	// 	SnakeCase string
	// }{
	// 	"goTemplateInfrastructure",
	// 	"GoTemplateInfrastructure",
	// 	"go-template-infrastructure",
	// 	"g",
	// 	"GO_TEMPLATE_INFRASTRUCTURE",
	// 	"Go Template Infrastructure",
	// 	"go_template_infrastructure",
	// }
	if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go [serviceName]")
        os.Exit(1)
    }
    serviceName := os.Args[1]
	data := generateStruct(serviceName)
	
	// Remove the output directory ./output
	err := os.RemoveAll("./output")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Run template recursively over files in config/template directory
	inputDir := "./config/template"
	outputDir := "./output"
	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			// Get the relative path from the input directory
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}

			outputSubdir := filepath.Join(outputDir, data.Dashed)

			// Create output subdirectory if it doesn't exist
			if _, err := os.Stat(outputSubdir); os.IsNotExist(err) {
				err := os.MkdirAll(outputSubdir, 0755)
				if err != nil {
					return err
				}
			}

			output, err := runFileTemplate(path, data)
			if err != nil {
				return err
			}

			// Write output to corresponding file in output directory
			outputPath := filepath.Join(outputDir, relPath)
			outputPath_ := strings.Replace(outputPath, "go-template-infrastructure", data.Dashed, -1)
			outputPath_ = strings.Replace(outputPath_, ".tpl", "", -1)
			if err := os.MkdirAll(filepath.Dir(outputPath_), 0755); err != nil {
				return err
			}
			err = os.WriteFile(outputPath_, []byte(output), 0644)
			if err != nil {
				return err
			}
			fmt.Printf("Template output written to: %s\n", outputPath_)
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func runFileTemplate(filePath string, data interface{}) (string, error) {
	// Create template from filename
	fileName := cleanFileName(filePath) + ".tpl"
	tpl := template.New(fileName).Delims("[[", "]]")
	tpl, err := tpl.ParseFiles(filePath)
	if err != nil {
		return "", err
	}

	// Execute template to get returned string
	tplBytes := new(bytes.Buffer)
	err = tpl.Execute(tplBytes, data)
	if err != nil {
		return "", err
	}
	out := tplBytes.String()
	return out, nil
}

func cleanFileName(filePath string) string {
	return filepath.Base(filePath[:len(filePath)-len(filepath.Ext(filePath))])
}

func generateStruct(sentence string) Data {
	words := strings.Fields(sentence)
	if len(words) == 0 {
		return Data{}
	}
	lowered := make([]string, len(words))
	for i, w := range words {
		lowered[i] = strings.ToLower(w)
	}
	camel := lowered[0]
	for _, word := range lowered[1:] {
		camel += strings.ToUpper(word[:1]) + word[1:]
	}
	capsCamel := ""
	for _, word := range lowered {
		capsCamel += strings.ToUpper(word[:1]) + word[1:]
	}
	dashed := strings.Join(lowered, "-")
	letter := ""
	for _, word := range lowered {
		if len(word) > 0 {
			letter += string(word[0])
		}
	}
	namespace := strings.ToUpper(strings.Join(lowered, "_"))
	sentenceOut := strings.Join(words, " ")
	snake := strings.Join(lowered, "_")
	return Data{
		Camel:     camel,
		CapsCamel: capsCamel,
		Dashed:    dashed,
		Letter:    letter,
		Namespace: namespace,
		Sentence:  sentenceOut,
		Snake:     snake,
	}
}
