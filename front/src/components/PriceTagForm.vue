<template>
  <div class="card">
    <label>Товар</label>
    <select v-model="selectedProductId">
      <option disabled value="">-- Выберите товар --</option>
      <option v-for="p in products" :key="p.id" :value="p.id">
        {{ p.name }}
      </option>
    </select>

    <label>Шаблон ценника</label>
    <select v-model="selectedPriceTagTemplate">
      <option disabled value="">-- Выберите шаблон --</option>
      <option v-for="tpl in priceTagTemplates" :key="tpl" :value="tpl">
        {{ tpl }}
      </option>
    </select>

    <label>Логотип магазина</label>
    <input type="file" accept="image/*" @change="onFileChange" />

    <h3>Превью ценника:</h3>
    <div>
      <PriceTagPreview
        :product="selectedProduct"
        :storeLogo="storeLogo"
        :templateUrl="templateUrl"
      />
    </div>

    <button @click="addPriceTag">Создать ценник</button>
  </div>
</template>

<script lang="ts" setup>
import { usePriceTagStore } from '@/stores/priceTagStore';
import { useProductStore } from '@/stores/productStore';
import { useTemplateStore } from '@/stores/templateStore';
import { computed, onMounted, ref } from 'vue';
import PriceTagPreview from './PriceTagPreview.vue';

const productStore = useProductStore();
const priceTagStore = usePriceTagStore();
const templatesStore = useTemplateStore();

const selectedPriceTagTemplate = ref<string | null>(null);
const selectedProductId = ref<number | null>(null);
const storeLogo = ref<File | null>(null);

const products = computed(() => productStore.products);
const selectedProduct = computed(
  () => products.value.find(p => p.id === selectedProductId.value) || null
);

const priceTagTemplates = computed(() => templatesStore.priceTagTemplates);
const templateUrl = computed(() =>
  selectedPriceTagTemplate.value
    ? `/pricetags/${selectedPriceTagTemplate.value}`
    : null
);

const onFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement;
  if (target.files && target.files.length > 0)
    storeLogo.value = target.files[0];
};

const addPriceTag = async () => {
  if (!selectedProduct.value || !storeLogo.value) {
    alert('Выберите товар и логотип');
    return;
  }

  if (!selectedPriceTagTemplate.value) {
    alert('Шаюлон не выбран');
    return;
  }

  try {
    const formData = new FormData();
    formData.append('productId', selectedProduct.value.id.toString());
    formData.append('storeLogo', storeLogo.value);
    formData.append('template', selectedPriceTagTemplate.value);

    await priceTagStore.addPriceTag(formData);

    selectedProductId.value = null;
    storeLogo.value = null;
    alert('Логотип загружен!');
  } catch (err) {
    console.error('Ошибка при создании ценника:', err);
    alert('Ошибка при создании ценника');
  }
};

onMounted(async () => {
  await productStore.generateTestProducts();
  productStore.fetchProducts();
  await templatesStore.fetchPriceTagTemplates();
  selectedPriceTagTemplate.value = templatesStore.priceTagTemplates[0] || null;
  console.log(selectedPriceTagTemplate.value);
});
</script>

<style scoped>
.card {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 16px;
  max-width: 600px;
  background-color: #fff;
}

select,
input {
  display: block;
  width: 100%;
  padding: 8px;
  margin-bottom: 12px;
}

button {
  background-color: #1e3a8a;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 8px;
  cursor: pointer;
  margin-top: 12px;
}

button:hover {
  background-color: #163570;
}
</style>
