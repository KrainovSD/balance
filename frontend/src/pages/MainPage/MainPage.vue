<script setup lang="ts">
  import { VButton, VDivider, VText } from "@krainovsd/vue-ui";
  import { computed, onMounted } from "vue";
  import { MODES, useModeStore } from "@/entities/mode";
  import { usePaymentsStore } from "@/entities/payments";
  import { useReceiptsStore } from "@/entities/receipts";
  import PaymentTemplates from "./PaymentTemplates/PaymentTemplates.vue";
  import Payments from "./Payments/Payments.vue";

  const paymentsStore = usePaymentsStore();
  const receiptsStore = useReceiptsStore();
  const modeStore = useModeStore();

  const paymentTemplateVisible = computed(() => paymentsStore.paymentTemplates.length > 0);
  const paymentVisible = computed(() => paymentsStore.payments.length > 0);

  onMounted(() => {
    void paymentsStore.getPayments();
    void receiptsStore.getReceipts();
    void paymentsStore.getPaymentTemplates();
    void receiptsStore.getReceiptTemplates();
  });
</script>

<template>
  <div :class="$style.base">
    <div v-if="modeStore.mode === MODES.Payment" :class="$style.payments">
      <VText v-if="paymentTemplateVisible" size="sm" type="secondary">Шаблоны:</VText>
      <PaymentTemplates />
      <VDivider v-if="paymentTemplateVisible && paymentVisible" />
      <VText v-if="paymentVisible" size="sm" type="secondary">Расходы:</VText>
      <Payments />
    </div>
    <div v-if="modeStore.mode === MODES.Receipt" :class="$style.receipts">
      <VButton type="default">Создать шаблон дохода</VButton>
    </div>
  </div>
</template>

<style lang="scss" module>
  .base {
    display: flex;
    gap: var(--ksd-padding-lg);
  }

  .payments {
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-sm);
    flex: 1;
  }

  .receipts {
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-sm);
    flex: 1;
  }
</style>
