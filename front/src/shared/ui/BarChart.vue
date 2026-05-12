<template>
  <div class="bars-wrap">
    <div class="bar-row" v-for="bar in bars" :key="bar.label">
      <div class="bar-meta">
        <span class="bar-name">{{ bar.label }}</span>
        <span class="bar-value">{{ toPct(bar.value) }}</span>
      </div>
      <div class="bar-track">
        <div
          class="bar-fill"
          :style="{ width: toPct(bar.value), background: color }"
        />
      </div>
    </div>

    <div class="verdict-pill" :style="{ background: bg, color }">
      <div class="verdict-dot" :style="{ background: color }" />
      {{ verdict }}
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { toPct, getScoreColor, getScoreBg, getVerdict } from '@/shared/lib/score.js'

const props = defineProps({
  lexical:    { type: Number, default: 0 },
  structural: { type: Number, default: 0 },
  general:    { type: Number, default: 0 },
})

const bars = computed(() => [
  { label: 'Лексический',   value: props.lexical },
  { label: 'Структурный',   value: props.structural },
])

const color   = computed(() => getScoreColor(props.general))
const bg      = computed(() => getScoreBg(props.general))
const verdict = computed(() => getVerdict(props.general))
</script>

<style scoped>
.bars-wrap {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.bar-row {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.bar-meta {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
}

.bar-name {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-muted);
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.bar-value {
  font-family: var(--mono);
  font-size: 13px;
  font-weight: 500;
  color: var(--text);
}

.bar-track {
  height: 6px;
  background: var(--bg-elevated);
  border-radius: 3px;
  overflow: hidden;
}

.bar-fill {
  height: 100%;
  border-radius: 3px;
  transition: width 0.9s cubic-bezier(0.4, 0, 0.2, 1) 0.2s;
}

.verdict-pill {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-family: var(--mono);
  font-size: 11px;
  letter-spacing: 0.06em;
  padding: 5px 12px;
  border-radius: 20px;
  margin-top: 0.25rem;
}

.verdict-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  flex-shrink: 0;
}
</style>
