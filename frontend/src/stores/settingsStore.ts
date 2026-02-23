// frontend/src/stores/settingsStore.ts
import { defineStore } from 'pinia'
import { SetDisplayMode, SetResolution } from '../../wailsjs/go/main/App'

export const useSettingsStore = defineStore('settings', {
  state: () => ({
    display: {
      resolution: { w: 1280, h: 720 },
      mode: 'windowed' // 'windowed', 'fullscreen', 'borderless'
    },
    theme: {
      background: '#000000',
      text: '#33ff33'
    },
    audio: {
      master_volume: 0.8
    },
    debug: {
      skip_splash: false,
      dev_mode: true
    }
  }),
  actions: {
    applyDisplaySettings() {
      SetResolution(this.display.resolution.w, this.display.resolution.h)
      SetDisplayMode(this.display.mode)
    },
    applyTheme() {
      document.documentElement.style.setProperty('--bg-color', this.theme.background);
      document.documentElement.style.setProperty('--text-color', this.theme.text);
      
      const hex = this.theme.text.replace('#', '');
      const r = parseInt(hex.substring(0, 2), 16);
      const g = parseInt(hex.substring(2, 4), 16);
      const b = parseInt(hex.substring(4, 6), 16);
      document.documentElement.style.setProperty('--text-glow', `0 0 5px rgba(${r}, ${g}, ${b}, 0.5)`);
    }
  }
})