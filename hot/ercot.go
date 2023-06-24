package hot

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErcotResponse struct {
	LastUpdate       string `json:"lastUpdate"`
	CurrentCondition struct {
		ConditionNote    string `json:"condition_note"`
		EeaLevel         int    `json:"eea_level"`
		EnergyLevelValue int    `json:"energy_level_value"`
		State            string `json:"state"`
		Title            string `json:"title"`
		PrcValue         string `json:"prc_value"`
		Index            int    `json:"index"`
		Datetime         int    `json:"datetime"`
	} `json:"current_condition"`
	Data []struct {
		Prc      int   `json:"prc"`
		Interval int64 `json:"interval"`
	} `json:"data"`
}

type ErcotData struct {
	Note  string
	State string
	Prc   string
}

const ercotUrl = "https://www.ercot.com/api/1/services/read/dashboards/daily-prc.json"

func GetErcotData() ErcotData {

	resp, err := http.Get(ercotUrl)

	if err != nil {
		log.Fatal("Failed to retrieve Ercot data")
	}

	var ercotResponse ErcotResponse

	if err := json.NewDecoder(resp.Body).Decode(&ercotResponse); err != nil {
		log.Fatal("Failed to deserialize Ercot response")
	}

	r := ErcotData{
		Note:  ercotResponse.CurrentCondition.ConditionNote,
		State: ercotResponse.CurrentCondition.State,
		Prc:   ercotResponse.CurrentCondition.PrcValue,
	}

	return r
}
