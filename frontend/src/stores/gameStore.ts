// frontend/src/stores/gameStore.ts
import { defineStore } from 'pinia'
import { CreateNewGame, LoadGameData } from '../../wailsjs/go/main/App'
import yaml from 'js-yaml'

export const useGameStore = defineStore('game', {
  state: () => ({
    isLoaded: false,
    activeSaveId: '',
    saveState: null as any, // Holds Captain, Ship, Credits, etc.
    planets: null as any,
    planetTraits: null as any,
    commodities: null as any,
    modules: null as any,
    events: null as any,
    crew: null as any,
    gameSettings: null as any
  }),
  actions: {
    async startNewGame(captain: string, ship: string) {
      try {
        // 1. Tell Go to create the folder and copy the files
        const saveId = await CreateNewGame(captain, ship)
        if (!saveId) throw new Error("Backend failed to create save folder.")
        
        this.activeSaveId = saveId
        
        // 2. Load the data directly from the newly created folder
        await this.loadGameFromId(saveId)
        return true
      } catch (error) {
        console.error("New game sequence failed:", error)
        return false
      }
    },
    async loadGameFromId(saveId: string) {
      try {
        const rawData = await LoadGameData(saveId)

        this.saveState = yaml.load(rawData['save_state.yaml'])
        this.planets = yaml.load(rawData['planets.yaml'])
        this.planetTraits = yaml.load(rawData['planet_traits.yaml'])
        this.commodities = yaml.load(rawData['commodities.yaml'])
        this.modules = yaml.load(rawData['modules.yaml'])
        this.events = yaml.load(rawData['events.yaml'])
        this.crew = yaml.load(rawData['crew_roles.yaml'])
        this.gameSettings = yaml.load(rawData['game_settings.yaml'])

        this.isLoaded = true
        console.log(`Successfully hydrated Save ID: ${saveId}`)
      } catch (error) {
        console.error("Failed to load save data:", error)
      }
    }
  }
})