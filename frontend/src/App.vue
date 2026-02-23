<script setup lang="ts">
import { ref, onMounted } from 'vue'
import SplashScreen from './components/SplashScreen.vue'
import MainMenu from './components/MainMenu.vue'
import OptionsMenu from './components/OptionsMenu.vue'
import GameUI from './components/GameUI.vue'

// Read this from your Pinia settings store on mount
const skipSplash = ref(true) // Set to true based on dev_mode

// The Application State Machine
const currentView = ref('splash')

onMounted(() => {
  if (skipSplash.value) {
    currentView.value = 'main_menu'
  } else {
    // Normal splash screen flow
    setTimeout(() => {
      currentView.value = 'main_menu'
    }, 3000) // Show splash for 3 seconds
  }
})

// Navigation function passed to child components
const navigateTo = (view: string) => {
  currentView.value = view
}
</script>

<template>
  <main class="app-container">
    <SplashScreen v-if="currentView === 'splash'" />
    
    <MainMenu 
      v-if="currentView === 'main_menu'" 
      @navigate="navigateTo" 
    />
    
    <OptionsMenu 
      v-if="currentView === 'options'" 
      @navigate="navigateTo" 
    />
    
    <GameUI 
      v-if="currentView === 'game'" 
      @navigate="navigateTo" 
    />
  </main>
</template>

<style>
/* Global CSS reset and retro terminal font setup goes here */
.app-container {
  width: 100vw;
  height: 100vh;
  background-color: black;
  color: #33ff33; /* Classic terminal green */
  overflow: hidden; /* Prevent scrollbars in the native window */
}
</style>