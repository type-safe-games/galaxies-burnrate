<script setup lang="ts">
import { QuitGame } from '../../wailsjs/go/main/App'
import pkg from '../../package.json' 

const emit = defineEmits(['navigate'])
const appVersion = pkg.version 

// Updated to point to the new setup screen instead of directly to 'game'
const newGame = () => emit('navigate', 'new_game_setup')
const loadGame = () => emit('navigate', 'load_game')
const openOptions = () => emit('navigate', 'options')
const quitApp = () => QuitGame()
</script>

<template>
  <div class="menu-container">
    <div class="header">
      <h1>GALAXIES</h1>
      <h2>BURN RATE</h2>
    </div>
    
    <nav class="button-group">
      <button class="terminal-btn" @click="newGame">New Vessel</button>
      <button class="terminal-btn" @click="loadGame">Load Registry</button> 
      <button class="terminal-btn" @click="openOptions">Options</button>
      <button class="terminal-btn" @click="quitApp">Quit Game</button>
    </nav>

    <div class="footer">
      <span>v{{ appVersion }} // TYPE SAFE GAMES</span>
    </div>
  </div>
</template>

<style scoped>
.menu-container {
  display: flex;
  flex-direction: column;
  justify-content: center; 
  align-items: center;
  height: 100%;
  padding: 60px 0;
  box-sizing: border-box;
  position: relative; 
}

.header { 
  text-align: center; 
  margin-bottom: 80px; 
}

.header h1 {
  font-size: 110px;
  margin: 0;
  letter-spacing: 12px;
  text-shadow: var(--text-glow);
}

.header h2 {
  font-size: 54px;
  margin: -10px 0 0 0;
  letter-spacing: 28px;
  opacity: 0.8;
}

.button-group {
  display: flex;
  flex-direction: column;
  gap: 25px;
  width: 450px;
}

.footer {
  position: absolute; /* Pins it to the container edges */
  bottom: 20px;
  left: 30px;
  font-size: 18px; /* Slightly smaller */
  opacity: 0.3; /* Pushed further into the background */
  letter-spacing: 2px;
  pointer-events: none; /* Prevents it from interfering with any mouse clicks */
}
</style>