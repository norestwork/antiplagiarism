<template>
  <header class="header">
    <div class="logo">Антиплагиат</div>
  </header>
</template>

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
}

.logo {
  font-family: var(--mono);
  font-size: 13px;
  font-weight: 500;
  color: var(--accent);
  letter-spacing: 0.08em;
  text-transform: uppercase;
}
</style>
