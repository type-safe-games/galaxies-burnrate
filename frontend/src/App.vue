<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useSettingsStore } from './stores/settingsStore'
import SplashScreen from './components/SplashScreen.vue'
import MainMenu from './components/MainMenu.vue'
import OptionsMenu from './components/OptionsMenu.vue'
import NewGameSetup from './components/NewGameSetup.vue'
import LoadGameMenu from './components/LoadGameMenu.vue'
import GameUI from './components/GameUI.vue'

const settings = useSettingsStore()
const currentView = ref('splash')

onMounted(() => {
  settings.applyTheme()

  if (settings.debug.skip_splash) {
    currentView.value = 'main_menu'
  }
})

const navigateTo = (view: string) => {
  currentView.value = view
}
</script>

<template>
  <main class="app-container">
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
  /* Let the app take up 100% of the viewport */
  width: 100vw;
  height: 100vh;
  position: relative;
  background-color: var(--bg-color);
  overflow: hidden;
}
</style>