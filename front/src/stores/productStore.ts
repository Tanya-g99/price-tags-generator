import { defineStore } from 'pinia';
import type { IProduct } from '@/types/models';
import axios from 'axios';

export const useProductStore = defineStore('productStore', {
  state: () => ({
    products: [] as IProduct[], // всегда массив
  }),
  actions: {
    async fetchProducts() {
      const res = await axios.get<IProduct[]>('/api/products');
      this.products = res.data || [];
    },
    async addProduct(product: Omit<IProduct, 'id'>) {
      const res = await axios.post<IProduct>('/api/products', product);
      if (!this.products) this.products = []; // на всякий случай
      this.products.push(res.data);
    },
    async generateTestProducts() {
      await this.fetchProducts();

      if (this.products.length != 0) {
        return;
      }
      const testProducts: ReadonlyArray<Omit<IProduct, 'id'>> = [
        {
          name: 'Продукт1',
          price: 120,
        },
        {
          name: 'Продукт2',
          price: 120000,
        },
        {
          name: 'Продукт3',
          price: 5,
        },
      ];

      testProducts.forEach(product => {
        this.addProduct(product);
      });
    },
  },
});
