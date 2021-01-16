package domain

// Digimon contains all the requirements for triggering an evolution
type Digimon struct {
	Name       string `json:"name"`
	HP         string `json:"hp"`
	MP         string `json:"mp"`
	Atk        string `json:"atk"`
	Def        string `json:"def"`
	Spd        string `json:"spd"`
	Int        string `json:"int"`
	Weight     string `json:"weight"`
	Mistake    string `json:"mistake"`
	Happiness  string `json:"happiness"`
	Discipline string `json:"discipline"`
	Battles    string `json:"battles"`
	Techs      string `json:"techs"`
	Decode     string `json:"decode"`
	Quota      string `json:"quota"`
}
