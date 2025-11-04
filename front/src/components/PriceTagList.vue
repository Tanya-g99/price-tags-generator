<template>
  <div class="price-tag-list">
    <div v-for="tag in priceTags" :key="tag.id" class="card">
      <h3>{{ tag.productName }} {{ tag.price }} ₽</h3>

      <div v-if="tag.priceTagImage" class="pricetag-image">
        <img
          :src="`${API_URL}/${tag.priceTagImage}`"
          alt="Сформированный ценник"
        />
        <!-- {{ tag.priceTagImage }} -->
      </div>

      <!-- {{ tag }} -->
    </div>
  </div>
</template>

<script lang="ts" setup>
import { usePriceTagStore } from '@/stores/priceTagStore';
import { storeToRefs } from 'pinia';
import { onMounted } from 'vue';

const API_URL = import.meta.env.VITE_API_URL;

const store = usePriceTagStore();
const { priceTags } = storeToRefs(store);

onMounted(() => store.fetchPriceTags());
</script>

<style scoped>
.price-tag-list {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
}

.card {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 12px;
  width: 220px;
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.pricetag-image img {
  max-width: 200px;
  max-height: 150px;
  border: 1px solid #333;
  object-fit: contain;
}
</style>
