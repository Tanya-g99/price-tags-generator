<script lang="ts" setup>
import { useTemplateStore } from '@/stores/templateStore';
import {
  BaseElement,
  ImageElement,
  PriceElement,
  PriceTagTemplate,
  QrElement,
  RectElement,
  TextElement,
  isImage,
  isImageOrQr,
  isPrice,
  isQr,
  isRect,
  isText,
  isTextOrPrice,
} from '@/types/templateModels';
import { alerts } from '@/utils/alerts';
import { useToast } from 'primevue/usetoast';
import { computed, onMounted, ref } from 'vue';

const template = ref<PriceTagTemplate>(new PriceTagTemplate());
const draggingElement = ref<BaseElement | null>(null);
const selectedElement = ref<BaseElement | null>(null);
const selectedIndex = ref<number | null>(null);
const offset = ref({ x: 0, y: 0 });

const templateStore = useTemplateStore();
const toast = useToast();

// Инициализируем alerts с toast instance
onMounted(() => {
  alerts.setToastInstance(toast);
});

// Опции для preserveAspectRatio
const preserveAspectRatioOptions = [
  { label: 'Растянуть по размеру', value: 'none' },
  { label: 'Сохранить пропорции (вписать)', value: 'xMidYMid meet' },
  { label: 'Сохранить пропорции (обрезать)', value: 'xMidYMid slice' },
  { label: 'Выровнять по левому краю', value: 'xMinYMid meet' },
  { label: 'Выровнять по правому краю', value: 'xMaxYMid meet' },
  { label: 'Выровнять по верху', value: 'xMidYMin meet' },
  { label: 'Выровнять по низу', value: 'xMidYMax meet' },
];

function normalizeColor(color: string | undefined): string {
  if (!color) return '#000000';

  if (color.startsWith('#')) {
    return color;
  }

  return `#${color}`;
}

function addText() {
  template.value.text = new TextElement();
  selectedElement.value = template.value.text;
}

function addPrice() {
  template.value.price = new PriceElement();
  selectedElement.value = template.value.price;
}

function addQRCode() {
  template.value.qr = new QrElement();
  selectedElement.value = template.value.qr;
}

function addImage() {
  template.value.image = new ImageElement();
  selectedElement.value = template.value.image;
}

function addRect() {
  const newRect = new RectElement();
  template.value.rects.push(newRect);
  selectedElement.value = newRect;
}

function startDrag(element: BaseElement, event: MouseEvent, index: number) {
  draggingElement.value = element;
  selectedElement.value = element;
  selectedIndex.value = index;
  offset.value = {
    x: event.clientX - element.x,
    y: event.clientY - element.y,
  };
}

function onDrag(event: MouseEvent) {
  if (draggingElement.value) {
    const elementWidth =
      'width' in draggingElement.value
        ? (draggingElement.value as any).width
        : 0;
    const elementHeight =
      'height' in draggingElement.value
        ? (draggingElement.value as any).height
        : 0;
    draggingElement.value.x = Math.max(
      0,
      Math.min(
        template.value.width - elementWidth,
        event.clientX - offset.value.x
      )
    );
    draggingElement.value.y = Math.max(
      0,
      Math.min(
        template.value.height - elementHeight,
        event.clientY - offset.value.y
      )
    );
  }
}

function endDrag() {
  draggingElement.value = null;
}

function deleteElement() {
  if (selectedElement.value) {
    if (isText(selectedElement.value)) {
      template.value.text = undefined;
    } else if (isPrice(selectedElement.value)) {
      template.value.price = undefined;
    } else if (isImageOrQr(selectedElement.value)) {
      if (isQr(selectedElement.value)) {
        template.value.qr = undefined;
      } else if (isImage(selectedElement.value)) {
        template.value.image = undefined;
      }
    } else if (isRect(selectedElement.value)) {
      const index = template.value.rects.findIndex(
        rect => rect.id === selectedElement.value!.id
      );
      if (index !== -1) {
        template.value.rects.splice(index, 1);
      }
    }

    selectedElement.value = null;
    selectedIndex.value = null;
  }
}

function bringToFront() {
  if (selectedElement.value && isRect(selectedElement.value)) {
    const rects = template.value.rects;
    const idx = rects.findIndex(r => r.id === selectedElement.value!.id);
    if (idx !== -1 && idx < rects.length - 1) {
      const temp = rects[idx + 1];
      rects[idx + 1] = rects[idx];
      rects[idx] = temp;
    }
  }
}

