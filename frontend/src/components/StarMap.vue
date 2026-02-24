<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import { useGameStore } from '../stores/gameStore'

const game = useGameStore()

const canvasRef = ref<HTMLCanvasElement | null>(null)
const containerRef = ref<HTMLElement | null>(null)
let resizeObserver: ResizeObserver | null = null

const selectedPlanetId = ref<string | null>(null)

// --- ZOOM LOGIC ---
// Represents the total units visible across the shortest screen dimension.
// 10 = highly zoomed in (5 units each way)
// 40 = default (20 units each way)
// 70 = highly zoomed out (35 units each way)
const zoomLevels = [10, 20, 30, 40, 50, 60, 70]
const currentZoomIndex = ref(3) // Start at index 3 (value: 40)

const zoomIn = () => {
  if (currentZoomIndex.value > 0) currentZoomIndex.value--
}

const zoomOut = () => {
  if (currentZoomIndex.value < zoomLevels.length - 1) currentZoomIndex.value++
}

// --- INTERACTION ---
const handleClick = (e: MouseEvent) => {
  if (!canvasRef.value) return
  
  const rect = canvasRef.value.getBoundingClientRect()
  const mouseX = e.clientX - rect.left
  const mouseY = e.clientY - rect.top

  const w = canvasRef.value.width
  const h = canvasRef.value.height
  const cx = w / 2
  const cy = h / 2
  
  // Use the reactive zoom level
  const scale = Math.min(w, h) / zoomLevels[currentZoomIndex.value]

  const clicked = game.planets?.planets?.find((p: any) => {
    const px = cx + (p.coordinates.x * scale)
    const py = cy - (p.coordinates.y * scale)
    const dist = Math.sqrt(Math.pow(mouseX - px, 2) + Math.pow(mouseY - py, 2))
    return dist < 15 
  })

  selectedPlanetId.value = clicked ? clicked.id : null
  draw()
}

// --- DRAWING ---
const draw = () => {
  if (!canvasRef.value || !containerRef.value) return
  const ctx = canvasRef.value.getContext('2d')
  if (!ctx) return

  const w = canvasRef.value.width
  const h = canvasRef.value.height
  if (w === 0 || h === 0) return

  const style = getComputedStyle(document.documentElement)
  const bgColor = style.getPropertyValue('--bg-color').trim() || '#000000'
  const textColor = style.getPropertyValue('--text-color').trim() || '#33ff33'

  const cx = w / 2
  const cy = h / 2
  
  // Use the reactive zoom level here too
  const scale = Math.min(w, h) / zoomLevels[currentZoomIndex.value]

  ctx.fillStyle = bgColor
  ctx.fillRect(0, 0, w, h)

  // Minor Grid (1 Unit)
  ctx.strokeStyle = textColor
  ctx.globalAlpha = 0.05
  ctx.lineWidth = 1
  const minorGridSize = 1 * scale 
  ctx.beginPath()
  for (let x = cx; x < w; x += minorGridSize) { ctx.moveTo(x, 0); ctx.lineTo(x, h); }
  for (let x = cx; x > 0; x -= minorGridSize) { ctx.moveTo(x, 0); ctx.lineTo(x, h); }
  for (let y = cy; y < h; y += minorGridSize) { ctx.moveTo(0, y); ctx.lineTo(w, y); }
  for (let y = cy; y > 0; y -= minorGridSize) { ctx.moveTo(0, y); ctx.lineTo(w, y); }
  ctx.stroke()

  // Major Grid (5 Units)
  ctx.globalAlpha = 0.15 
  const majorGridSize = 5 * scale 
  ctx.beginPath()
  for (let x = cx; x < w; x += majorGridSize) { ctx.moveTo(x, 0); ctx.lineTo(x, h); }
  for (let x = cx; x > 0; x -= majorGridSize) { ctx.moveTo(x, 0); ctx.lineTo(x, h); }
  for (let y = cy; y < h; y += majorGridSize) { ctx.moveTo(0, y); ctx.lineTo(w, y); }
  for (let y = cy; y > 0; y -= majorGridSize) { ctx.moveTo(0, y); ctx.lineTo(w, y); }
  ctx.stroke()
  
  ctx.globalAlpha = 1.0

  // Center Crosshair
  ctx.strokeStyle = textColor
  ctx.lineWidth = 2
  ctx.beginPath()
  ctx.moveTo(cx, cy - 10); ctx.lineTo(cx, cy + 10)
  ctx.moveTo(cx - 10, cy); ctx.lineTo(cx + 10, cy)
  ctx.stroke()

  if (!game.planets || !game.planets.planets) return

  const currentLocationId = game.saveState?.current_location || 'plt_prime'
  const currentPlanet = game.planets.planets.find((p: any) => p.id === currentLocationId)
  const selectedPlanet = game.planets.planets.find((p: any) => p.id === selectedPlanetId.value)

  // Flight Path
  if (currentPlanet && selectedPlanet && currentPlanet.id !== selectedPlanet.id) {
    const startX = cx + (currentPlanet.coordinates.x * scale)
    const startY = cy - (currentPlanet.coordinates.y * scale)
    const endX = cx + (selectedPlanet.coordinates.x * scale)
    const endY = cy - (selectedPlanet.coordinates.y * scale)

    ctx.beginPath()
    ctx.moveTo(startX, startY)
    ctx.lineTo(endX, endY)
    ctx.setLineDash([8, 8]) 
    ctx.strokeStyle = textColor
    ctx.globalAlpha = 0.5
    ctx.lineWidth = 2
    ctx.stroke()
    ctx.setLineDash([]) 
    ctx.globalAlpha = 1.0
  }

  // Planets
  game.planets.planets.forEach((p: any) => {
    const px = cx + (p.coordinates.x * scale)
    const py = cy - (p.coordinates.y * scale) 
    
    const isCurrent = (p.id === currentLocationId)
    const isSelected = (p.id === selectedPlanetId.value)

    if (isCurrent) {
      ctx.beginPath()
      ctx.arc(px, py, 14, 0, Math.PI * 2)
      ctx.fillStyle = bgColor 
      ctx.fill()
      ctx.strokeStyle = '#ffffff' 
      ctx.lineWidth = 2
      ctx.stroke()
    }

    if (isSelected && !isCurrent) {
      ctx.beginPath()
      ctx.arc(px, py, 10, 0, Math.PI * 2)
      ctx.strokeStyle = textColor
      ctx.lineWidth = 2
      ctx.stroke()
    }

    ctx.beginPath()
    ctx.arc(px, py, isCurrent ? 6 : 4, 0, Math.PI * 2)
    ctx.fillStyle = isCurrent ? '#ffffff' : textColor
    ctx.fill()

    ctx.font = '24px "VT323"'
    ctx.fillStyle = isCurrent ? '#ffffff' : textColor
    ctx.textAlign = 'left'
    ctx.fillText(p.name, px + 15, py + 6) 
  })
}

