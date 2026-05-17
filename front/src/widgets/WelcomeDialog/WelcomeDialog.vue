<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="modelValue" class="overlay" @click.self="$emit('update:modelValue', false)">
        <div class="dialog">
          <div class="dialog-header">
            <div class="dialog-title">Добро пожаловать</div>
            <div class="dialog-subtitle">Инструкция по использованию</div>
          </div>

          <div class="steps">
            <div class="step" v-for="(step, i) in steps" :key="i">
              <div class="step-num">{{ i + 1 }}</div>
              <div class="step-body">
                <div class="step-title">{{ step.title }}</div>
                <div class="step-desc">{{ step.desc }}</div>
              </div>
            </div>
          </div>

          <div class="legend">
            <div class="legend-title">Интерпретация результата</div>
            <div class="legend-items">
              <div class="legend-item">
                <div class="legend-dot" style="background: var(--green)"></div>
                <span><b>0–20%</b> — тексты уникальны</span>
              </div>
              <div class="legend-item">
                <div class="legend-dot" style="background: var(--yellow)"></div>
                <span><b>21–50%</b> — частичное совпадение</span>
              </div>
              <div class="legend-item">
                <div class="legend-dot" style="background: var(--red)"></div>
                <span><b>51–100%</b> — высокое сходство</span>
              </div>
            </div>
          </div>

          <button class="close-btn" @click="$emit('update:modelValue', false)">
            Понятно, начать работу →
          </button>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
defineProps({
  modelValue: { type: Boolean, default: true }
})

defineEmits(['update:modelValue'])

const steps = [
  {
    title: 'Введите или загрузите тексты',
    desc: 'В каждом блоке выберите режим «ввод» для ручного набора или «файл» для загрузки .txt документа. Поддерживается перетаскивание файла в область блока.'
  },
  {
    title: 'Нажмите «Рассчитать»',
    desc: 'Кнопка становится активной когда оба текста заполнены. После нажатия запрос отправляется на сервер — дождитесь завершения анализа.'
  },
  {
    title: 'Изучите результат',
    desc: 'В диалоге отображается общая оценка сходства в виде кругового индикатора, а также отдельные показатели лексического и структурного анализа.'
  },
]
</script>

<style scoped>
.overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.72);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 200;
  padding: 1.5rem;
}

.dialog {
  background: var(--bg-surface);
  border: 0.5px solid var(--border-hover);
  border-radius: var(--radius-lg);
  width: 100%;
  max-width: 560px;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.dialog-header {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.dialog-title {
  font-size: 20px;
  font-weight: 500;
  color: var(--text);
}

.dialog-subtitle {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-muted);
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.steps {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.step {
  display: flex;
  gap: 1rem;
  align-items: flex-start;
}

.step-num {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: var(--accent-dim);
  border: 0.5px solid var(--accent);
  color: var(--accent);
  font-family: var(--mono);
  font-size: 12px;
  font-weight: 500;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  margin-top: 2px;
}

.step-title {
  font-size: 14px;
  font-weight: 500;
  color: var(--text);
  margin-bottom: 3px;
}

.step-desc {
  font-size: 13px;
  color: var(--text-muted);
  line-height: 1.6;
}

.legend {
  background: var(--bg-elevated);
  border-radius: var(--radius);
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.6rem;
}

.legend-title {
  font-family: var(--mono);
  font-size: 10px;
  color: var(--text-dim);
  letter-spacing: 0.08em;
  text-transform: uppercase;
  margin-bottom: 0.25rem;
}

.legend-items {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  font-size: 13px;
  color: var(--text-muted);
}

.legend-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  flex-shrink: 0;
}

.close-btn {
  font-family: var(--mono);
  font-size: 12px;
  letter-spacing: 0.06em;
  padding: 0.75rem 1.5rem;
  background: var(--accent);
  color: #0e0e0f;
  border: none;
  border-radius: var(--radius);
  cursor: pointer;
  font-weight: 500;
  transition: opacity 0.2s;
  align-self: flex-end;
}

.close-btn:hover {
  opacity: 0.88;
}

.fade-enter-active, .fade-leave-active { transition: opacity 0.2s; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
