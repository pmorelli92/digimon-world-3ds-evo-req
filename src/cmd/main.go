package main

import (
	"digimon-world-3ds-evo-req-cmd/domain"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/olekukonko/tablewriter"
)

func main() {
	// TODO: Replace with real API call
	digimonByName, _ := fetchDigimonByName()

	// Create question
	question := survey.Select{
		Message:  "Choose your Digimon:",
		Options:  digimonByName,
		Default:  "Agumon",
		PageSize: 3,
	}

	var chosen string
	err := survey.AskOne(&question, &chosen)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// TODO: Replace with real API call
	evolutions, _ := findPossibleEvolutions(chosen)

	// Create the requirements header
	header, err := domain.GetDigimonOrderedFieldsName()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(header)

	// Append rows
	for _, e := range evolutions {
		values, err := e.GetDigimonOrderedValues()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		table.Append(values)
	}

	table.Render()
}

func fetchDigimonByName() ([]string, error) {
	return []string{"Agumon", "Greymon", "GeoGreymon"}, nil
}

func findPossibleEvolutions(digimon string) ([]domain.Digimon, error) {
	return []domain.Digimon{
		{
			Name:       "Greymon",
			HP:         "1000",
			MP:         "-",
			Atk:        "80",
			Def:        "80",
			Spd:        "-",
			Int:        "-",
			Weight:     "25 or more",
			Mistake:    "3 or less",
			Happiness:  "65 or more",
			Discipline: "-",
			Battles:    "-",
			Techs:      "28 or more",
			Decode:     "12 or more",
			Quota:      "3",
		},
		{
			Name:       "GeoGreymon",
			HP:         "-",
			MP:         "-",
			Atk:        "120",
			Def:        "-",
			Spd:        "80",
			Int:        "60",
			Weight:     "28 or more",
			Mistake:    "5 or less",
			Happiness:  "-",
			Discipline: "-",
			Battles:    "15 or more",
			Techs:      "35 or more",
			Decode:     "15 or more",
			Quota:      "3",
		},
	}, nil
}
