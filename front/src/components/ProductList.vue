<template>
  <div class="product-list">
    <Card v-for="p in products" :key="p.id" class="product-card">
      <template #content>
        <div class="product-card__content">
          <h3 class="product-card__title">{{ p.name }}</h3>
          <p class="product-card__price">{{ p.price }} ₽</p>
        </div>
      </template>
    </Card>
  </div>
</template>

<script lang="ts" setup>
import { useProductStore } from '@stores/productStore';
import { storeToRefs } from 'pinia';

const store = useProductStore();
const { products } = storeToRefs(store);
</script>

<style scoped lang="scss">
.product-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 1.5rem;
  padding: 1rem;

  @include mobile {
    grid-template-columns: 1fr;
    gap: 1rem;
    padding: 0.75rem;
  }
}
</style>
