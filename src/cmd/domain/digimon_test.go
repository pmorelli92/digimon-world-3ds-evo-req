package domain

import (
	"reflect"
	"testing"
)

func TestGetDigimonOrderedFieldsName(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		{
			name: "Success",
			want: []string{
				"Name",
				"HP",
				"MP",
				"Atk",
				"Def",
				"Spd",
				"Int",
				"Weight",
				"Mistake",
				"Happiness",
				"Discipline",
				"Battles",
				"Techs",
				"Decode",
				"Quota",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDigimonOrderedFieldsName()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDigimonOrderedFieldsName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDigimonOrderedFieldsName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDigimon_GetDigimonOrderedValues(t *testing.T) {
	tests := []struct {
		name    string
		digimon Digimon
		want    []string
		wantErr bool
	}{
		{
			name: "Success",
			digimon: Digimon{
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
			want: []string{
				"Greymon",
				"1000",
				"-",
				"80",
				"80",
				"-",
				"-",
				"25 or more",
				"3 or less",
				"65 or more",
				"-",
				"-",
				"28 or more",
				"12 or more",
				"3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.digimon.GetDigimonOrderedValues()
			if (err != nil) != tt.wantErr {
				t.Errorf("Digimon.GetDigimonOrderedValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Digimon.GetDigimonOrderedValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
