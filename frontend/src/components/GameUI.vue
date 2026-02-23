<script setup lang="ts">
import { useGameStore } from '../stores/gameStore'

const emit = defineEmits(['navigate'])
const game = useGameStore()

const abortMission = () => {
  emit('navigate', 'main_menu')
}
</script>

<template>
  <div class="game-container">
    <div class="hud-header">
      <div>
        <h2>CPT. {{ game.saveState?.captain_name }}</h2>
        <h3>VESSEL: {{ game.saveState?.ship_name }}</h3>
      </div>
      <div class="credits">
        <h2>{{ game.saveState?.credits }} CR</h2>
        <h3>DAY {{ game.saveState?.current_day }}</h3>
      </div>
    </div>

    <div class="placeholder-center">
      <h1>SYSTEMS NOMINAL</h1>
      <button class="terminal-btn" @click="abortMission">SAVE & QUIT TO MENU</button>
    </div>
  </div>
</template>

<style scoped>
.game-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
  padding: 40px;
}
.hud-header {
  display: flex;
  justify-content: space-between;
  border-bottom: 2px solid var(--text-color);
  padding-bottom: 20px;
}
h2, h3 { margin: 0; }
h2 { font-size: 48px; text-shadow: var(--text-glow); }
h3 { font-size: 32px; opacity: 0.8; }
.credits { text-align: right; }
.placeholder-center {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 40px;
}
</style>