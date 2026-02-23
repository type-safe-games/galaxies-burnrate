<script setup lang="ts">
import { useGameStore } from '../stores/gameStore'
import StarMap from './StarMap.vue'

const emit = defineEmits(['navigate'])
const game = useGameStore()

const abortMission = () => {
  emit('navigate', 'main_menu')
}
</script>

<template>
  <div class="game-container">
    
    <StarMap />

    <div class="hud-layer">
      
      <div class="hud-header panel">
        <div class="captain-info">
          <h2>CPT. {{ game.saveState?.captain_name }}</h2>
          <h3>VESSEL: {{ game.saveState?.ship_name }}</h3>
        </div>
        <div class="credits-info">
          <h2>{{ game.saveState?.credits }} CR</h2>
          <h3>DAY {{ game.saveState?.current_day }}</h3>
        </div>
      </div>

      <div class="hud-footer">
        <div class="panel bottom-panel">
          <h3>SYSTEMS NOMINAL</h3>
        </div>
        
        <div class="action-buttons">
          <button class="terminal-btn" @click="abortMission">SAVE & QUIT</button>
        </div>
      </div>

    </div>

  </div>
</template>

<style scoped>
.game-container {
  position: relative;
  width: 100%;
  height: 100%;
  background-color: var(--bg-color); /* Fallback */
}

/* Overlay Wrapper */
.hud-layer {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  padding: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  
  /* Critical: Lets mouse clicks pass through the empty center to the Canvas below */
  pointer-events: none; 
  z-index: 10;
}

/* UI Panels */
.panel {
  background: var(--bg-color);
  border: 2px solid var(--text-color);
  padding: 20px 30px;
  box-shadow: 0 0 15px rgba(0, 0, 0, 0.8);
  /* Re-enable clicks for actual UI elements */
  pointer-events: auto; 
}

.hud-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.captain-info h2, .credits-info h2 {
  font-size: 48px;
  margin: 0;
  text-shadow: var(--text-glow);
}

.captain-info h3, .credits-info h3 {
  font-size: 32px;
  margin: 5px 0 0 0;
  opacity: 0.8;
}

.credits-info {
  text-align: right;
}

.hud-footer {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
}

.bottom-panel {
  min-width: 400px;
}

.bottom-panel h3 {
  margin: 0;
  font-size: 32px;
}

.action-buttons {
  pointer-events: auto;
}
</style>