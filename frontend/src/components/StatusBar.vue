<script setup lang="ts">
import { useGameStore } from '../stores/gameStore'

const emit = defineEmits(['navigate'])
const game = useGameStore()

const abortMission = () => {
  emit('navigate', 'main_menu')
}
</script>

<template>
  <aside class="sidebar">
    
    <div class="sidebar-header">
      <div class="status-row">
        <div class="status-left">
          <span class="dim">Captain:</span> <span class="bright">{{ game.saveState?.captain_name }}</span>
        </div>
        <div class="status-right bright">
          {{ game.saveState?.credits?.toLocaleString() }} CR
        </div>
      </div>

      <div class="status-row">
        <div class="status-left">
          <span class="dim">Ship:</span> <span class="bright">{{ game.saveState?.ship_name }}</span>
        </div>
        <div class="status-right bright">
          Day {{ game.saveState?.current_day }}
        </div>
      </div>
    </div>

    <div class="target-info-section">
      </div>

    <div class="sidebar-footer">
      <h3>SYSTEMS NOMINAL</h3>
      <button class="terminal-btn" @click="abortMission">SAVE & QUIT</button>
    </div>

  </aside>
</template>

<style scoped>
.sidebar {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  border-right: 2px solid var(--text-color);
  padding: 30px; 
  box-shadow: 5px 0 20px rgba(0, 0, 0, 0.9);
  z-index: 10;
}

/* --- COMPACT HEADER STYLES --- */
.sidebar-header {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding-bottom: 20px;
  border-bottom: 2px dashed var(--text-color);
}

.status-row {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  font-size: 28px;
}

.status-left {
  text-align: left;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  padding-right: 15px; 
}

.status-right {
  text-align: right;
  white-space: nowrap;
}

.dim {
  opacity: 0.6;
}

.bright {
  text-shadow: var(--text-glow);
}

.target-info-section {
  flex: 1; 
  margin: 20px 0;
}

.sidebar-footer {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.sidebar-footer h3 {
  margin: 0;
  font-size: 28px;
  text-align: center;
  padding-bottom: 20px;
  border-bottom: 2px dashed var(--text-color);
}

.terminal-btn {
  width: 100%;
}
</style>