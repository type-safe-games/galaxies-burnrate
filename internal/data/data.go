package data

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// ClientSettings strictly matches your clientsettings.yaml structure
type ClientSettings struct {
	Display struct {
		Width      int  `yaml:"width"`
		Height     int  `yaml:"height"`
		Fullscreen bool `yaml:"fullscreen"`
	} `yaml:"display"`
	Debug struct {
		SkipSplash bool `yaml:"skip_splash"`
		DevMode    bool `yaml:"dev_mode"`
	} `yaml:"debug"`
}

// GameSettings represents the static configuration in game_settings.yaml
type GameSettings struct {
	PlayerDefaults struct {
		StartingCredits  int    `yaml:"starting_credits"`
		StartingLocation string `yaml:"starting_location"`
		StartingFuel     int    `yaml:"starting_fuel"`
	} `yaml:"player_defaults"`
}

// SaveState represents the dynamic player data in save_state.yaml
type SaveState struct {
	SaveID          string `json:"save_id"`
	CaptainName     string `yaml:"captain_name" json:"captain_name"`
	ShipName        string `yaml:"ship_name" json:"ship_name"`
	Credits         int    `yaml:"credits" json:"credits"`
	CurrentDay      int    `yaml:"current_day" json:"current_day"`
	CurrentLocation string `yaml:"current_location" json:"current_location"`
	CurrentFuel     int    `yaml:"current_fuel" json:"current_fuel"`
}

// SaveMeta holds the summary data to display on the Load Screen.
type SaveMeta struct {
	SaveID          string `json:"save_id"`
	CaptainName     string `json:"captain_name" yaml:"captain_name"`
	ShipName        string `json:"ship_name" yaml:"ship_name"`
	Credits         int    `json:"credits" yaml:"credits"`
	CurrentDay      int    `json:"current_day" yaml:"current_day"`
	CurrentLocation string `json:"current_location" yaml:"current_location"`
	CurrentFuel     int    `json:"current_fuel" yaml:"current_fuel"`
}

// LoadClientSettings handles the Extraction Pattern for user preferences.
// Notice the Capital 'L' - this exports the function so main.go can use it!
// We also accept 'fs embed.FS' so main.go can hand us the embedded files.
func LoadClientSettings(fs embed.FS) ClientSettings {
	var settings ClientSettings

	// 1. Get the OS-specific config directory (~/.config or AppData)
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Warning: Could not find OS config dir, defaulting to local folder")
		configDir = "."
	}

	// Create a dedicated folder for the game: ~/.config/GalaxiesBurnRate/
	appConfigDir := filepath.Join(configDir, "GalaxiesBurnRate")
	settingsPath := filepath.Join(appConfigDir, "client_settings.yaml")

	// 2. Check if the file already exists on the player's hard drive
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		fmt.Println("First launch detected. Extracting default client_settings.yaml...")

		// Create the directory
		os.MkdirAll(appConfigDir, 0755)

		// Read the default settings using the 'fs' passed from main.go
		defaultSettings, err := fs.ReadFile("game_data/client_settings.yaml")
		if err == nil {
			// Write the embedded file to the user's hard drive
			os.WriteFile(settingsPath, defaultSettings, 0644)
		} else {
			fmt.Println("Critical Error: Default client_settings.yaml missing from embed.")
		}
	}

	// 3. Read the file from the hard drive
	fileData, err := os.ReadFile(settingsPath)
	if err != nil {
		fmt.Println("Error reading settings file:", err)
		return settings
	}

	// 4. Parse the YAML into the Go struct
	err = yaml.Unmarshal(fileData, &settings)
	if err != nil {
		fmt.Println("Error parsing settings YAML:", err)
	}

	return settings
}

// LoadEmbeddedGameData scoops up all the read-only YAMLs to send to Vue.
// Capital 'L' to export, and accepts 'fs embed.FS'.
func LoadEmbeddedGameData(fs embed.FS) map[string]string {
	gameData := make(map[string]string)

	filesToLoad := []string{
		"planets.yaml",
		"planet_traits.yaml",
		"commodities.yaml",
		"modules.yaml",
		"events.yaml",
		"crew_roles.yaml",
		"game_settings.yaml",
	}

	for _, filename := range filesToLoad {
		// Read directly using the 'fs' passed from main.go
		data, err := fs.ReadFile("game_data/" + filename)
		if err == nil {
			gameData[filename] = string(data)
		} else {
			fmt.Println("Failed to load embedded file:", filename)
		}
	}

	return gameData
}

