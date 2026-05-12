import { ref } from 'vue'

export function useFileReader() {
  const text = ref('')
  const fileName = ref('')
  const isDragging = ref(false)

  function readFile(file) {
    if (!file || !file.name.endsWith('.txt')) return
    fileName.value = file.name
    const reader = new FileReader()
    reader.onload = (e) => { text.value = e.target.result }
    reader.readAsText(file, 'UTF-8')
  }

  function onFileChange(e) {
    readFile(e.target.files[0])
  }

  function onDrop(e) {
    isDragging.value = false
    readFile(e.dataTransfer.files[0])
  }

  function onDragOver() { isDragging.value = true }
  function onDragLeave() { isDragging.value = false }

  function reset() {
    text.value = ''
    fileName.value = ''
  }

  return { text, fileName, isDragging, onFileChange, onDrop, onDragOver, onDragLeave, reset }
}
