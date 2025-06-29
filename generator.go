package main

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"
)

type GenConfig struct {
	CacheType   string
	StorageType string
	OutputPath  string
}

type TemplConfig struct {
	TemplatePath string
	OutputPath   string
}

var templates = []TemplConfig{
	{"templates/config.yaml.tmpl", "config"},
	{"templates/script.lua", "scripts"},
	{"templates/docker-compose.yaml.tmpl", ""},
	{"templates/Taskfile.yaml", ""},
}

func generateFiles(config *GenConfig) {
	for _, t := range templates {
		_, name := path.Split(t.TemplatePath)
		if err := generateFile(t.TemplatePath, name, t.OutputPath, config); err != nil {
			fmt.Printf("%s", err.Error())
		}
	}
}

func generateFile(tmplPath, tmplName, outputPath string, config *GenConfig) error {
	//парсим файл
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}

	outPath := path.Join(outputPath, wipeTmplExt(tmplName))

	//если задали выходную папку
	if config.OutputPath != "" {
		os.MkdirAll(path.Join(config.OutputPath, outputPath), os.ModeDir)
		outPath = path.Join(config.OutputPath, outPath)
	} else {
		os.MkdirAll(outputPath, os.ModeDir)
	}

	//создаем и записываем в файл
	outputFile, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	defer outputFile.Close()

	return tmpl.Execute(outputFile, config)
}

func wipeTmplExt(path string) string {
	stringShards := strings.Split(path, ".")
	countShard := len(stringShards)

	if stringShards[countShard-1] == "tmpl" {
		stringShards[countShard-1] = ""
	}

	return strings.Join(stringShards, ".")
}
