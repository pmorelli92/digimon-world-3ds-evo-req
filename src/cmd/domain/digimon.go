package domain

import (
	"fmt"
	"reflect"
	"strconv"
)

// Digimon contains all the requirements for triggering an evolution
type Digimon struct {
	Name       string `order:"0" json:"name"`
	HP         string `order:"1" json:"hp"`
	MP         string `order:"2" json:"mp"`
	Atk        string `order:"3" json:"atk"`
	Def        string `order:"4" json:"def"`
	Spd        string `order:"5" json:"spd"`
	Int        string `order:"6" json:"int"`
	Weight     string `order:"7" json:"weight"`
	Mistake    string `order:"8" json:"mistake"`
	Happiness  string `order:"9" json:"happiness"`
	Discipline string `order:"10" json:"discipline"`
	Battles    string `order:"11" json:"battles"`
	Techs      string `order:"12" json:"techs"`
	Decode     string `order:"13" json:"decode"`
	Quota      string `order:"14" json:"quota"`
}

// GetDigimonOrderedFieldsName returns all the fields name
// of the Digimon struct ordered using the order tag
func GetDigimonOrderedFieldsName() ([]string, error) {
	t := reflect.TypeOf(Digimon{})
	orderedFields := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		order := field.Tag.Get("order")
		orderInt, err := strconv.Atoi(order)
		if err != nil {
			return nil, fmt.Errorf("error parsing order tag on struct: %w", err)
		}

		orderedFields[orderInt] = field.Name
	}

	return orderedFields, nil
}

// GetDigimonOrderedValues returns all the values
// of a Digimon struct instance ordered using the order tag
func (d Digimon) GetDigimonOrderedValues() ([]string, error) {
	orderedFields, err := GetDigimonOrderedFieldsName()
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(d)
	orderedValues := make([]string, len(orderedFields))

	for i, f := range orderedFields {
		value := v.FieldByName(f).Interface()
		orderedValues[i] = value.(string)
	}

	return orderedValues, nil
}
