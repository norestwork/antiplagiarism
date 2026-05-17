<template>
  <header class="header">
    <div class="logo">Антиплагиат</div>
    <button class="theme-btn" @click="toggle" :title="isDark ? 'Светлая тема' : 'Тёмная тема'">
      <span v-if="isDark">☀</span>
      <span v-else>☾</span>
    </button>
  </header>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const isDark = ref(true)

onMounted(() => {
  const saved = localStorage.getItem('theme')
  if (saved === 'light') {
    isDark.value = false
    document.documentElement.setAttribute('data-theme', 'light')
  }
})

function toggle() {
  isDark.value = !isDark.value
  const theme = isDark.value ? 'dark' : 'light'
  localStorage.setItem('theme', theme)
  if (isDark.value) {
    document.documentElement.removeAttribute('data-theme')
  } else {
    document.documentElement.setAttribute('data-theme', 'light')
  }
}
</script>

<style scoped>
.header {
  border-bottom: 0.5px solid var(--border);
  padding: 0 2.5rem;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: var(--bg);
  position: sticky;
  top: 0;
  z-index: 10;
  transition: background 0.25s, border-color 0.25s;
}

.logo {
  font-family: var(--mono);
  font-size: 13px;
  font-weight: 500;
  color: var(--accent);
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.theme-btn {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  border: 0.5px solid var(--border-hover);
  background: var(--bg-elevated);
  color: var(--text-muted);
  cursor: pointer;
  font-size: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.theme-btn:hover {
  color: var(--accent);
  border-color: var(--accent);
}
</style>
