<template>
  <Navbar />
  <main>
    <router-view />
  </main>
  <Toast />
</template>

<script lang="ts" setup>
import Navbar from '@/components/Navbar.vue';
import { usePriceTagStore } from '@/stores/priceTagStore';
import { useProductStore } from '@/stores/productStore';
import { useTemplateStore } from '@/stores/templateStore';
import { onMounted } from 'vue';

const productStore = useProductStore();
const priceTagStore = usePriceTagStore();
const templateStore = useTemplateStore();

onMounted(async () => {
  await productStore.fetchProducts();
  await priceTagStore.fetchPriceTags();
  await templateStore.fetchPriceTagTemplates();
});
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
