package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"
)

//go:embed templates/*
var templateFS embed.FS

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
	{"config.yaml.tmpl", "config"},
	{"script.lua", "scripts"},
	{"docker-compose.yaml.tmpl", ""},
	{"Taskfile.yaml", ""},
}

func generateFiles(config *GenConfig) {
	for _, t := range templates {
		if err := generateFile(t.TemplatePath, t.OutputPath, config); err != nil {
			fmt.Printf("Error generating %s: %v\n", t.TemplatePath, err)
		}
	}
}

func generateFile(tmplName, outputPath string, config *GenConfig) error {
	// Читаем шаблон из встроенной файловой системы
	tmplContent, err := templateFS.ReadFile("templates/" + tmplName)
	if err != nil {
		return fmt.Errorf("error reading template: %w", err)
	}

	// Создаем шаблон из строки
	tmpl, err := template.New(tmplName).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Формируем путь для выходного файла
	outputFileName := wipeTmplExt(tmplName)
	fullOutputPath := path.Join(outputPath, outputFileName)
	if config.OutputPath != "" {
		fullOutputPath = path.Join(config.OutputPath, fullOutputPath)
	}

	// Создаем директории, если нужно
	dir := path.Dir(fullOutputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directories: %w", err)
	}

	// Создаем и записываем файл
	outputFile, err := os.Create(fullOutputPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer outputFile.Close()

	return tmpl.Execute(outputFile, config)
}

func wipeTmplExt(path string) string {
	if strings.HasSuffix(path, ".tmpl") {
		return strings.TrimSuffix(path, ".tmpl")
	}
	return path
}
