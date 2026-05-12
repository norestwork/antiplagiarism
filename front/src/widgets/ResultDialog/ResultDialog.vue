<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="modelValue" class="overlay" @click.self="$emit('update:modelValue', false)">
        <div class="dialog">
          <button class="dialog-close" @click="$emit('update:modelValue', false)">✕</button>

          <div v-if="loading" class="loading-state">
            <div class="spinner-ring" />
            <div class="loading-text">Анализируем тексты...</div>
          </div>

          <div v-else-if="error" class="error-state">
            <div class="error-icon">⚠</div>
            <div class="error-text">{{ error }}</div>
          </div>

          <div v-else>
            <div class="result-header">
              <div class="result-label">Результат анализа</div>
              <div class="result-title">Оценка схожести текстов</div>
            </div>
            <div class="result-body">
              <CircleScore :value="result.general" />
              <BarChart
                :general="result.general"
                :lexical="result.lexical"
                :structural="result.structural"
              />
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import CircleScore from '@/shared/ui/CircleScore.vue'
import BarChart from '@/shared/ui/BarChart.vue'

defineProps({
  modelValue: { type: Boolean, default: false },
  loading:    { type: Boolean, default: false },
  error:      { type: String,  default: null },
  result: {
    type: Object,
    default: () => ({ general: 0, structural: 0, lexical: 0 }),
  },
})

defineEmits(['update:modelValue'])
</script>

<style scoped>
.overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.72);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
  padding: 1.5rem;
}

.dialog {
  background: var(--bg-surface);
  border: 0.5px solid var(--border-hover);
  border-radius: var(--radius-lg);
  width: 100%;
  max-width: 680px;
  padding: 2rem;
  position: relative;
}

.dialog-close {
  position: absolute;
  top: 1rem;
  right: 1rem;
  background: transparent;
  border: none;
  color: var(--text-muted);
  cursor: pointer;
  font-size: 18px;
  padding: 4px;
  transition: color 0.15s;
}

.dialog-close:hover { color: var(--text); }

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 1.25rem;
  padding: 3rem 0;
}

.spinner-ring {
  width: 48px;
  height: 48px;
  border: 2px solid var(--border-hover);
  border-top-color: var(--accent);
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

.loading-text {
  font-family: var(--mono);
  font-size: 12px;
  color: var(--text-muted);
  letter-spacing: 0.06em;
}

.error-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
  padding: 3rem 0;
}

.error-icon {
  font-size: 28px;
  color: var(--red);
}

.error-text {
  font-family: var(--mono);
  font-size: 12px;
  color: var(--text-muted);
  text-align: center;
}

.result-header {
  margin-bottom: 1.75rem;
}

.result-label {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-muted);
  letter-spacing: 0.1em;
  text-transform: uppercase;
  margin-bottom: 0.25rem;
}

.result-title {
  font-size: 18px;
  font-weight: 500;
  color: var(--text);
}

.result-body {
  display: grid;
  grid-template-columns: 200px 1fr;
  gap: 2.5rem;
  align-items: center;
}

.fade-enter-active, .fade-leave-active { transition: opacity 0.2s; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