function sendToBack() {
  if (selectedElement.value && isRect(selectedElement.value)) {
    const rects = template.value.rects;
    const idx = rects.findIndex(r => r.id === selectedElement.value!.id);
    if (idx !== -1 && idx > 0) {
      const temp = rects[idx - 1];
      rects[idx - 1] = rects[idx];
      rects[idx] = temp;
    }
  }
}

const elements = computed(() => {
  const els: BaseElement[] = [];
  if (template.value.text) els.push(template.value.text);
  if (template.value.price) els.push(template.value.price);
  if (template.value.qr) els.push(template.value.qr);
  if (template.value.image) els.push(template.value.image);
  els.push(...template.value.rects);
  return els;
});

// Масштабирование для превью
const previewScale = computed(() => {
  const maxWidth = 400; // Максимальная ширина превью
  const maxHeight = 300; // Максимальная высота превью

  const scaleX = maxWidth / template.value.width;
  const scaleY = maxHeight / template.value.height;

  // Используем минимальный масштаб, чтобы поместить весь ценник
  return Math.min(scaleX, scaleY, 1); // Не увеличиваем больше 100%
});

const previewWidth = computed(() => template.value.width * previewScale.value);
const previewHeight = computed(
  () => template.value.height * previewScale.value
);

// Список доступных шрифтов
const fontFamilies = [
  { label: 'Arial', value: 'Arial' },
  { label: 'Times New Roman', value: 'Times New Roman' },
  { label: 'Helvetica', value: 'Helvetica' },
  { label: 'Georgia', value: 'Georgia' },
  { label: 'Verdana', value: 'Verdana' },
  { label: 'Courier New', value: 'Courier New' },
  { label: 'Impact', value: 'Impact' },
  { label: 'Comic Sans MS', value: 'Comic Sans MS' },
  { label: 'Trebuchet MS', value: 'Trebuchet MS' },
  { label: 'Palatino', value: 'Palatino' },
];

async function addPriceTagTemplate() {
  if (!template.value.name) {
    alerts.templateNameRequired();
    return;
  }
  const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
  svg.setAttribute('width', template.value.width.toString());
  svg.setAttribute('height', template.value.height.toString());
  svg.setAttribute(
    'viewBox',
    `0 0 ${template.value.width} ${template.value.height}`
  );

  // Фон
  const background = document.createElementNS(
    'http://www.w3.org/2000/svg',
    'rect'
  );
  background.setAttribute('width', template.value.width.toString());
  background.setAttribute('height', template.value.height.toString());
  background.setAttribute(
    'fill',
    normalizeColor(template.value.backgroundColor)
  );
  background.setAttribute('rx', '12');
  background.setAttribute('ry', '12');
  svg.appendChild(background);

  // Сначала добавляем прямоугольники (области) - они будут фоном
  elements.value.forEach(el => {
    if (isRect(el)) {
      const rect = document.createElementNS(
        'http://www.w3.org/2000/svg',
        'rect'
      );
      rect.setAttribute('x', el.x.toString());
      rect.setAttribute('y', el.y.toString());
      rect.setAttribute('width', el.width.toString());
      rect.setAttribute('height', el.height.toString());
      rect.setAttribute('fill', normalizeColor(el.color));
      rect.setAttribute('opacity', (el.opacity || 1.0).toString());
      rect.setAttribute(
        'stroke',
        el.useMainColorAsStroke
          ? normalizeColor(el.color)
          : normalizeColor(el.strokeColor)
      );
      svg.appendChild(rect);
    }
  });

  // Затем добавляем остальные элементы (текст, изображения, QR) - они будут поверх областей
  elements.value.forEach(el => {
    if (isTextOrPrice(el)) {
      const text = document.createElementNS(
        'http://www.w3.org/2000/svg',
        'text'
      );
      text.setAttribute('x', el.x.toString());
      text.setAttribute('y', el.y.toString());
      text.setAttribute('font-family', el.fontFamily || 'Arial');
      text.setAttribute('font-size', (el.fontSize || 16).toString());
      text.setAttribute('font-weight', el.bold ? 'bold' : 'normal');
      text.setAttribute('font-style', el.italic ? 'italic' : 'normal');
      text.setAttribute('fill', normalizeColor(el.color));
      if (isPrice(el)) {
        text.setAttribute('id', 'price');
        text.textContent = 'Цена ₽';
      } else if (isText(el)) {
        text.setAttribute('id', 'productName');
        text.textContent = 'Название продукта';
      }
      svg.appendChild(text);
    } else if (isQr(el)) {
      const rect = document.createElementNS(
        'http://www.w3.org/2000/svg',
        'image'
      );
      rect.setAttribute('x', el.x.toString());
      rect.setAttribute('y', el.y.toString());
      rect.setAttribute('width', (el.size || 50).toString());
      rect.setAttribute('height', (el.size || 50).toString());
      rect.setAttribute('href', el.href);
      rect.setAttribute('preserveAspectRatio', 'xMaxYMid meet');
      svg.appendChild(rect);
    } else if (isImage(el)) {
      const image = document.createElementNS(
        'http://www.w3.org/2000/svg',
        'image'
      );
      image.setAttribute('id', 'storeLogo');
      image.setAttribute('x', el.x.toString());
      image.setAttribute('y', el.y.toString());
      image.setAttribute('width', (el.width || 50).toString());
      image.setAttribute('height', (el.height || 50).toString());
      image.setAttribute('preserveAspectRatio', el.preserveAspectRatio);
      image.setAttribute('href', el.href);
      svg.appendChild(image);
    }
  });

  const svgString = new XMLSerializer().serializeToString(svg);

  try {
    // Создаем File из SVG строки
    const blob = new Blob([svgString], { type: 'image/svg+xml' });
    const file = new File([blob], `${template.value.name || 'price-tag'}.svg`, {
      type: 'image/svg+xml',
    });

    // Используем store для сохранения
    const result = await templateStore.addPriceTagTemplate(file);
    console.log('Шаблон сохранен:', result);
    alerts.templateSaved();
  } catch (error) {
    console.error('Ошибка сохранения шаблона:', error);
    alerts.templateSaveError(
      error instanceof Error ? error.message : undefined
    );
  }
}
</script>

