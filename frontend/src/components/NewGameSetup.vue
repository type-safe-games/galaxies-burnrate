<script setup lang="ts">
import { ref } from 'vue'
import { useGameStore } from '../stores/gameStore'

const emit = defineEmits(['navigate'])
const game = useGameStore()

const captainName = ref('')
const shipName = ref('')
const isGenerating = ref(false)

const handleLaunch = async () => {
  if (!captainName.value || !shipName.value) return
  
  isGenerating.value = true
  // Generates the folder, extracts the YAMLs, and loads them into memory
  const success = await game.startNewGame(captainName.value, shipName.value)
  
  if (success) {
    emit('navigate', 'game')
  } else {
    isGenerating.value = false
    alert("CRITICAL ERROR: Failed to instantiate sector data.")
  }
}

const abortLaunch = () => {
  emit('navigate', 'main_menu')
}
</script>

<template>
  <div class="setup-container">
    <h2>INITIALIZE NEW VESSEL</h2>
    
    <div v-if="!isGenerating" class="form-wrapper">
      <div class="input-group">
        <label>CAPTAIN DESIGNATION:</label>
        <input type="text" v-model="captainName" class="terminal-input" placeholder="ENTER NAME_" maxlength="20" autofocus>
      </div>

      <div class="input-group">
        <label>VESSEL REGISTRY NAME:</label>
        <input type="text" v-model="shipName" class="terminal-input" placeholder="ENTER SHIP NAME_" maxlength="24">
      </div>

      <nav class="button-group">
        <button class="terminal-btn" :disabled="!captainName || !shipName" @click="handleLaunch">CONFIRM REGISTRY</button>
        <button class="terminal-btn" @click="abortLaunch">ABORT</button>
      </nav>
    </div>

    <div v-else class="generating-wrapper">
      <h3>GENERATING SECTOR DATA...</h3>
      <p>Allocating local market variables...</p>
      <p>Hydrating stellar cartography...</p>
    </div>
  </div>
</template>

<style scoped>
.setup-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}
h2 {
  font-size: 54px;
  margin-bottom: 60px;
  text-shadow: var(--text-glow);
  letter-spacing: 4px;
}
.form-wrapper {
  display: flex;
  flex-direction: column;
  gap: 40px;
  width: 600px;
}
.input-group {
  display: flex;
  flex-direction: column;
  gap: 10px;
  font-size: 32px;
}
.terminal-input {
  background: var(--bg-color);
  color: var(--text-color);
  border: 2px solid var(--text-color);
  padding: 15px;
  font-size: 32px;
  outline: none;
  width: 100%;
}
.terminal-input:focus {
  box-shadow: 0 0 15px var(--text-glow);
}
.terminal-input::placeholder {
  text-transform: uppercase; 
  opacity: 0.5;
}
.terminal-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
  box-shadow: none;
  border-color: var(--text-color);
  color: var(--text-color);
  background: transparent;
}
.button-group {
  margin-top: 40px;
  display: flex;
  justify-content: space-between;
}
.generating-wrapper {
  text-align: center;
  animation: pulse 1.5s infinite;
}
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}
</style>