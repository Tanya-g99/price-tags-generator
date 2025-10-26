export interface IProduct {
  id: number;
  name: string;
  price: number;
}

export interface IPriceTag {
  id: number;
  productId: number;
  productName?: string; // название товара
  price?: number; // цена товара
  priceTagImage?: string; // сформированный PNG ценника
}
