package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

//go:embed main.go.tmpl
var MAIN_TMPL string

//go:embed day.go.tmpl
var DAY_TMPL string

type mainConfig struct {
	Year int
}

type dayConfig struct {
	Day int
}

func main() {
	mainFile, err := os.Create("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer mainFile.Close()

	mainTmpl, err := template.New("main.go").Parse(MAIN_TMPL)
	if err != nil {
		log.Fatal(err)
	}

	err = mainTmpl.Execute(mainFile, mainConfig{Year: time.Now().Year()})
	if err != nil {
		log.Fatal(err)
	}

	dayTmpl, err := template.New("day.go").Parse(DAY_TMPL)
	if err != nil {
		log.Fatal(err)
	}

	for day := 1; day <= 25; day++ {
		dayFile, err := os.Create(fmt.Sprintf("day%02d.go", day))
		if err != nil {
			log.Fatal(err)
		}

		defer dayFile.Close()
		err = dayTmpl.Execute(dayFile, dayConfig{Day: day})
		if err != nil {
			log.Fatal(err)
		}
	}
}
