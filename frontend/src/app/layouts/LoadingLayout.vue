<script setup lang="ts">
  import { VLoader } from "@krainovsd/vue-ui";
  import { computed } from "vue";
  import { usePaymentsStore } from "@/entities/payments";
  import { useReceiptsStore } from "@/entities/receipts";

  const paymentsStore = usePaymentsStore();
  const receiptsStore = useReceiptsStore();
  const loading = computed(
    () =>
      paymentsStore.getPaymentTemplatesLoading ||
      paymentsStore.createPaymentTemplatesLoading ||
      paymentsStore.deletePaymentTemplatesLoading ||
      paymentsStore.updatePaymentTemplatesLoading ||
      paymentsStore.getPaymentsLoading ||
      paymentsStore.createPaymentsLoading ||
      paymentsStore.deletePaymentsLoading ||
      paymentsStore.updatePaymentsLoading ||
      receiptsStore.getReceiptTemplatesLoading ||
      receiptsStore.createReceiptTemplatesLoading ||
      receiptsStore.deleteReceiptTemplatesLoading ||
      receiptsStore.updateReceiptTemplatesLoading ||
      receiptsStore.getReceiptsLoading ||
      receiptsStore.createReceiptsLoading ||
      receiptsStore.deleteReceiptsLoading ||
      receiptsStore.updateReceiptsLoading,
  );
</script>

<template>
  <div v-if="loading" :class="$style.loading">
    <VLoader type="primary" />
  </div>
</template>

<style lang="scss" module>
  .loading {
    position: absolute;
    bottom: var(--ksd-padding-lg);
    right: var(--ksd-padding-lg);
    width: 25px;
    height: 25px;
    --ksd-loader-primary-size: 25px;
  }
</style>
