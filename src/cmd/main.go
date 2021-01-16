package main

import (
	"digimon-world-3ds-evo-req-cmd/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	type response struct {
		Digimons []string `json:"digimons"`
	}

	resp, err := http.Get("https://digimon-api-pmorelli92.cloud.okteto.net/api/digimon")
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respStruct response
	err = json.Unmarshal(bytes, &respStruct)

	return respStruct.Digimons, nil
}

func findPossibleEvolutions(digimon string) ([]domain.Digimon, error) {
	type response struct {
		Digimons []domain.Digimon `json:"digimons"`
	}

	url := fmt.Sprintf("https://digimon-api-pmorelli92.cloud.okteto.net/api/digimon/%s", digimon)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respStruct response
	err = json.Unmarshal(bytes, &respStruct)

	return respStruct.Digimons, nil
}
