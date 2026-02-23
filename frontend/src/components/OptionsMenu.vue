<script setup lang="ts">
import { watch } from 'vue'
import { useSettingsStore } from '../stores/settingsStore'

const emit = defineEmits(['navigate'])
const settings = useSettingsStore()

// Real-time theme updates
watch(() => settings.theme, () => {
  settings.applyTheme()
}, { deep: true })

// Expanded Resolutions Array covering Legacy, Standard, Productivity, and Ultrawide
const availableResolutions = [
  // 4:3 Legacy / CRT
  { label: "640 x 480 (4:3)", val: { w: 640, h: 480 } },
  { label: "800 x 600 (4:3)", val: { w: 800, h: 600 } },
  { label: "1024 x 768 (4:3)", val: { w: 1024, h: 768 } },
  { label: "1152 x 864 (4:3)", val: { w: 1152, h: 864 } },
  
  // 16:10 Productivity / Laptops
  { label: "1280 x 800 (16:10)", val: { w: 1280, h: 800 } },
  { label: "1440 x 900 (16:10)", val: { w: 1440, h: 900 } },
  { label: "1920 x 1200 (16:10)", val: { w: 1920, h: 1200 } },
  { label: "2560 x 1600 (16:10)", val: { w: 2560, h: 1600 } },

  // 16:9 Standard HD / Modern
  { label: "1280 x 720 (16:9)", val: { w: 1280, h: 720 } },
  { label: "1366 x 768 (16:9)", val: { w: 1366, h: 768 } }, // Extremely common budget laptop res
  { label: "1600 x 900 (16:9)", val: { w: 1600, h: 900 } },
  { label: "1920 x 1080 (16:9)", val: { w: 1920, h: 1080 } },
  { label: "2560 x 1440 (16:9)", val: { w: 2560, h: 1440 } },
  { label: "3840 x 2160 (16:9)", val: { w: 3840, h: 2160 } },
  { label: "7680 x 4320 (16:9)", val: { w: 7680, h: 4320 } },

  // 21:9 & 32:9 Ultrawides
  { label: "2560 x 1080 (21:9)", val: { w: 2560, h: 1080 } },
  { label: "3440 x 1440 (21:9)", val: { w: 3440, h: 1440 } },
  { label: "5120 x 1440 (32:9)", val: { w: 5120, h: 1440 } }
]

const displayModes = [
  { label: "WINDOWED", val: "windowed" },
  { label: "FULLSCREEN", val: "fullscreen" },
  { label: "BORDERLESS", val: "borderless" }
]

const applySettings = () => {
  settings.applyDisplaySettings()
}

const goBack = () => {
  emit('navigate', 'main_menu')
}
</script>

<template>
  <div class="options-container">
    <h2>SYSTEM CONFIGURATION</h2>
    
    <div class="settings-grid">
      
      <div class="setting-row">
        <label>RESOLUTION</label>
        <select v-model="settings.display.resolution" class="terminal-input">
          <option v-for="res in availableResolutions" :key="res.label" :value="res.val">
            {{ res.label }}
          </option>
        </select>
      </div>

      <div class="setting-row">
        <label>DISPLAY MODE</label>
        <div class="toggle-group">
          <span 
            v-for="mode in displayModes" 
            :key="mode.val"
            class="custom-toggle"
            @click="settings.display.mode = mode.val"
          >
            [{{ settings.display.mode === mode.val ? 'X' : ' ' }}] {{ mode.label }}
          </span>
        </div>
      </div>

      <div class="setting-row">
        <label>BACKGROUND COLOR</label>
        <div class="color-picker-wrapper">
          <input type="text" v-model="settings.theme.background" class="terminal-input hex-input" maxlength="7">
          <div class="color-preview-box">
            <input type="color" v-model="settings.theme.background" class="invisible-color-picker">
          </div>
        </div>
      </div>

      <div class="setting-row">
        <label>TERMINAL TEXT</label>
        <div class="color-picker-wrapper">
          <input type="text" v-model="settings.theme.text" class="terminal-input hex-input" maxlength="7">
          <div class="color-preview-box" :style="{ backgroundColor: settings.theme.text }">
            <input type="color" v-model="settings.theme.text" class="invisible-color-picker">
          </div>
        </div>
      </div>

    </div>

    <nav class="button-group">
      <button class="terminal-btn" @click="applySettings">APPLY DISPLAY</button>
      <button class="terminal-btn" @click="goBack">RETURN</button>
    </nav>
  </div>
</template>

<style scoped>
.options-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}

h2 {
  font-size: 54px;
  margin-bottom: 50px;
  text-shadow: var(--text-glow);
  letter-spacing: 4px;
}

.settings-grid {
  display: flex;
  flex-direction: column;
  gap: 30px;
  width: 750px;
}

.setting-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 32px;
}

.terminal-input {
  background: var(--bg-color);
  color: var(--text-color);
  border: 2px solid var(--text-color);
  padding: 8px 12px;
  font-size: 28px;
  cursor: pointer;
  outline: none;
  min-width: 280px;
}

.terminal-input:focus {
  box-shadow: 0 0 10px var(--text-glow);
}

/* CUSTOM TOGGLE BRACKETS */
.toggle-group {
  display: flex;
  gap: 15px;
  font-size: 28px;
}

.custom-toggle {
  cursor: pointer;
  transition: all 0.2s ease;
}

.custom-toggle:hover {
  text-shadow: var(--text-glow);
}

/* CUSTOM COLOR PICKER */
.color-picker-wrapper {
  display: flex;
  align-items: center;
  gap: 15px;
}

.hex-input {
  width: 140px;
  min-width: 140px;
  text-align: center;
}

.color-preview-box {
  width: 48px;
  height: 48px;
  border: 2px solid var(--text-color);
  position: relative;
  overflow: hidden;
  cursor: pointer;
  background-color: var(--bg-color); /* Overridden inline for text color */
}

/* Hide the ugly native OS color picker, but keep it clickable behind the preview box */
.invisible-color-picker {
  opacity: 0;
  position: absolute;
  top: -10px;
  left: -10px;
  width: 100px;
  height: 100px;
  cursor: pointer;
}

.button-group {
  margin-top: 80px;
  display: flex;
  gap: 40px;
}
</style>