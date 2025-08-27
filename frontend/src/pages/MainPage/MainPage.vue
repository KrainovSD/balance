<script setup lang="ts">
  import { VFilterOutlined } from "@krainovsd/vue-icons";
  import { type SelectItem, VButton, VDivider, VSelect, VText } from "@krainovsd/vue-ui";
  import { computed, onMounted, ref } from "vue";
  import { useDateStore } from "@/entities/date";
  import { MODES, useModeStore } from "@/entities/mode";
  import { usePaymentsStore } from "@/entities/payments";
  import { useReceiptsStore } from "@/entities/receipts";
  import { isHasFeature } from "@/lib/check-feature";
  import PaymentTemplates from "./PaymentTemplates/PaymentTemplates.vue";
  import Payments from "./Payments/Payments.vue";

  const paymentsStore = usePaymentsStore();
  const receiptsStore = useReceiptsStore();
  const modeStore = useModeStore();
  const dateStore = useDateStore();

  const filter = ref<number[] | null>([]);

  const paymentTemplateVisible = computed(() => paymentsStore.paymentTemplates.length > 0);
  const paymentVisible = computed(() => paymentsStore.payments.length > 0);

  const filteredByDatePayments = computed(() => {
    if (modeStore.mode !== MODES.Payment) return [];

    return paymentsStore.payments.filter((payment) => {
      if (!dateStore.date) return false;
      const date = new Date(payment.date);

      return date.getFullYear() === dateStore.date[0] && date.getMonth() === dateStore.date[1];
    });
  });
  const filteredPayments = computed(() =>
    filteredByDatePayments.value.filter((payment) => {
      return !filter.value || filter.value.length === 0 || filter.value.includes(payment.paymentId);
    }),
  );
  const paymentsAmountMap = computed(() => {
    const map: Record<string, number> = {};
    filteredByDatePayments.value.forEach((payment) => {
      map[payment.paymentId] ??= 0;
      map[payment.paymentId] += payment.amount;
    });

    return map;
  });
  const paymentsTotalAmount = computed(() =>
    filteredByDatePayments.value.reduce((acc, payment) => {
      acc += payment.amount;

      return acc;
    }, 0),
  );

  const filterOptions = computed(() => {
    switch (modeStore.mode) {
      case MODES.Payment: {
        return paymentsStore.paymentTemplates.map<SelectItem>((template) => ({
          label: template.name,
          value: template.id,
        }));
      }
      case MODES.Receipt: {
        return receiptsStore.receiptTemplates.map<SelectItem>((template) => ({
          label: template.name,
          value: template.id,
        }));
      }
      default: {
        return [];
      }
    }
  });

  onMounted(() => {
    void paymentsStore.getPayments();
    void receiptsStore.getReceipts();
    void paymentsStore.getPaymentTemplates();
    void receiptsStore.getReceiptTemplates();
  });
</script>

<template>
  <div :class="$style.base">
    <div :class="$style.filter">
      <VFilterOutlined />
      <VSelect
        v-model="filter"
        :multiple="true"
        :options="filterOptions"
        :class="$style.select"
        :placeholder="'Фильтр'"
      />
    </div>
    <div
      v-if="modeStore.mode === MODES.Payment"
      :class="[
        $style.content,
        isHasFeature('hidden-scroll-content') && $style.content_hiddenScroll,
      ]"
    >
      <VText v-if="paymentTemplateVisible" size="sm" type="secondary">Шаблоны:</VText>
      <PaymentTemplates :payments-amount-map="paymentsAmountMap" />
      <VDivider v-if="paymentTemplateVisible && paymentVisible" />
      <VText v-if="paymentVisible" size="sm" type="secondary">Расходы:</VText>
      <Payments :payments="filteredPayments" />
    </div>
    <div
      v-if="modeStore.mode === MODES.Receipt"
      :class="[
        $style.content,
        isHasFeature('hidden-scroll-content') && $style.content_hiddenScroll,
      ]"
    >
      <VButton type="default">Создать шаблон дохода</VButton>
    </div>
    <VText size="lg" :strong="true" :class="$style.total">
      {{
        modeStore.mode === MODES.Payment
          ? `Всего: ${new Intl.NumberFormat("ru-RU").format(paymentsTotalAmount)} ₽`
          : `Всего: ${new Intl.NumberFormat("ru-RU").format(0)} ₽`
      }}
    </VText>
  </div>
</template>

<style lang="scss" module>
  .base {
    display: flex;
    flex-direction: column;
    overflow: hidden;
    height: 100%;
  }

  .select {
    width: 200px;
  }

  .filter {
    font-size: var(--ksd-font-size-lg);
    display: flex;
    align-items: center;
    gap: var(--ksd-padding-sm);
    padding-inline: var(--ksd-padding);
    margin-block: var(--ksd-padding-lg);
  }

  .total {
    padding-inline: var(--ksd-padding);
    margin-block: var(--ksd-padding-lg);
  }

  .content {
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-sm);
    flex: 1;
    overflow: auto;
    padding-inline: var(--ksd-padding);
    padding-block-end: var(--ksd-padding);

    &_hiddenScroll {
      &::-webkit-scrollbar {
        width: 0;
        height: 0;
        display: none;
      }

      scrollbar-width: none;
      -ms-overflow-style: none;
    }
  }
</style>
