import type { IPriceTag } from '@/types/models';

export class PriceTagService {
  private baseUrl = 'http://localhost:8080/api/pricetags';

  async generateBarcode(data: {
    name: string;
    price: string;
    logoFile?: File;
  }): Promise<Blob> {
    const formData = new FormData();
    formData.append('name', data.name);
    formData.append('price', data.price);
    if (data.logoFile) formData.append('logo', data.logoFile);

    const res = await fetch(`${this.baseUrl}/barcode`, {
      method: 'POST',
      body: formData,
    });
    if (!res.ok) throw new Error('Ошибка генерации штрихкода');
    return res.blob();
  }

  async uploadPriceTag(
    dataUrl: string
  ): Promise<{ status: string; file: string }> {
    const res = await fetch(`${this.baseUrl}/upload`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ image: dataUrl }),
    });
    if (!res.ok) throw new Error('Ошибка сохранения ценника');
    return res.json();
  }

  async listPriceTags(): Promise<IPriceTag[]> {
    const res = await fetch(this.baseUrl);
    if (!res.ok) throw new Error('Ошибка получения ценников');
    return res.json();
  }
}

export const priceTagService = new PriceTagService();
