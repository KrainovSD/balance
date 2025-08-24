<script setup lang="ts">
  import { VButton } from "@krainovsd/vue-ui";
  import { computed } from "vue";
  import { type IDate, getMonthName, useDateStore } from "@/entities/date";
  import { usePaymentsStore } from "@/entities/payments";
  import { useReceiptsStore } from "@/entities/receipts";

  const paymentsStore = usePaymentsStore();
  const receiptsStore = useReceiptsStore();
  const dateStore = useDateStore();

  const dates = computed(() => {
    const dates = new Set<string>();

    paymentsStore.payments.forEach((payment) => {
      const date = new Date(payment.date);
      const year = date.getFullYear();
      const month = date.getMonth();

      dates.add(`${year}-${month}`);
    });
    receiptsStore.receipts.forEach((receipt) => {
      const date = new Date(receipt.date);
      const year = date.getFullYear();
      const month = date.getMonth();

      dates.add(`${year}-${month}`);
    });

    return Array.from(dates).map<IDate>((date) => {
      const [year, month] = date.split("-");

      return [+year, +month];
    });
  });
</script>

<template>
  <aside :class="$style.root">
    <VButton v-for="date in dates" :key="`${date[0]}-${date[1]}`">
      {{ date[0] }} {{ getMonthName(date[1], "short") }}
    </VButton>
  </aside>
</template>

<style lang="scss" module>
  .root {
    --width-max: 310px;
    --width-min: 53px;
    height: 100%;
    width: 310px;
    display: flex;
    flex-direction: column;

    &[data-open="true"] {
      animation: appear var(--ksd-transition-slow) ease both;
    }
    &[data-open="false"] {
      animation: disappear var(--ksd-transition-slow) ease both;
    }

    &.appear {
      width: var(--width-max);
    }
    &.disappear {
      width: var(--width-min);
    }
  }

  @keyframes appear {
    from {
      width: var(--width-min);
    }
    to {
      width: var(--width-max);
    }
  }
  @keyframes disappear {
    from {
      width: var(--width-max);
    }
    to {
      width: var(--width-min);
    }
  }
</style>
