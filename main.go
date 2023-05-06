package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	simul "github.com/Guilherme-De-Marchi/nbody-go/simulation"
	"github.com/Guilherme-De-Marchi/nbody-go/ui"
)

var (
	simulConf SimulConfig
)

type SimulConfig struct {
	GenerationType string         `json:"generation_type,omitempty"`
	Universe       simul.Universe `json:"universe,omitempty"`
	RandOpt        simul.RandOpt  `json:"random_options,omitempty"`
	EditOpt        simul.EditOpt  `json:"edit_options,omitempty"`
}

func main() {
	confJ, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("[INTERNAL ERROR]: ", err)
	}
	json.Unmarshal(confJ, &simulConf)

	var universe *simul.Universe
	if simulConf.GenerationType == "randomized" {
		rand.Seed(time.Now().UnixNano())
		universe = simul.NewRandomUniverse(
			simulConf.Universe.Size,
			simul.G,
			simulConf.RandOpt.MassR,
			simulConf.RandOpt.RadR,
			simulConf.RandOpt.ObjectQtt,
		)
	} else if simulConf.GenerationType == "prefab" {
		log.Fatal("Generation type 'prefab' not implemented")
	} else {
		log.Fatal("Invalid value for field 'generation_type'")
	}

	s := simul.NewSimulation(universe, simulConf.RandOpt, simulConf.EditOpt)
	(*ui.Game)(s).Init()
}
