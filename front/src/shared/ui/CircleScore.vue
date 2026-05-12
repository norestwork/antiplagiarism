<template>
  <div class="circle-wrap">
    <svg class="circle-svg" width="140" height="140" viewBox="0 0 140 140">
      <circle class="circle-bg" cx="70" cy="70" r="56" />
      <circle
        class="circle-fg"
        cx="70" cy="70" r="56"
        :stroke="color"
        :stroke-dasharray="circumference"
        :stroke-dashoffset="dashOffset"
      />
      <text class="circle-text" x="70" y="70">{{ toPct(value) }}</text>
    </svg>
    <span class="circle-label">общая схожесть</span>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { toPct, getScoreColor } from '@/shared/lib/score.js'

const props = defineProps({
  value: { type: Number, default: 0 },
})

const circumference = 2 * Math.PI * 56

const dashOffset = computed(() => circumference * (1 - props.value))
const color = computed(() => getScoreColor(props.value))
</script>

<style scoped>
.circle-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
}

.circle-svg {
  transform: rotate(-90deg);
}

.circle-bg {
  fill: none;
  stroke: var(--bg-elevated);
  stroke-width: 8;
}

.circle-fg {
  fill: none;
  stroke-width: 8;
  stroke-linecap: round;
  transition: stroke-dashoffset 0.9s cubic-bezier(0.4, 0, 0.2, 1), stroke 0.4s;
}

.circle-text {
  font-family: var(--mono);
  font-size: 22px;
  font-weight: 500;
  fill: var(--text);
  text-anchor: middle;
  dominant-baseline: middle;
  transform: rotate(90deg);
  transform-origin: center;
}

.circle-label {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-muted);
  letter-spacing: 0.06em;
}
</style>
