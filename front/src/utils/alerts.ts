// Простая система уведомлений без зависимости от useToast
class AlertService {
  private static toastInstance: any = null;

  static setToastInstance(toast: any) {
    this.toastInstance = toast;
  }

  private static getToast() {
    if (!this.toastInstance) {
      console.warn(
        'Toast instance not initialized. Please call setToastInstance first.'
      );
      return null;
    }
    return this.toastInstance;
  }

  static success(message: string, summary: string = 'Успех') {
    const toast = this.getToast();
    if (toast) {
      toast.add({
        severity: 'success',
        summary,
        detail: message,
        life: 3000,
      });
    } else {
      console.log(`✅ ${summary}: ${message}`);
    }
  }

  static error(message: string, summary: string = 'Ошибка') {
    const toast = this.getToast();
    if (toast) {
      toast.add({
        severity: 'error',
        summary,
        detail: message,
        life: 5000,
      });
    } else {
      console.error(`❌ ${summary}: ${message}`);
    }
  }

  static info(message: string, summary: string = 'Информация') {
    const toast = this.getToast();
    if (toast) {
      toast.add({
        severity: 'info',
        summary,
        detail: message,
        life: 3000,
      });
    } else {
      console.log(`ℹ️ ${summary}: ${message}`);
    }
  }

  static warning(message: string, summary: string = 'Предупреждение') {
    const toast = this.getToast();
    if (toast) {
      toast.add({
        severity: 'warn',
        summary,
        detail: message,
        life: 4000,
      });
    } else {
      console.warn(`⚠️ ${summary}: ${message}`);
    }
  }

  // Специфичные методы для приложения
  static templateSaved() {
    this.success('Шаблон успешно сохранен!');
  }

  static templateSaveError(error?: string) {
    this.error(
      error || 'Произошла ошибка при сохранении шаблона',
      'Ошибка сохранения'
    );
  }

  static templateNameRequired() {
    this.error('Название шаблона не может быть пустым', 'Обязательное поле');
  }

  static productSaved() {
    this.success('Товар успешно сохранен!');
  }

  static productSaveError(error?: string) {
    this.error(
      error || 'Произошла ошибка при сохранении товара',
      'Ошибка сохранения'
    );
  }

  static priceTagGenerated() {
    this.success('Ценник успешно сгенерирован!');
  }

  static priceTagGenerateError(error?: string) {
    this.error(
      error || 'Произошла ошибка при генерации ценника',
      'Ошибка генерации'
    );
  }
}

// Экспортируем экземпляр для удобства
export const alerts = AlertService;