onMounted(() => {
  if (containerRef.value && canvasRef.value) {
    resizeObserver = new ResizeObserver(() => {
      if (containerRef.value && canvasRef.value) {
        canvasRef.value.width = Math.floor(containerRef.value.clientWidth)
        canvasRef.value.height = Math.floor(containerRef.value.clientHeight)
        draw()
      }
    })
    resizeObserver.observe(containerRef.value)
  }
  setTimeout(draw, 100)
})

onUnmounted(() => {
  if (resizeObserver) resizeObserver.disconnect()
})

watch(() => game.planets, draw, { deep: true })
watch(selectedPlanetId, draw) 

// Force a redraw anytime the user clicks a zoom button
watch(currentZoomIndex, draw)
</script>

<template>
  <div ref="containerRef" class="map-container">
    <canvas ref="canvasRef" @click="handleClick"></canvas>

    <div class="zoom-controls">
      <div class="zoom-label">MAG</div>
      <button class="zoom-btn" @click.stop="zoomIn" :disabled="currentZoomIndex === 0">+</button>
      <div class="zoom-indicator">[{{ currentZoomIndex }}]</div>
      <button class="zoom-btn" @click.stop="zoomOut" :disabled="currentZoomIndex === zoomLevels.length - 1">-</button>
    </div>
  </div>
</template>

<style scoped>
.map-container { 
  width: 100%; 
  height: 100%; 
  position: absolute; 
  top: 0;
  left: 0;
  overflow: hidden; 
  z-index: 0;
}

canvas {
  display: block; 
  width: 100%;
  height: 100%;
  cursor: crosshair;
}

/* ZOOM UI STYLES */
.zoom-controls {
  position: absolute;
  right: 40px;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  background: rgba(0, 0, 0, 0.7);
  padding: 15px 10px;
  border: 1px solid var(--text-color);
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.8);
  /* Enable clicks specifically for this panel so it intercepts them from the canvas */
  pointer-events: auto;
}

.zoom-label {
  font-size: 20px;
  opacity: 0.7;
}

.zoom-indicator {
  font-size: 24px;
  margin: 5px 0;
}

.zoom-btn {
  background: var(--bg-color);
  color: var(--text-color);
  border: 2px solid var(--text-color);
  font-family: 'VT323', monospace;
  font-size: 32px;
  width: 45px;
  height: 45px;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: all 0.15s;
}

.zoom-btn:hover:not(:disabled) {
  background: var(--text-color);
  color: var(--bg-color);
  box-shadow: 0 0 10px var(--text-glow);
}

.zoom-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
  border-color: rgba(255, 255, 255, 0.2);
  color: rgba(255, 255, 255, 0.2);
}
</style>