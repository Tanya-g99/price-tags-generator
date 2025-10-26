import type { IProduct } from '@/types/models';

export class ProductService {
  private baseUrl = 'http://localhost:8080/api/products';

  async addProduct(product: {
    name: string;
    price: number;
    logoUrl: string;
  }): Promise<IProduct> {
    const res = await fetch(this.baseUrl, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(product),
    });
    if (!res.ok) throw new Error('Ошибка при добавлении товара');
    return res.json();
  }

  async listProducts(): Promise<IProduct[]> {
    const res = await fetch(this.baseUrl);
    if (!res.ok) throw new Error('Ошибка при получении товаров');
    return res.json();
  }
}

export const productService = new ProductService();
