<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useSettingsStore } from './stores/settingsStore'
import SplashScreen from './components/SplashScreen.vue'
import MainMenu from './components/MainMenu.vue'
import OptionsMenu from './components/OptionsMenu.vue'
import NewGameSetup from './components/NewGameSetup.vue'
import LoadGameMenu from './components/LoadGameMenu.vue'
import GameUI from './components/GameUI.vue'

const settings = useSettingsStore()
const currentView = ref('splash')

// --- GAME CANVAS SCALING ---
const scale = ref(1)

const updateScale = () => {
  const designWidth = 1920
  const designHeight = 1080
  const windowWidth = window.innerWidth
  const windowHeight = window.innerHeight
  
  const scaleX = windowWidth / designWidth
  const scaleY = windowHeight / designHeight
  
  // Maintain 16:9 aspect ratio by scaling to the most restrictive dimension
  scale.value = Math.min(scaleX, scaleY)
}

onMounted(() => {
  settings.applyTheme()

  if (settings.debug.skip_splash) {
    currentView.value = 'main_menu'
  }

  updateScale()
  window.addEventListener('resize', updateScale)
})

onUnmounted(() => {
  window.removeEventListener('resize', updateScale)
})
// ---------------------------

const navigateTo = (view: string) => {
  currentView.value = view
}
</script>

<template>
  <main class="app-container" :style="{ transform: `scale(${scale})` }">
    <SplashScreen v-if="currentView === 'splash'" @navigate="navigateTo" />
    <MainMenu v-if="currentView === 'main_menu'" @navigate="navigateTo" />
    <OptionsMenu v-if="currentView === 'options'" @navigate="navigateTo" />
    <NewGameSetup v-if="currentView === 'new_game_setup'" @navigate="navigateTo" />
    <LoadGameMenu v-if="currentView === 'load_game'" @navigate="navigateTo" />
    <GameUI v-if="currentView === 'game'" @navigate="navigateTo" />
  </main>
</template>

<style scoped>
.app-container {
  width: 1920px;
  height: 1080px;
  transform-origin: center center;
  position: relative;
  background-color: var(--bg-color);
  overflow: hidden;
}
</style>