<template>
  <div class="text-block" :class="{ focused }">
    <div class="block-header">
      <span class="block-label">Текст {{ index }}</span>
      <div class="mode-toggle">
        <button
          class="mode-btn"
          :class="{ active: mode === 'text' }"
          @click="mode = 'text'"
        >ввод</button>
        <button
          class="mode-btn"
          :class="{ active: mode === 'file' }"
          @click="mode = 'file'"
        >файл</button>
      </div>
    </div>

    <div class="block-body">
      <textarea
        v-if="mode === 'text'"
        :value="modelValue"
        :placeholder="`Вставьте текст ${index}...`"
        @input="$emit('update:modelValue', $event.target.value)"
        @focus="focused = true"
        @blur="focused = false"
      />

      <div
        v-else
        class="file-drop"
        :class="{ 'drag-over': fileReader.isDragging.value }"
        @dragover.prevent="fileReader.onDragOver"
        @dragleave="fileReader.onDragLeave"
        @drop.prevent="onDrop"
      >
        <input type="file" accept=".txt" @change="onFileChange" />

        <template v-if="!fileReader.fileName.value">
          <span class="file-icon">📄</span>
          <span class="file-drop-text">Перетащите .txt файл<br>или нажмите для выбора</span>
        </template>

        <template v-else>
          <span class="file-icon">✅</span>
          <span class="file-name">{{ fileReader.fileName.value }}</span>
          <span class="file-drop-text file-drop-text--sm">
            {{ wordCount }} слов загружено
          </span>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useFileReader } from '@/shared/lib/useFileReader.js'

const props = defineProps({
  modelValue: { type: String, default: '' },
  index:      { type: Number, required: true },
})

const emit = defineEmits(['update:modelValue'])

const mode    = ref('text')
const focused = ref(false)
const fileReader = useFileReader()

const wordCount = computed(() =>
  fileReader.text.value.trim()
    ? fileReader.text.value.trim().split(/\s+/).length
    : 0
)

watch(fileReader.text, (val) => emit('update:modelValue', val))

function onFileChange(e) {
  fileReader.onFileChange(e)
}

function onDrop(e) {
  fileReader.onDrop(e)
}
</script>

<style scoped>
.text-block {
  background: var(--bg-surface);
  border: 0.5px solid var(--border);
  border-radius: var(--radius-lg);
  overflow: hidden;
  transition: border-color 0.2s;
}

.text-block.focused {
  border-color: var(--border-hover);
}

.block-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.75rem 1rem;
  border-bottom: 0.5px solid var(--border);
}

.block-label {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-muted);
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.mode-toggle {
  display: flex;
  gap: 2px;
  background: var(--bg-elevated);
  border-radius: 6px;
  padding: 2px;
}

.mode-btn {
  font-family: var(--mono);
  font-size: 10px;
  padding: 3px 8px;
  border-radius: 4px;
  border: none;
  cursor: pointer;
  color: var(--text-muted);
  background: transparent;
  transition: all 0.15s;
  letter-spacing: 0.04em;
}

.mode-btn.active {
  background: var(--bg-surface);
  color: var(--accent);
  border: 0.5px solid var(--border-hover);
}

.block-body {
  padding: 1rem;
}

textarea {
  width: 100%;
  height: 220px;
  background: transparent;
  border: none;
  color: var(--text);
  font-family: var(--serif);
  font-size: 14px;
  line-height: 1.7;
  resize: none;
  outline: none;
}

textarea::placeholder {
  color: var(--text-dim);
}

.file-drop {
  height: 220px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  border: 0.5px dashed var(--border-hover);
  border-radius: var(--radius);
  cursor: pointer;
  transition: all 0.2s;
  position: relative;
}

.file-drop:hover,
.file-drop.drag-over {
  border-color: var(--accent);
  background: var(--accent-dim);
}

.file-drop input[type='file'] {
  position: absolute;
  inset: 0;
  opacity: 0;
  cursor: pointer;
}

.file-icon {
  font-size: 28px;
}

.file-drop-text {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--text-muted);
  text-align: center;
  letter-spacing: 0.04em;
}

.file-drop-text--sm {
  font-size: 10px;
}

.file-name {
  font-family: var(--mono);
  font-size: 11px;
  color: var(--accent);
  background: var(--accent-dim);
  padding: 4px 10px;
  border-radius: 4px;
  max-width: 180px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
