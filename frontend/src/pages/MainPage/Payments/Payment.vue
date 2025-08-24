<script setup lang="ts">
  import { dateFormat } from "@krainovsd/js-helpers";
  import { VText } from "@krainovsd/vue-ui";
  import type { IPayment } from "@/entities/payments";
  import { DATE_WITH_TIME_FORMAT } from "@/entities/tech";

  type Props = {
    payment: IPayment;
  };
  type Emits = {
    open: [id: number];
  };

  const props = defineProps<Props>();
  const emit = defineEmits<Emits>();
</script>

<template>
  <div
    :class="$style.base"
    role="button"
    tabindex="0"
    @click="emit('open', $props.payment.id)"
    @keydown="
      (event) => {
        if (event.key === 'Enter' || event.key === ' ') {
          emit('open', $props.payment.id);
        }
      }
    "
  >
    <div :class="$style.wrap">
      <VText size="lg">{{ $props.payment.name }}</VText>
      <div :class="$style.info">
        <VText :strong="true" size="lg" :class="$style.amount">
          {{ new Intl.NumberFormat("ru-RU").format($props.payment.amount) }} ₽
        </VText>
        <VText size="sm">{{ dateFormat($props.payment.date, DATE_WITH_TIME_FORMAT) }}</VText>
      </div>
    </div>
    <VText>{{ props.payment.description }}</VText>
  </div>
</template>

<style lang="scss" module>
  .base {
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-sm);
    background-color: var(--ksd-bg-fill-color);
    border-radius: var(--ksd-border-radius-sm);
    padding: var(--ksd-padding-sm) var(--ksd-padding);
    transition: all var(--ksd-transition-mid) ease;
    cursor: pointer;
    min-width: 300px;

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

  .wrap {
    display: flex;
    justify-content: space-between;
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
