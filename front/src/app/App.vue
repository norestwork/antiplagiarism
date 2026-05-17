<template>
  <AnalysePage />
  <WelcomeDialog v-model="showWelcome" />
</template>

<script setup>
import { ref } from 'vue'
import AnalysePage from '@/pages/analyse/ui/AnalysePage.vue'
import WelcomeDialog from '@/widgets/WelcomeDialog/WelcomeDialog.vue'

const showWelcome = ref(true)
</script>
