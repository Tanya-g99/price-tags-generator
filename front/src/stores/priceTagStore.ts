import { defineStore } from 'pinia';
import type { IPriceTag } from '@/types/models';
import axios from 'axios';

export const usePriceTagStore = defineStore('priceTagStore', {
  state: () => ({
    priceTags: [] as IPriceTag[],
  }),
  actions: {
    async fetchPriceTags() {
      try {
        const res = await axios.get<IPriceTag[]>('/api/pricetags');
        this.priceTags = res.data || [];
        console.log('PriceTags', this.priceTags);
      } catch (err) {
        console.error('Ошибка при получении ценников:', err);
        this.priceTags = [];
      }
    },

    async addPriceTag(formData: FormData) {
      try {
        const res = await axios.post<IPriceTag>('/api/pricetags', formData);
        if (!this.priceTags) this.priceTags = [];
        this.priceTags.push(res.data);
      } catch (err) {
        console.error('Ошибка при создании ценника:', err);
        throw new Error('Ошибка при создании ценника');
      }
    },
  },
});
