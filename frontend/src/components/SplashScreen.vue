<script setup lang="ts">
import { onMounted, ref } from 'vue'

const emit = defineEmits(['navigate'])
const isFading = ref(false)

onMounted(() => {
  // Wait 4 seconds for the animations to play out
  setTimeout(() => {
    isFading.value = true // Trigger the fade out class
    
    // Wait 1 second for the fade to finish, then go to menu
    setTimeout(() => {
      emit('navigate', 'main_menu')
    }, 1000)
  }, 4000) 
})
</script>

<template>
  <div class="splash-container" :class="{ 'fade-out': isFading }">
    <div class="animation-wrapper">
      <h1 class="title typing-effect">GALAXIES: BURN RATE<span class="cursor">_</span></h1>
      <h2 class="studio fade-in-effect">A TYPE SAFE GAMES PRODUCTION</h2>
    </div>
  </div>
</template>

<style scoped>
.splash-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  background-color: var(--bg-color);
  transition: opacity 1s ease-in-out;
  opacity: 1;
}

.fade-out {
  opacity: 0;
}

.animation-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.title {
  font-size: 84px;
  margin-bottom: 10px;
  text-shadow: var(--text-glow);
  white-space: nowrap;
  overflow: hidden;
  margin: 0 auto;
}

/* CSS Typewriter effect */
.typing-effect {
  width: 0;
  animation: typing 1.5s steps(20, end) forwards;
}

.studio {
  font-size: 36px;
  opacity: 0;
  margin-top: 20px;
}

/* Fade in the studio text AFTER the typing finishes */
.fade-in-effect {
  animation: fadeIn 1s ease-in-out 1.5s forwards;
}

.cursor {
  animation: blink 1s step-end infinite;
}

@keyframes typing {
  from { width: 0; }
  to { width: 100%; }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 0.7; }
}

@keyframes blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}
</style>