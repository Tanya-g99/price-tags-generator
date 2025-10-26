<template>
  <div class="form-field">
    <label :for="fieldId" class="form-field__label">{{ label }}</label>
    <slot />
  </div>
</template>

<script lang="ts" setup>
interface Props {
  label: string;
  fieldId?: string;
}

const props = withDefaults(defineProps<Props>(), {
  fieldId: () => `field-${Math.random().toString(36).substr(2, 9)}`,
});
</script>

<style scoped lang="scss">
.form-field {
  margin-bottom: 1rem;
  display: flex;
  flex-direction: column;

  &__label {
    display: block;
    font-weight: 500;
    margin-bottom: 0.5rem;
    color: var(--p-gray-400);
  }
}
</style>