// CreateNewSave creates a unique folder, extracts the embedded YAMLs to it, and initializes save_state.yaml
func CreateNewSave(fs embed.FS, captainName string, shipName string) (string, error) {
	configDir, _ := os.UserConfigDir()
	savesDir := filepath.Join(configDir, "GalaxiesBurnRate", "saves")
	os.MkdirAll(savesDir, 0755)

	saveID := fmt.Sprintf("save_%d", time.Now().Unix())
	newSavePath := filepath.Join(savesDir, saveID)
	os.MkdirAll(newSavePath, 0755)

	filesToCopy := []string{
		"planets.yaml",
		"planet_traits.yaml",
		"commodities.yaml",
		"modules.yaml",
		"events.yaml",
		"crew_roles.yaml",
		"game_settings.yaml",
	}

	for _, file := range filesToCopy {
		data, err := fs.ReadFile("game_data/" + file)
		if err == nil {
			os.WriteFile(filepath.Join(newSavePath, file), data, 0644)
		}
	}

	// --- NEW: Read the defaults directly from the embedded game_settings.yaml ---
	var settings GameSettings
	settingsData, err := fs.ReadFile("game_data/game_settings.yaml")
	if err == nil {
		yaml.Unmarshal(settingsData, &settings)
	} else {
		// Absolute fallback just in case the file is missing/corrupted
		settings.PlayerDefaults.StartingCredits = 5000
		settings.PlayerDefaults.StartingLocation = "plt_prime"
		settings.PlayerDefaults.StartingFuel = 100
	}

	// Generate the initial save_state.yaml using the parsed YAML defaults
	initialState := SaveState{
		CaptainName:     captainName,
		ShipName:        shipName,
		Credits:         settings.PlayerDefaults.StartingCredits,
		CurrentDay:      1,
		CurrentLocation: settings.PlayerDefaults.StartingLocation,
		CurrentFuel:     settings.PlayerDefaults.StartingFuel,
	}

	stateData, _ := yaml.Marshal(&initialState)
	os.WriteFile(filepath.Join(newSavePath, "save_state.yaml"), stateData, 0644)

	return saveID, nil
}

// LoadSaveGame reads the YAMLs directly from the player's specific save folder
func LoadSaveGame(saveID string) map[string]string {
	gameData := make(map[string]string)
	configDir, _ := os.UserConfigDir()
	savePath := filepath.Join(configDir, "GalaxiesBurnRate", "saves", saveID)

	filesToLoad := []string{
		"planets.yaml",
		"planet_traits.yaml",
		"commodities.yaml",
		"modules.yaml",
		"events.yaml",
		"crew_roles.yaml",
		"game_settings.yaml",
		"save_state.yaml", // Don't forget to load the dynamic state!
	}

	for _, filename := range filesToLoad {
		data, err := os.ReadFile(filepath.Join(savePath, filename))
		if err == nil {
			gameData[filename] = string(data)
		} else {
			fmt.Println("Failed to load save file:", filename)
		}
	}

	return gameData
}

func GetAvailableSaves() []SaveMeta {
	var saves []SaveMeta

	configDir, err := os.UserConfigDir()
	if err != nil {
		return saves
	}

	savesDir := filepath.Join(configDir, "GalaxiesBurnRate", "saves")
	entries, err := os.ReadDir(savesDir)
	if err != nil {
		return saves // Returns empty array if no saves exist yet
	}

	for _, entry := range entries {
		if entry.IsDir() {
			saveID := entry.Name()
			statePath := filepath.Join(savesDir, saveID, "save_state.yaml")

			// If a save_state.yaml exists in this folder, read it
			fileData, err := os.ReadFile(statePath)
			if err == nil {
				var meta SaveMeta
				err = yaml.Unmarshal(fileData, &meta)
				if err == nil {
					meta.SaveID = saveID
					saves = append(saves, meta)
				}
			}
		}
	}

	return saves
}
