package domain

import (
	"fmt"
	"reflect"
	"strconv"
)

// Digimon contains all the requirements for triggering an evolution
type Digimon struct {
	Name       string `order:"0"`
	HP         string `order:"1"`
	MP         string `order:"2"`
	Atk        string `order:"3"`
	Def        string `order:"4"`
	Spd        string `order:"5"`
	Int        string `order:"6"`
	Weight     string `order:"7"`
	Mistake    string `order:"8"`
	Happiness  string `order:"9"`
	Discipline string `order:"10"`
	Battles    string `order:"11"`
	Techs      string `order:"12"`
	Decode     string `order:"13"`
	Quota      string `order:"14"`
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