<template>
  <div class="template-form">
    <h1 class="template-form__title">Конструктор ценника</h1>

    <Card>
      <template #content>
        <div class="template-form__main-layout">
          <div class="template-form__left-column">
            <div class="template-form__section">
              <h3 class="template-form__section-title">
                <i class="pi pi-cog template-form__icon"></i>
                Настройки шаблона
              </h3>
              <div class="template-form__settings-grid">
                <FormField label="Название шаблона" field-id="template-name">
                  <InputText v-model="template.name" id="template-name" />
                </FormField>
                <FormField label="Цвет фона" field-id="template-bg">
                  <ColorPicker
                    v-model="template.backgroundColor"
                    id="template-bg"
                  />
                </FormField>
                <FormField label="Ширина" field-id="template-width">
                  <InputNumber
                    v-model="template.width"
                    :min="100"
                    id="template-width"
                  />
                </FormField>
                <FormField label="Высота" field-id="template-height">
                  <InputNumber
                    v-model="template.height"
                    :min="50"
                    id="template-height"
                  />
                </FormField>
              </div>
            </div>

            <div class="template-form__section">
              <h3 class="template-form__section-title">
                <i class="pi pi-plus-circle template-form__icon"></i>
                Добавить элементы
              </h3>
              <div class="template-form__button-group">
                <Button
                  @click="addText"
                  icon="pi pi-font"
                  label="Добавить текст"
                  severity="info"
                />
                <Button
                  @click="addPrice"
                  icon="pi pi-dollar"
                  label="Добавить цену"
                  severity="info"
                />
                <Button
                  @click="addQRCode"
                  icon="pi pi-qrcode"
                  label="Добавить QR"
                  severity="info"
                />
                <Button
                  @click="addImage"
                  icon="pi pi-image"
                  label="Добавить картинку"
                  severity="info"
                />
                <Button
                  @click="addRect"
                  icon="pi pi-square"
                  label="Добавить область"
                  severity="info"
                />
              </div>
            </div>

            <div class="template-form__section">
              <h3 class="template-form__section-title">
                <i class="pi pi-eye template-form__icon"></i>
                Превью шаблона
                <span v-if="previewScale < 1" class="template-form__scale-info">
                  ({{ Math.round(previewScale * 100) }}% Масштаб предпросмотра)
                </span>
              </h3>
              <!-- {{ elements }} -->
              <div class="template-form__preview-container">
                <svg
                  :width="previewWidth"
                  :height="previewHeight"
                  :viewBox="`0 0 ${template.width} ${template.height}`"
                  :style="{
                    backgroundColor: normalizeColor(template.backgroundColor),
                  }"
                  @mousemove="onDrag"
                  @mouseup="endDrag"
                >
                  <template
                    v-for="(el, index) in elements"
                    :key="`element-${index}`"
                  >
                    <template
                      v-for="(el, index) in elements.filter(e => isRect(e))"
                      :key="`rect-${el.id || index}`"
                    >
                      <rect
                        :x="el.x"
                        :y="el.y"
                        :width="el.width"
                        :height="el.height"
                        :fill="normalizeColor(el.color)"
                        :opacity="el.opacity || 1.0"
                        :stroke="
                          el.useMainColorAsStroke
                            ? normalizeColor(el.color)
                            : normalizeColor(el.strokeColor)
                        "
                        @mousedown="startDrag(el, $event, elements.indexOf(el))"
                        style="cursor: pointer"
                      />
                    </template>
                    <template
                      v-for="(el, index) in elements.filter(e => !isRect(e))"
                      :key="`element-${el.id || index}`"
                    >
                      <text
                        v-if="isTextOrPrice(el)"
                        :x="el.x"
                        :y="el.y"
                        :fill="normalizeColor(el.color)"
                        :font-size="el.fontSize || 16"
                        :font-family="el.fontFamily || 'Arial'"
                        :font-weight="el.bold ? 'bold' : 'normal'"
                        :font-style="el.italic ? 'italic' : 'normal'"
                        @mousedown="startDrag(el, $event, elements.indexOf(el))"
                        style="cursor: pointer"
                      >
                        {{ el.content }}
                      </text>
                      <image
                        v-else-if="isQr(el)"
                        :x="el.x"
                        :y="el.y"
                        :width="el.size || 50"
                        :height="el.size || 50"
                        :href="el.href"
                        preserveAspectRatio="xMaxYMid meet"
                        @mousedown="startDrag(el, $event, elements.indexOf(el))"
                        style="cursor: pointer"
                      />
                      <image
                        v-else-if="isImage(el)"
                        :x="el.x"
                        :y="el.y"
                        :width="el.width || 50"
                        :height="el.height || 50"
                        :href="el.href"
                        :preserveAspectRatio="el.preserveAspectRatio"
                        @mousedown="startDrag(el, $event, elements.indexOf(el))"
                        style="cursor: pointer"
                      />
                    </template>
                  </template>
                </svg>
              </div>
            </div>
          </div>

          <div class="template-form__right-column">
            <div class="template-form__section">
              <h3 class="template-form__section-title">
                <i class="pi pi-pencil template-form__icon"></i>
                Редактирование элемента
              </h3>

              <div v-if="!selectedElement" class="template-form__no-selection">
                <p>Выберите элемент для редактирования</p>
              </div>
              <div v-else class="template-form__element-editor">
                <div class="template-form__settings-grid">
                  <!-- Общие поля для всех элементов -->
                  <FormField label="Позиция X" field-id="element-x">
                    <InputNumber v-model="selectedElement!.x" id="element-x" />
                  </FormField>
                  <FormField label="Позиция Y" field-id="element-y">
                    <InputNumber v-model="selectedElement!.y" id="element-y" />
                  </FormField>

                  <!-- Поля для текста и цены -->
                  <template v-if="isTextOrPrice(selectedElement)">
                    <FormField label="Текст" field-id="element-text">
                      <InputText
                        v-model="
                          (selectedElement as TextElement | PriceElement)
                            .content
                        "
                        id="element-text"
                      />
                    </FormField>
                    <FormField label="Цвет" field-id="element-color">
                      <ColorPicker
                        v-model="
                          (selectedElement as TextElement | PriceElement).color
                        "
                        id="element-color"
                      />
                    </FormField>
                    <FormField
                      label="Размер шрифта"
                      field-id="element-font-size"
                    >
                      <InputNumber
                        v-model="
                          (selectedElement as TextElement | PriceElement)
                            .fontSize
                        "
                        :min="8"
                        :max="72"
                        id="element-font-size"
                      />
                    </FormField>
                    <FormField label="Шрифт" field-id="element-font-family">
                      <Dropdown
                        v-model="
                          (selectedElement as TextElement | PriceElement)
                            .fontFamily
                        "
                        :options="fontFamilies"
                        optionLabel="label"
                        optionValue="value"
                        placeholder="Выберите шрифт"
                        id="element-font-family"
                      />
                    </FormField>
                    <FormField label="Жирный" field-id="bold">
                      <Checkbox
                        :model-value="
                          (selectedElement as TextElement | PriceElement).bold
                            ? true
                            : false
                        "
                        @update:model-value="
                          val => {
                            if (selectedElement)
                              (
                                selectedElement as TextElement | PriceElement
                              ).bold = !!val;
                          }
                        "
                        :binary="true"
                        inputId="bold"
                      />
                    </FormField>
                    <FormField label="Курсив" field-id="italic">
                      <Checkbox
                        :model-value="
                          (selectedElement as TextElement | PriceElement).italic
                            ? true
                            : false
                        "
                        @update:model-value="
                          val => {
                            if (selectedElement)
                              (
                                selectedElement as TextElement | PriceElement
                              ).italic = !!val;
                          }
                        "
                        :binary="true"
                        inputId="italic"
                      />
                    </FormField>
                  </template>
                </div>

                <!-- Поля для изображений и QR -->
                <template v-if="isQr(selectedElement)">
                  <FormField label="Размер" field-id="element-size">
                    <InputNumber
                      v-model="(selectedElement as QrElement).size"
                      :min="20"
                      :max="template.width"
                      id="element-size"
                    />
                  </FormField>
                </template>
                <template v-if="isImage(selectedElement)">
                  <FormField label="Ширина" field-id="element-width">
                    <InputNumber
                      v-model="(selectedElement as ImageElement).width"
                      :min="20"
                      :max="template.width"
                      id="element-width"
                    />
                  </FormField>
                  <FormField label="Высота" field-id="element-height">
                    <InputNumber
                      v-model="(selectedElement as ImageElement).height"
                      :min="20"
                      :max="template.height"
                      id="element-height"
                    />
                  </FormField>
                </template>
                <template v-if="isImageOrQr(selectedElement)">
                  <FormField
                    v-if="selectedElement && isImage(selectedElement)"
                    label="Режим масштабирования"
                    field-id="element-preserve-aspect-ratio"
                  >
                    <Dropdown
                      v-model="
                        (selectedElement as ImageElement).preserveAspectRatio
                      "
                      :options="preserveAspectRatioOptions"
                      option-label="label"
                      option-value="value"
                      id="element-preserve-aspect-ratio"
                    />
                  </FormField>
                </template>

                <!-- Поля для прямоугольников -->
                <template v-if="selectedElement && isRect(selectedElement)">
                  <FormField label="Ширина" field-id="rect-width">
                    <InputNumber
                      v-model="(selectedElement as RectElement).width"
                      :min="1"
                      :max="template.width"
                      id="rect-width"
                    />
                  </FormField>
                  <FormField label="Высота" field-id="rect-height">
                    <InputNumber
                      v-model="(selectedElement as RectElement).height"
                      :min="1"
                      :max="template.height"
                      id="rect-height"
                    />
                  </FormField>
                  <FormField label="Цвет области">
                    <ColorPicker
                      v-model="(selectedElement as RectElement).color"
                    />
                  </FormField>
                  <FormField label="Прозрачность" field-id="rect-opacity">
                    <Slider
                      v-model="(selectedElement as RectElement).opacity"
                      :min="0"
                      :max="1"
                      :step="0.1"
                      id="rect-opacity"
                    />
                  </FormField>
                  <FormField
                    label="Использовать основной цвет как обводку"
                    field-id="rect-use-main-color"
                  >
                    <Checkbox
                      :model-value="
                        (selectedElement as RectElement).useMainColorAsStroke
                          ? true
                          : false
                      "
                      @update:model-value="
                        val => {
                          if (selectedElement)
                            (
                              selectedElement as RectElement
                            ).useMainColorAsStroke = !!val;
                        }
                      "
                      :binary="true"
                      inputId="rect-use-main-color"
                    />
                  </FormField>
                  <FormField
                    v-if="
                      selectedElement &&
                      !(selectedElement as RectElement).useMainColorAsStroke
                    "
                    label="Цвет обводки"
                  >
                    <ColorPicker
                      v-model="(selectedElement as RectElement).strokeColor"
                    />
                  </FormField>
                </template>
              </div>

              <Divider />

              <div
                v-if="selectedElement"
                class="template-form__actions-container"
              >
                <Button
                  @click="deleteElement"
                  icon="pi pi-trash"
                  label="Удалить"
                  severity="danger"
                />
                <Button
                  @click="bringToFront"
                  icon="pi pi-arrow-up"
                  label="На передний план"
                  severity="secondary"
                />
                <Button
                  @click="sendToBack"
                  icon="pi pi-arrow-down"
                  label="На задний план"
                  severity="secondary"
                />
              </div>
            </div>
          </div>
        </div>

        <Divider />
        <div class="template-form__save-container">
          <Button
            @click="addPriceTagTemplate"
            icon="pi pi-save"
            label="Сохранить шаблон"
            severity="success"
          />
        </div>
      </template>
    </Card>
  </div>
