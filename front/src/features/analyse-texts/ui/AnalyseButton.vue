<template>
  <button
    class="analyse-btn"
    :disabled="disabled"
    @click="$emit('click')"
  >
    Рассчитать →
  </button>
</template>

<script setup>
defineProps({
  disabled: { type: Boolean, default: false },
})

defineEmits(['click'])
</script>

<style scoped>
.analyse-btn {
  font-family: var(--mono);
  font-size: 12px;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  padding: 0.75rem 2rem;
  background: var(--accent);
  color: #0e0e0f;
  border: none;
  border-radius: var(--radius);
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s;
}

.analyse-btn:hover:not(:disabled) {
  opacity: 0.88;
  transform: translateY(-1px);
}

.analyse-btn:active:not(:disabled) {
  transform: translateY(0);
}

.analyse-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}
</style>
