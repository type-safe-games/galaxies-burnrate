<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { GetAvailableSaves } from '../../wailsjs/go/main/App'
import { useGameStore } from '../stores/gameStore'

const emit = defineEmits(['navigate'])
const game = useGameStore()

const savedGames = ref<any[]>([])
const isLoading = ref(false)

onMounted(async () => {
  // Call the Go backend to scan the hard drive
  const results = await GetAvailableSaves()
  // Sort them so the newest saves are at the top (based on the Unix timestamp in the SaveID)
  savedGames.value = results ? results.sort((a: any, b: any) => b.save_id.localeCompare(a.save_id)) : []
})

const loadSave = async (saveId: string) => {
  isLoading.value = true
  await game.loadGameFromId(saveId)
  emit('navigate', 'game')
}

const goBack = () => {
  emit('navigate', 'main_menu')
}
</script>

<template>
  <div class="load-container">
    <h2>RESTORE SECTOR DATA</h2>
    
    <div v-if="isLoading" class="loading-wrapper">
      <h3>DECRYPTING ARCHIVES...</h3>
    </div>

    <div v-else class="saves-list">
      <div v-if="savedGames.length === 0" class="no-saves">
        <p>NO ARCHIVES FOUND ON LOCAL STORAGE.</p>
      </div>

      <button 
        v-for="save in savedGames" 
        :key="save.save_id" 
        class="terminal-btn save-slot"
        @click="loadSave(save.save_id)"
      >
        <div class="save-header">
          <span class="captain-name">CPT. {{ save.captain_name }}</span>
          <span class="day-count">DAY {{ save.current_day }}</span>
        </div>
        <div class="save-details">
          <span>VESSEL: {{ save.ship_name }}</span>
          <span>FUNDS: {{ save.credits }} CR</span>
        </div>
      </button>
    </div>

    <nav class="button-group">
      <button class="terminal-btn" @click="goBack" :disabled="isLoading">ABORT</button>
    </nav>
  </div>
</template>

<style scoped>
.load-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}

h2 {
  font-size: 54px;
  margin-bottom: 40px;
  text-shadow: var(--text-glow);
  letter-spacing: 4px;
}

.saves-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 700px;
  max-height: 500px;
  overflow-y: auto; /* Allows scrolling if they have dozens of saves */
  padding-right: 15px;
}

/* Custom Scrollbar for Terminal aesthetic */
.saves-list::-webkit-scrollbar { width: 10px; }
.saves-list::-webkit-scrollbar-track { background: var(--bg-color); border: 1px solid var(--text-color); }
.saves-list::-webkit-scrollbar-thumb { background: var(--text-color); }

.no-saves {
  text-align: center;
  font-size: 32px;
  opacity: 0.5;
  margin: 50px 0;
}

.save-slot {
  display: flex;
  flex-direction: column;
  text-align: left;
  padding: 20px;
}

.save-header {
  display: flex;
  justify-content: space-between;
  font-size: 36px;
  margin-bottom: 10px;
}

.save-details {
  display: flex;
  justify-content: space-between;
  font-size: 24px;
  opacity: 0.8;
}

.loading-wrapper {
  margin: 100px 0;
  text-align: center;
  animation: pulse 1.5s infinite;
}

.button-group {
  margin-top: 50px;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}
</style>