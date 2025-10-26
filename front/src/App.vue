<template>
  <Navbar />
  <main>
    <router-view />
  </main>
  <Toast />
</template>

<script lang="ts" setup>
import Navbar from '@/components/Navbar.vue';
</script>

<style scoped>
main {
  max-width: 1200px;
  margin: 24px auto;
  padding: 0 16px;
  gap: 24px;
  display: flex;
  flex-direction: column;
}
</style>
