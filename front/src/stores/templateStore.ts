import axios from 'axios';
import { defineStore } from 'pinia';

export const useTemplateStore = defineStore('templateStore', {
  state: () => ({
    priceTagTemplates: [] as string[],
  }),
  actions: {
    async fetchPriceTagTemplates() {
      try {
        const res = await axios.get<string[]>('/api/templates-pricetags');
        console.log(res);
        this.priceTagTemplates = res.data || [];
      } catch (err) {
        console.error('Ошибка при получении шаблонов ценников', err);
        this.priceTagTemplates = [];
      }
    },
    async fetchTemplate(templatePath: string) {
      try {
        if (!templatePath) {
          return;
        }
        const res = await axios.get(`/templates${templatePath}`, {});

        console.log(`/templates${templatePath}`, res);
        return res.data;
      } catch (err) {
        console.error('Ошибка при получении шаблона', err);
      }
    },
    async addPriceTagTemplate(file: File) {
      try {
        const formData = new FormData();
        formData.append('file', file);

        const res = await axios.post('/api/templates-pricetags', formData, {
          headers: { 'Content-Type': 'multipart/form-data' },
        });

        console.log('Шаблон загружен:', res.data);

        // После успешной загрузки обновим список шаблонов
        await this.fetchPriceTagTemplates();

        return res.data;
      } catch (err) {
        console.error('Ошибка при загрузке шаблона', err);
        throw err;
      }
    },
  },
});
