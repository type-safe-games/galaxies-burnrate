<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import { useGameStore } from '../stores/gameStore'

const game = useGameStore()

const canvasRef = ref<HTMLCanvasElement | null>(null)
const containerRef = ref<HTMLElement | null>(null)
let resizeObserver: ResizeObserver | null = null

// --- CONSTANTS ---
const MAP_SCALE = 40 // Controls how "zoomed in" the grid is

const draw = () => {
  if (!canvasRef.value || !containerRef.value) return
  const ctx = canvasRef.value.getContext('2d')
  if (!ctx) return

  const w = canvasRef.value.width
  const h = canvasRef.value.height
  if (w === 0 || h === 0) return

  // We'll read the CSS variables to ensure the map respects the user's theme choices
  const style = getComputedStyle(document.documentElement)
  const bgColor = style.getPropertyValue('--bg-color').trim() || '#000000'
  const textColor = style.getPropertyValue('--text-color').trim() || '#33ff33'

  const cx = w / 2
  const cy = h / 2
  const scale = Math.min(w, h) / MAP_SCALE

  // Clear Background
  ctx.fillStyle = bgColor
  ctx.fillRect(0, 0, w, h)

  // Draw Grid (Alpha lowered so it sits in the background)
  ctx.strokeStyle = textColor
  ctx.globalAlpha = 0.15
  ctx.lineWidth = 1
  const gridSize = 5 * scale 
  
  ctx.beginPath()
  for (let x = cx; x < w; x += gridSize) { ctx.moveTo(x, 0); ctx.lineTo(x, h); }
  for (let x = cx; x > 0; x -= gridSize) { ctx.moveTo(x, 0); ctx.lineTo(x, h); }
  for (let y = cy; y < h; y += gridSize) { ctx.moveTo(0, y); ctx.lineTo(w, y); }
  for (let y = cy; y > 0; y -= gridSize) { ctx.moveTo(0, y); ctx.lineTo(w, y); }
  ctx.stroke()
  
  ctx.globalAlpha = 1.0

  // Center Crosshair (0,0)
  ctx.strokeStyle = textColor
  ctx.lineWidth = 2
  ctx.beginPath()
  ctx.moveTo(cx, cy - 10); ctx.lineTo(cx, cy + 10)
  ctx.moveTo(cx - 10, cy); ctx.lineTo(cx + 10, cy)
  ctx.stroke()

  // Draw Planets from YAML Data
  if (!game.planets || !game.planets.planets) return

  game.planets.planets.forEach((p: any) => {
    // Map YAML {x, y} to Canvas coordinates
    const px = cx + (p.coordinates.x * scale)
    const py = cy - (p.coordinates.y * scale) // Invert Y so positive is "up"

    // Planet Node
    ctx.beginPath()
    ctx.arc(px, py, 4, 0, Math.PI * 2)
    ctx.fillStyle = textColor
    ctx.fill()

    // Planet Label
    ctx.font = '24px "VT323"'
    ctx.fillStyle = textColor
    ctx.textAlign = 'left'
    ctx.fillText(p.name.toUpperCase(), px + 12, py + 6)
  })
}

// --- RESIZE LOGIC ---
onMounted(() => {
  if (containerRef.value && canvasRef.value) {
    resizeObserver = new ResizeObserver(() => {
      if (containerRef.value && canvasRef.value) {
        // Force integer pixels for sharpness
        canvasRef.value.width = Math.floor(containerRef.value.clientWidth)
        canvasRef.value.height = Math.floor(containerRef.value.clientHeight)
        draw()
      }
    })
    resizeObserver.observe(containerRef.value)
  }
  // Short timeout to ensure CSS fonts/variables are loaded before first draw
  setTimeout(draw, 100)
})

onUnmounted(() => {
  if (resizeObserver) resizeObserver.disconnect()
})

// Redraw map if data hydrates
watch(() => game.planets, draw, { deep: true })
</script>

<template>
  <div ref="containerRef" class="map-container">
    <canvas ref="canvasRef"></canvas>
  </div>
</template>

<style scoped>
.map-container { 
  width: 100%; 
  height: 100%; 
  position: absolute; /* Sits behind the UI */
  top: 0;
  left: 0;
  overflow: hidden; 
  z-index: 0;
}

canvas {
  display: block; 
  width: 100%;
  height: 100%;
}
</style>