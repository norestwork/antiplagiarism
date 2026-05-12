import { ref } from 'vue'
import { analyseTexts } from '@/shared/api/analyse.js'

export function useAnalyse() {
  const isOpen   = ref(false)
  const loading  = ref(false)
  const error    = ref(null)
  const result   = ref({ general: 0, structural: 0, lexical: 0 })

  async function run(text1, text2) {
    isOpen.value  = true
    loading.value = true
    error.value   = null
    result.value  = { general: 0, structural: 0, lexical: 0 }

    try {
      result.value = await analyseTexts(text1, text2)
    } catch (e) {
      error.value = e.message
    } finally {
      loading.value = false
    }
  }

  return { isOpen, loading, error, result, run }
}
