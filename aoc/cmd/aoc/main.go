package main

import (
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"time"
)

//go:embed go.mod.tmpl
var GO_MOD_TMPL string

//go:embed main.go.tmpl
var MAIN_TMPL string

//go:embed day.go.tmpl
var DAY_TMPL string

type goModConfig struct {
	Module string
}

type mainConfig struct {
	Year int
}

type dayConfig struct {
	Day int
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("It must be used like this: %s module-name", os.Args[0])
	}

	runTemplate("go.mod", GO_MOD_TMPL, goModConfig{os.Args[1]})
	runTemplate("main.go", MAIN_TMPL, mainConfig{Year: time.Now().Year()})

	for day := 1; day <= 25; day++ {
		runTemplate(fmt.Sprintf("day%02d.go", day), DAY_TMPL, dayConfig{Day: day})
	}
}

func runTemplate(filename string, tmplContent string, data any) {
	_, err := os.Stat(filename)
	if !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("File %s already exist, stopping here to avoid erasing anything.", filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tmpl, _ := template.New(filename).Parse(tmplContent)
	_ = tmpl.Execute(file, data)

	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
