package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pkg/profile"
)

//	var rootCmd = &cobra.Command{
//		Use:   "mycli",
//		Short: "A simple CLI created with Cobra",
//		Long:  `A simple example CLI created with Cobra.`,
//		Run: func(cmd *cobra.Command, args []string) {
//			fmt.Println("Hello from mycli!")
//		},
//	}
type Ammo struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	ShortName       string  `json:"shortName"`
	Weight          float64 `json:"weight"`
	Caliber         string  `json:"caliber"`
	StackMaxSize    int     `json:"stackMaxSize"`
	Tracer          bool    `json:"tracer"`
	TracerColor     string  `json:"tracerColor"`
	AmmoType        string  `json:"ammoType"`
	ProjectileCount int     `json:"projectileCount"`
	Ballistics      struct {
		Damage              int     `json:"damage"`
		ArmorDamage         int     `json:"armorDamage"`
		FragmentationChance float64 `json:"fragmentationChance"`
		RicochetChance      float64 `json:"ricochetChance"`
		PenetrationChance   float64 `json:"penetrationChance"`
		PenetrationPower    int     `json:"penetrationPower"`
		Accuracy            int     `json:"accuracy"`
		Recoil              int     `json:"recoil"`
		InitialSpeed        int     `json:"initialSpeed"`
	} `json:"ballistics"`
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.MemProfile).Stop()

	// Enable memory profiling
	profile.MemProfileRate(1)

	// Measure performance
	startTime := time.Now()

	// Read and process JSON data
	err := processJSON("Data/ammunition.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calculate and print elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("Elapsed Time: %s\n", elapsedTime)
}

func processJSON(filename string) error {
	// Read JSON data from file
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Create a map to hold the parsed Ammo objects
	var ammoMap map[string]Ammo

	// Unmarshal JSON data into the map
	err = json.Unmarshal(jsonData, &ammoMap)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	//Access and print information about each Ammo object
	for id, ammo := range ammoMap {
		fmt.Printf("ID: %s\n", id)
		fmt.Printf("Name: %s\n", ammo.Name)
		fmt.Printf("Weight: %f\n", ammo.Weight)
		fmt.Printf("Caliber: %s\n", ammo.Caliber)
		fmt.Printf("Damage: %d\n", ammo.Ballistics.Damage)
		fmt.Println("------------------------")
	}

	return nil
}
