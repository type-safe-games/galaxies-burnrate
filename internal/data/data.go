package data

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

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
	settingsPath := filepath.Join(appConfigDir, "clientsettings.yaml")

	// 2. Check if the file already exists on the player's hard drive
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		fmt.Println("First launch detected. Extracting default clientsettings.yaml...")

		// Create the directory
		os.MkdirAll(appConfigDir, 0755)

		// Read the default settings using the 'fs' passed from main.go
		defaultSettings, err := fs.ReadFile("game_data/clientsettings.yaml")
		if err == nil {
			// Write the embedded file to the user's hard drive
			os.WriteFile(settingsPath, defaultSettings, 0644)
		} else {
			fmt.Println("Critical Error: Default clientsettings.yaml missing from embed.")
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
		"crew.yaml",
		"gamesettings.yaml",
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
