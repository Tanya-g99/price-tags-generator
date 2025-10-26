<template>
  <div class="product-list">
    <div v-for="p in products" :key="p.id" class="card">
      <h3>{{ p.name }}</h3>
      <p>{{ p.price }} ₽</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia';
import { useProductStore } from '@stores/productStore';

const store = useProductStore();
const { products } = storeToRefs(store);
</script>
