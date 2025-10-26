import '@/assets/styles.scss';
import { router } from '@/router';
import { createPinia } from 'pinia';
import Toast from 'primevue/toast';
import { createApp } from 'vue';
import App from './App.vue';
import FormField from './components/FormField.vue';

// PrimeVue
import Button from 'primevue/button';
import Card from 'primevue/card';
import Checkbox from 'primevue/checkbox';
import ColorPicker from 'primevue/colorpicker';
import PrimeVue from 'primevue/config';
import Divider from 'primevue/divider';
import Dropdown from 'primevue/dropdown';
import InputNumber from 'primevue/inputnumber';
import InputText from 'primevue/inputtext';
import Panel from 'primevue/panel';
import Slider from 'primevue/slider';
import ToastService from 'primevue/toastservice';

// PrimeVue Theme
import { definePreset } from '@primevue/themes';
import Aura from '@primevue/themes/aura';
// PrimeIcons
import 'primeicons/primeicons.css';

const app = createApp(App);

app
  .use(createPinia())
  .use(router)
  .use(ToastService)
  .use(PrimeVue, {
    theme: {
      preset: definePreset(Aura, {
        semantic: {
          primary: {
            50: '#f0f4ff',
            100: '#e0e7ff',
            200: '#c7d2fe',
            300: '#a5b4fc',
            400: '#818cf8',
            500: '#1e3a8a',
            600: '#1a3479',
            700: '#162e68',
            800: '#122857',
            900: '#0e2246',
            950: '#0a1c35',
          },
          gray: {
            50: '#f9fafb', // Очень светлый серый
            100: '#f3f4f6', // Светлый серый
            200: '#e5e7eb', // Светло-серый
            300: '#d1d5db', // Серый
            400: '#9ca3af', // Средне-серый
            500: '#6b7280', // Серый
            600: '#4b5563', // Темно-серый
            700: '#374151', // Очень темно-серый
            800: '#1f2937', // Почти черный
            900: '#111827', // Черный
          },
        },
      }),
      options: {
        prefix: 'p',
        darkModeSelector: '.dark',
        cssLayer: false,
      },
    },
  });

app.component('Button', Button);
app.component('InputText', InputText);
app.component('InputNumber', InputNumber);
app.component('ColorPicker', ColorPicker);
app.component('Slider', Slider);
app.component('Checkbox', Checkbox);
app.component('Card', Card);
app.component('Panel', Panel);
app.component('Divider', Divider);
app.component('Dropdown', Dropdown);
app.component('FormField', FormField);
app.component('Toast', Toast);

app.mount('#app');
