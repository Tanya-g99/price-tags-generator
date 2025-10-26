<template>
  <div class="card">
    <input v-model="name" placeholder="Название товара" />
    <input v-model.number="price" type="number" placeholder="Цена" />
    <button @click="addProduct">Добавить</button>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useProductStore } from '@stores/productStore';

const store = useProductStore();
const name = ref<string>('');
const price = ref<number | null>(null);

// Добавление товара
async function addProduct() {
  if (!name.value || price.value === null) return;

  await store.addProduct({
    name: name.value,
    price: price.value,
  });

  // Сброс формы
  name.value = '';
  price.value = null;
}
</script>

<style scoped>
.card {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 16px;
  max-width: 400px;
}

input {
  display: block;
  width: 100%;
  padding: 8px;
  margin-bottom: 12px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

button {
  background-color: #1e3a8a;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
}

button:hover {
  background-color: #163570;
}
</style>
