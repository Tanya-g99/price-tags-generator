<template>
  <div class="price-tag-preview-container">
    <div class="controls">
      <button @click="zoomIn" class="control-btn" title="Увеличить">+</button>
      <button @click="zoomOut" class="control-btn" title="Уменьшить">-</button>
      <button @click="resetZoom" class="control-btn" title="Сбросить масштаб">
        ⌂
      </button>
      <span class="zoom-level">{{ Math.round(zoom * 100) }}%</span>
    </div>
    <div
      ref="svgContainer"
      class="price-tag-preview"
      @wheel="handleWheel"
      @mousedown="startPan"
      @mousemove="pan"
      @mouseup="endPan"
      @mouseleave="endPan"
    >
      <div
        ref="svgWrapper"
        class="svg-wrapper"
        :style="{
          transform: `scale(${zoom}) translate(${panX}px, ${panY}px)`,
          transformOrigin: 'center center',
        }"
      >
        <!-- SVG будет вставлен сюда -->
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTemplateStore } from '@/stores/templateStore';
import { onMounted, ref, watch } from 'vue';

interface Props {
  product: { name: string; price: number } | null;
  storeLogo: File | null;
  templateUrl?: string | null;
}

const props = defineProps<Props>();
const templatesStore = useTemplateStore();

const svgContainer = ref<HTMLElement | null>(null);
const svgWrapper = ref<HTMLElement | null>(null);

// Состояние масштабирования и панорамирования
const zoom = ref(1);
const panX = ref(0);
const panY = ref(0);
const isPanning = ref(false);
const lastPanX = ref(0);
const lastPanY = ref(0);

const getTemplate = async () => {
  if (!props.templateUrl) return;
  const res = await templatesStore.fetchTemplate(props.templateUrl);
  if (svgWrapper.value) {
    svgWrapper.value.innerHTML = res;
    updateSvg();
    resetZoom();
  }
};

// Загружаем SVG при монтировании
onMounted(async () => {
  getTemplate();
});

watch(() => props.product, updateSvg, { deep: true });
watch(() => props.storeLogo, updateSvg);
watch(() => props.templateUrl, getTemplate);

const fileToDataUrl = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = reject;
    reader.readAsDataURL(file);
  });
};

async function updateSvg() {
  if (!svgWrapper.value) return;
  const svgEl = svgWrapper.value.querySelector('svg');
  if (!svgEl) return;

  // Название товара
  const productName = svgEl.querySelector<SVGTextElement>('#productName');
  if (productName && props.product)
    productName.textContent = props.product.name;

  // Цена
  const price = svgEl.querySelector<SVGTextElement>('#price');
  if (price && props.product) price.textContent = props.product.price + ' ₽';

  // Логотип
  if (props.storeLogo) {
    const logo = svgEl.querySelector<SVGImageElement>('#storeLogo');
    if (logo) {
      logo.setAttribute('href', await fileToDataUrl(props.storeLogo));
    }
  }
}

// Функции масштабирования
function zoomIn() {
  zoom.value = Math.min(zoom.value * 1.2, 5);
}

function zoomOut() {
  zoom.value = Math.max(zoom.value / 1.2, 0.1);
}

function resetZoom() {
  zoom.value = 1;
  panX.value = 0;
  panY.value = 0;
}

// Обработка колесика мыши для масштабирования
function handleWheel(event: WheelEvent) {
  event.preventDefault();
  const delta = event.deltaY > 0 ? 0.9 : 1.1;
  zoom.value = Math.max(0.1, Math.min(5, zoom.value * delta));
}

// Функции панорамирования
function startPan(event: MouseEvent) {
  if (event.button === 0) {
    // Левая кнопка мыши
    isPanning.value = true;
    lastPanX.value = event.clientX;
    lastPanY.value = event.clientY;
    event.preventDefault();
  }
}

function pan(event: MouseEvent) {
  if (isPanning.value) {
    const deltaX = event.clientX - lastPanX.value;
    const deltaY = event.clientY - lastPanY.value;

    panX.value += deltaX;
    panY.value += deltaY;

    lastPanX.value = event.clientX;
    lastPanY.value = event.clientY;
  }
}

function endPan() {
  isPanning.value = false;
}
</script>

<style scoped>
.price-tag-preview-container {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #f9fafb;
  overflow: hidden;
}

.controls {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: #f3f4f6;
  border-bottom: 1px solid #e5e7eb;
}

.control-btn {
  width: 32px;
  height: 32px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  background: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: bold;
  transition: all 0.2s;
}

.control-btn:hover {
  background: #f3f4f6;
  border-color: #9ca3af;
}

.control-btn:active {
  background: #e5e7eb;
}

.zoom-level {
  font-size: 14px;
  color: #6b7280;
  font-weight: 500;
  margin-left: 8px;
}

.price-tag-preview {
  width: 100%;
  height: 200px;
  overflow: hidden;
  cursor: grab;
  position: relative;
}

.price-tag-preview:active {
  cursor: grabbing;
}

.svg-wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.1s ease-out;
}

.svg-wrapper svg {
  max-width: 100%;
  max-height: 100%;
  display: block;
}
</style>
