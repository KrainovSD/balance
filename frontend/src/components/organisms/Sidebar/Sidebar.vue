<script setup lang="ts">
  import { VButton } from "@krainovsd/vue-ui";
  import { computed, watch } from "vue";
  import { type IDate, getMonthName, useDateStore } from "@/entities/date";
  import { usePaymentsStore } from "@/entities/payments";
  import { useReceiptsStore } from "@/entities/receipts";
  import { isHasFeature } from "@/lib/check-feature";

  type Props = {
    open: boolean;
  };

  type Emits = {
    close: [];
  };

  defineProps<Props>();
  defineEmits<Emits>();

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

    return Array.from(dates)
      .map<IDate>((date) => {
        const [year, month] = date.split("-");

        return [+year, +month];
      })
      .sort((dateA, dateB) => {
        if (dateA[0] !== dateB[0]) return dateB[0] - dateA[0];

        return dateB[1] - dateA[1];
      });
  });

  watch(
    dates,
    (dates) => {
      if (dates[0] != undefined) dateStore.$patch({ date: dates[0] });
    },
    { immediate: true },
  );
</script>

<template>
  <aside :class="[$style.root]" :data-open="$props.open">
    <div :class="[$style.list, isHasFeature('hidden-scroll-sidebar') && $style.list_scrollHidden]">
      <template v-for="date in dates" :key="`${date[0]}-${date[1]}`">
        <VButton
          :type="
            dateStore.date && dateStore.date[0] === date[0] && dateStore.date[1] === date[1]
              ? 'default'
              : 'text'
          "
          @click="dateStore.$patch({ date })"
        >
          {{ date[0] }} {{ getMonthName(date[1], "short") }}
        </VButton>
      </template>
    </div>
  </aside>
  <div v-if="$props.open" :class="$style.overlay" @click.stop="$emit('close')"></div>
</template>

<style lang="scss" module>
  .list {
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-sm);
    padding: var(--ksd-padding);
    overflow: auto;

    &_scrollHidden {
      &::-webkit-scrollbar {
        width: 0;
        height: 0;
        display: none;
      }

      scrollbar-width: none;
      -ms-overflow-style: none;
    }
  }

  .overlay {
    display: none;
    @media (width < 700px) {
      display: block;
      z-index: var(--ksd-modal-z-index);
      position: fixed;
      inset: 0;
      height: 100%;
      background-color: var(--ksd-bg-mask-color);
      font-family: var(--ksd-font-family);
      font-size: var(--ksd-font-size);
      overflow: auto;
      animation: modal-mask-in var(--ksd-transition-mid) linear;
      &_out {
        animation: modal-mask-out var(--ksd-transition-mid) linear;
      }

      @keyframes ksd-modal-mask-in {
        from {
          opacity: 0;
        }
        to {
          opacity: 1;
        }
      }

      @keyframes ksd-modal-mask-out {
        from {
          opacity: 1;
        }
        to {
          opacity: 0;
        }
      }
    }
  }

  .root {
    height: 100%;
    width: 150px;
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-sm);
    overflow: hidden;
    background-color: var(--ksd-bg-color);
    z-index: 101;

    @media (width < 700px) {
      position: absolute;
      --left-min: -150px;
      --left-max: 0px;

      &[data-open="true"] {
        animation: appear var(--ksd-transition-slow) ease both;
      }
      &[data-open="false"] {
        animation: disappear var(--ksd-transition-slow) ease both;
      }

      &.appear {
        left: var(--left-max);
      }
      &.disappear {
        left: var(--left-min);
      }
    }

    @keyframes appear {
      from {
        left: var(--left-min);
      }
      to {
        left: var(--left-max);
      }
    }
    @keyframes disappear {
      from {
        left: var(--left-max);
      }
      to {
        left: var(--left-min);
      }
    }
  }
</style>
