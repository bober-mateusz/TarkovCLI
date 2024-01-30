// ammo.go
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bober-mateusz/TarkovCLI/models"
	"github.com/spf13/cobra"
)

var caliberFlag string
var sortOption string
var ammoFlag string
var ammoList []models.Ammo // Declare ammoList as a package-level variable

var ammoCmd = &cobra.Command{
	Use:   "ammo",
	Short: "Display ammo types",
	Run: func(cmd *cobra.Command, args []string) {
		err := showAmmo(caliberFlag, sortOption, ammoFlag)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func showAmmo(caliber string, sortOption string, partialAmmo string) error {
	// Get the absolute path to the Data folder
	dataFolder := filepath.Join(".", "Data")

	// Read and process JSON data
	filename := filepath.Join(dataFolder, "ammunition.json")
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal JSON data into a slice of models.Ammo
	err = json.Unmarshal(jsonData, &ammoList)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Autofill the ammoFlag if it is a partial identifier
	if partialAmmo != "" {
		for _, ammo := range ammoList {
			if strings.HasPrefix(ammo.Caliber, partialAmmo) {
				ammoFlag = ammo.Caliber
				break
			}
		}
	}

	// Sort ammoList based on the specified option
	switch sortOption {
	case "damage":
		sort.Slice(ammoList, func(i, j int) bool {
			return ammoList[i].Ballistics.Damage > ammoList[j].Ballistics.Damage
		})
	case "penpower":
		sort.Slice(ammoList, func(i, j int) bool {
			return ammoList[i].Ballistics.PenetrationPower > ammoList[j].Ballistics.PenetrationPower
		})
	}

	// Display information about each Ammo object
	for _, ammo := range ammoList {
		// Check if the specified caliber is a substring of the actual caliber
		if caliber == "" || strings.Contains(ammo.Caliber, caliber) {
			fmt.Printf("Name: %s\n", ammo.Name)
			fmt.Printf("Caliber: %s\n", ammo.Caliber)
			fmt.Printf("Damage: %d\n", ammo.Ballistics.Damage)
			fmt.Printf("Penetration Power: %d\n", ammo.Ballistics.PenetrationPower)
			fmt.Println("------------------------")
		}
	}

	return nil
}

func init() {
	// Add the command to the root command
	rootCmd.AddCommand(ammoCmd)

	// Add the --caliber flag to the ammo command
	ammoCmd.Flags().StringVarP(&caliberFlag, "caliber", "c", "", "Filter ammo by caliber")

	// Add the --sort flag to the ammo command
	ammoCmd.Flags().StringVarP(&sortOption, "sort", "s", "", "Sort ammo by 'damage' or 'penpower'")

	// Add the --ammo flag to the ammo command
	ammoCmd.Flags().StringVarP(&ammoFlag, "ammo", "a", "", "Specify ammo caliber")

	// Mark --ammo flag for autocomplete
	_ = ammoCmd.RegisterFlagCompletionFunc("ammo", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Autocomplete based on the partial input
		var completions []string
		for _, ammo := range ammoList {
			if strings.HasPrefix(ammo.Caliber, toComplete) {
				completions = append(completions, ammo.Caliber)
			}
		}
		return completions, cobra.ShellCompDirectiveDefault
	})
}
