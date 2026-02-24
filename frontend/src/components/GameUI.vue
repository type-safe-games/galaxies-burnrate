<script setup lang="ts">
import { ref, onUnmounted } from 'vue'
import StarMap from './StarMap.vue'
import StatusBar from './StatusBar.vue' // Updated Name

const emit = defineEmits(['navigate'])

// --- RESIZE LOGIC ---
const sidebarWidth = ref(450) // Starting default width
const isResizing = ref(false)
const MIN_WIDTH = 350

const startResize = (e: MouseEvent) => {
  isResizing.value = true
  // Force the cursor to stay as a resize arrow globally while dragging
  document.body.style.cursor = 'col-resize' 
  window.addEventListener('mousemove', doResize)
  window.addEventListener('mouseup', stopResize)
}

const doResize = (e: MouseEvent) => {
  if (!isResizing.value) return
  
  // Cap the maximum width at exactly 50% of the player's current window
  const MAX_WIDTH = window.innerWidth / 2 
  let newWidth = e.clientX

  // Clamp the values
  if (newWidth < MIN_WIDTH) newWidth = MIN_WIDTH
  if (newWidth > MAX_WIDTH) newWidth = MAX_WIDTH

  sidebarWidth.value = newWidth
}

const stopResize = () => {
  isResizing.value = false
  document.body.style.cursor = '' // Release the cursor
  window.removeEventListener('mousemove', doResize)
  window.removeEventListener('mouseup', stopResize)
}

// Failsafe: Ensure event listeners are killed if the component unmounts
onUnmounted(() => {
  window.removeEventListener('mousemove', doResize)
  window.removeEventListener('mouseup', stopResize)
  document.body.style.cursor = ''
})
</script>

<template>
  <div class="game-layout">
    
    <div class="sidebar-wrapper" :style="{ width: sidebarWidth + 'px' }">
      <StatusBar @navigate="(view) => emit('navigate', view)" />
    </div>

    <div 
      class="resizer-seam" 
      @mousedown.prevent="startResize"
      :class="{ active: isResizing }"
    >
      <div class="resizer-handle top">||</div>
      <div class="resizer-handle bottom">||</div>
    </div>

    <main class="map-wrapper">
      <StarMap />
    </main>

    <div v-if="isResizing" class="drag-shield"></div>

  </div>
</template>

<style scoped>
.game-layout {
  display: flex;
  width: 100vw;
  height: 100vh;
  background-color: var(--bg-color);
  position: relative;
}

.sidebar-wrapper {
  height: 100%;
  flex-shrink: 0; /* Prevents the flex layout from squishing the sidebar */
}

/* --- RESIZER STYLES --- */
.resizer-seam {
  width: 16px;
  background-color: transparent;
  cursor: col-resize;
  display: flex;
  flex-direction: column;
  justify-content: space-between; /* Pushes one handle to the top, one to the bottom */
  align-items: center;
  z-index: 20;
  
  /* Pull the seam slightly left to perfectly overlap the StatusBar's right border */
  margin-left: -8px; 
  margin-right: -8px;
}

.resizer-handle {
  color: var(--text-color);
  font-size: 14px;
  font-weight: bold;
  padding: 20px 0;
  opacity: 0.3;
  transition: all 0.2s;
  user-select: none;
}

/* Glow up when hovered or actively dragging */
.resizer-seam:hover .resizer-handle,
.resizer-seam.active .resizer-handle {
  opacity: 1;
  text-shadow: var(--text-glow);
}

/* --- MAP STYLES --- */
.map-wrapper {
  flex: 1; 
  position: relative;
  overflow: hidden;
}

/* CRITICAL: When dragging fast, the mouse can leave the resizer and enter the Canvas. 
  Canvas elements notoriously swallow mouse events. This invisible shield covers the 
  whole screen during a drag so your window 'mousemove' events fire perfectly.
*/
.drag-shield {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 9999;
  cursor: col-resize;
}
</style>