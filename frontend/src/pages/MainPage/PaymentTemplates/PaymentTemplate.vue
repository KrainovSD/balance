<script setup lang="ts">
  import { VText } from "@krainovsd/vue-ui";
  import { computed } from "vue";
  import type { IPaymentTemplate } from "@/entities/payments";

  type Props = {
    template: IPaymentTemplate;
    used: number;
  };
  type Emits = {
    open: [id: number];
  };

  const props = defineProps<Props>();
  const emit = defineEmits<Emits>();
  const diff = computed(() => props.template.amount - props.used);
  const percent = computed(() => +Math.abs((diff.value * 100) / props.template.amount).toFixed(2));
</script>

<template>
  <div
    :class="$style.base"
    role="button"
    tabindex="0"
    @click="emit('open', $props.template.id)"
    @keydown="
      (event) => {
        if (event.key === 'Enter' || event.key === ' ') {
          emit('open', $props.template.id);
        }
      }
    "
  >
    <VText size="lg">{{ $props.template.name }}</VText>
    <div :class="$style.info">
      <VText :strong="true" size="lg" :class="$style.amount">
        {{ new Intl.NumberFormat("ru-RU").format($props.template.amount) }} ₽
      </VText>
      <VText :type="diff > 0 ? 'success' : 'error'" :class="$style.amount">
        {{ new Intl.NumberFormat("ru-RU").format(diff) }} ₽ ·
        {{ new Intl.NumberFormat("ru-RU").format(percent) }}%</VText
      >
    </div>
  </div>
</template>

<style lang="scss" module>
  .base {
    display: flex;
    background-color: var(--ksd-bg-fill-color);
    border-radius: var(--ksd-border-radius-sm);
    gap: var(--ksd-padding-sm);
    justify-content: space-between;
    padding: var(--ksd-padding-sm) var(--ksd-padding);
    transition: all var(--ksd-transition-mid) ease;
    min-width: 300px;
    cursor: pointer;

    &:focus-visible {
      outline: var(--ksd-outline-width) var(--ksd-outline-type) var(--ksd-outline-color);
      outline-offset: 1px;
      transition:
        outline-offset 0s,
        outline 0s;
    }

    &:hover {
      background-color: var(--ksd-bg-fill-hover-color);
    }

    &:active {
      background-color: var(--ksd-bg-active-color);
    }
  }

  .info {
    align-items: end;
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-xxs);
  }

  .amount {
    width: fit-content;
  }
</style>