</template>

<style scoped lang="scss">
.template-form {
  padding: 1rem;
  max-width: 100%;

  @include mobile {
    padding: 0.75rem;
  }

  @include desktop {
    padding: 1.5rem;
  }

  &__title {
    font-size: 1.75rem;
    font-weight: bold;
    color: #1f2937;
    margin-bottom: 1.5rem;
    text-align: center;

    @include mobile {
      font-size: 1.5rem;
      margin-bottom: 1rem;
    }

    @include desktop {
      font-size: 2rem;
      margin-bottom: 2rem;
      text-align: left;
    }
  }

  &__main-layout {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;

    @include desktop {
      flex-direction: row;
      gap: 2rem;
      align-items: flex-start;
    }
  }

  &__left-column {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    flex: 1;

    @include desktop {
      gap: 2rem;
    }
  }

  &__right-column {
    @include desktop {
      position: sticky;
      top: 1rem;
      max-height: calc(100vh - 2rem);
      overflow-y: auto;
      width: 400px;
      flex-shrink: 0;
    }
  }

  &__section {
    background: #f9fafb;
    border-radius: 0.75rem;
    padding: 1.25rem;
    border: 1px solid #e5e7eb;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.2s ease;

    &:hover {
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    }

    @include mobile {
      padding: 1rem;
      border-radius: 0.5rem;
    }

    @include desktop {
      padding: 1.5rem;
    }
  }

  &__section-title {
    display: flex;
    align-items: center;
    font-size: 1.125rem;
    font-weight: 600;
    color: #374151;
    margin-bottom: 1rem;

    @include mobile {
      font-size: 1rem;
      margin-bottom: 0.75rem;
    }
  }

  &__scale-info {
    font-size: 0.875rem;
    font-weight: 400;
    color: #6b7280;
    margin-left: 0.5rem;

    @include mobile {
      font-size: 0.75rem;
    }
  }

  &__icon {
    margin-right: 0.5rem;
  }

  &__button-group {
    display: flex;
    flex-wrap: wrap;
    gap: 0.75rem;

    @include mobile {
      gap: 0.5rem;
    }
  }

  &__settings-grid {
    display: grid;
    gap: 1rem;

    @include mobile {
      grid-template-columns: 1fr;
      gap: 0.75rem;
    }

    @include tablet {
      grid-template-columns: repeat(2, 1fr);
    }

    @include desktop {
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    }
  }

  &__element-editor {
    .template-form__settings-grid {
      margin-bottom: 1rem;
    }
  }

  &__field {
    margin-bottom: 1rem;

    &--checkbox {
      display: flex;
      align-items: center;

      .template-form__checkbox-label {
        margin-left: 0.5rem;
        font-size: 0.875rem;
        color: #374151;
      }
    }
  }

  &__preview-container {
    display: flex;
    justify-content: center;
    padding: 1rem;
    background: #ffffff;
    border-radius: 0.5rem;
    border: 2px dashed #d1d5db;

    svg {
      border: 1px solid #d1d5db;
      border-radius: 0.5rem;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
      max-width: 100%;
      height: auto;
    }

    @include mobile {
      padding: 0.75rem;
    }
  }

  &__actions-container {
    display: flex;
    flex-wrap: wrap;
    gap: 0.75rem;

    @include mobile {
      gap: 0.5rem;
      justify-content: center;
    }
  }

  &__save-container {
    display: flex;
    justify-content: center;
  }

  &__no-selection {
    padding: 2rem;
    text-align: center;
    color: var(--p-gray-500);
    font-style: italic;
  }
}
</style>
