interface IBaseElement {
  id: string;
  x: number;
  y: number;
  type: string;
}

export abstract class BaseElement implements IBaseElement {
  id: string;
  x: number;
  y: number;
  readonly type: string;

  protected constructor(type: string, defaults: { x?: number; y?: number }) {
    this.id = `${type}-${globalThis.Date.now()}`;
    this.type = type;
    this.x = defaults.x ?? 0;
    this.y = defaults.y ?? 0;
  }
}

export class TextElement extends BaseElement {
  readonly type = 'text' as const;
  content: string;
  fontSize: number;
  fontFamily: string;
  bold: boolean;
  italic: boolean;
  color: string;

  constructor(init?: Partial<Omit<TextElement, 'id' | 'type'>>) {
    super('text', { x: init?.x ?? 0, y: init?.y ?? 20 });
    this.content = init?.content ?? 'Название продукта';
    this.fontSize = init?.fontSize ?? 20;
    this.fontFamily = init?.fontFamily ?? 'Arial';
    this.bold = init?.bold ?? false;
    this.italic = init?.italic ?? false;
    this.color = init?.color ?? '#000000';
  }
}

export class PriceElement extends BaseElement {
  readonly type = 'price' as const;
  content: string;
  fontSize: number;
  fontFamily: string;
  bold: boolean;
  italic: boolean;
  color: string;

  constructor(init?: Partial<Omit<PriceElement, 'id' | 'type'>>) {
    super('price', { x: init?.x ?? 0, y: init?.y ?? 30 });
    this.content = init?.content ?? 'Цена ₽';
    this.fontSize = init?.fontSize ?? 30;
    this.fontFamily = init?.fontFamily ?? 'Arial';
    this.bold = init?.bold ?? true;
    this.italic = init?.italic ?? false;
    this.color = init?.color ?? '#000000';
  }
}

export class QrElement extends BaseElement {
  readonly type = 'qr' as const;
  size: number;
  href: string;

  constructor(init?: Partial<Omit<QrElement, 'id' | 'type'>>) {
    super('qr', { x: init?.x ?? 0, y: init?.y ?? 0 });
    this.size = init?.size ?? 80;
    this.href = init?.href ?? '/templates/assets/qr-code.svg';
  }
}

export class ImageElement extends BaseElement {
  readonly type = 'image' as const;
  width: number;
  height: number;
  href: string;
  preserveAspectRatio: string;

  constructor(init?: Partial<Omit<ImageElement, 'id' | 'type'>>) {
    super('image', { x: init?.x ?? 0, y: init?.y ?? 0 });
    this.width = init?.width ?? 60;
    this.height = init?.height ?? 40;
    this.href = init?.href ?? '/templates/assets/logo.svg';
    this.preserveAspectRatio = init?.preserveAspectRatio ?? 'xMaxYMid meet';
  }
}

export class RectElement extends BaseElement {
  readonly type = 'rect' as const;
  width: number;
  height: number;
  color: string;
  opacity: number;
  strokeColor: string;
  useMainColorAsStroke: boolean;

  constructor(init?: Partial<Omit<RectElement, 'id' | 'type'>>) {
    super('rect', { x: init?.x ?? 0, y: init?.y ?? 0 });
    this.width = init?.width ?? 80;
    this.height = init?.height ?? 40;
    this.color = init?.color ?? '#ff0000';
    this.opacity = init?.opacity ?? 1.0;
    this.strokeColor = init?.strokeColor ?? '#000000';
    this.useMainColorAsStroke = init?.useMainColorAsStroke ?? false;
  }
}

export class PriceTagTemplate {
  name?: string;
  backgroundColor: string;
  width: number;
  height: number;
  text?: TextElement;
  price?: PriceElement;
  qr?: QrElement;
  image?: ImageElement;
  rects: RectElement[];

  constructor(init?: Partial<PriceTagTemplate>) {
    this.backgroundColor = init?.backgroundColor ?? '#ffffff';
    this.width = init?.width ?? 300;
    this.height = init?.height ?? 150;
    this.rects = init?.rects ?? [];
  }
}

export function isTextOrPrice(
  el: BaseElement
): el is TextElement | PriceElement {
  return el.type === 'text' || el.type === 'price';
}

export function isText(el: BaseElement): el is TextElement {
  return el.type === 'text';
}

export function isPrice(el: BaseElement): el is PriceElement {
  return el.type === 'price';
}

export function isImageOrQr(el: BaseElement): el is ImageElement | QrElement {
  return el.type === 'image' || el.type === 'qr';
}

export function isQr(el: BaseElement): el is QrElement {
  return el.type === 'qr';
}

export function isImage(el: BaseElement): el is ImageElement {
  return el.type === 'image';
}

export function isRect(el: BaseElement): el is RectElement {
  return el.type === 'rect';
}
