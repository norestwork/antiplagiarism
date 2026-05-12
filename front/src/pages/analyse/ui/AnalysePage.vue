<template>
  <div>
    <AppHeader />

    <main class="main">
      <div class="page-title">Анализ текста</div>
      <div class="page-subtitle">Сравнение двух текстов на схожесть</div>

      <div class="texts-grid">
        <TextBlock v-model="texts[0]" :index="1" />
        <TextBlock v-model="texts[1]" :index="2" />
      </div>

      <div class="bottom-panel">
        <AnalyseButton :disabled="!canAnalyse" @click="onAnalyse" />
      </div>
    </main>

    <ResultDialog
      v-model="analyse.isOpen.value"
      :loading="analyse.loading.value"
      :error="analyse.error.value"
      :result="analyse.result.value"
    />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

import AppHeader    from '@/widgets/AppHeader/AppHeader.vue'
import ResultDialog from '@/widgets/ResultDialog/ResultDialog.vue'
import TextBlock    from '@/shared/ui/TextBlock.vue'
import AnalyseButton from '@/features/analyse-texts/ui/AnalyseButton.vue'
import { useAnalyse } from '@/features/analyse-texts/model/useAnalyse.js'

const texts = ref(['', ''])
const analyse = useAnalyse()

const canAnalyse = computed(() =>
  texts.value[0].trim().length > 0 && texts.value[1].trim().length > 0
)

function onAnalyse() {
  analyse.run(texts.value[0], texts.value[1])
}
</script>

<style scoped>
.main {
  max-width: 960px;
  margin: 0 auto;
  padding: 3rem 2rem 5rem;
}

.page-title {
  font-size: 13px;
  font-family: var(--mono);
  color: var(--text-muted);
  letter-spacing: 0.1em;
  text-transform: uppercase;
  margin-bottom: 0.5rem;
}

.page-subtitle {
  font-size: 28px;
  font-weight: 500;
  color: var(--text);
  margin-bottom: 2.5rem;
  line-height: 1.3;
}

.texts-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
  margin-bottom: 1.5rem;
}

.bottom-panel {
  display: flex;
  justify-content: flex-end;
}
</style>
