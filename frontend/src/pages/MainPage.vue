<script setup lang="ts">
  import { VButton } from "@krainovsd/vue-ui";
  import { usePaymentsStore } from "@/entities/payments";

  const paymentsStore = usePaymentsStore();

  function test() {
    void paymentsStore.getPaymentTemplates();
  }
</script>

<template>
  <VButton @click="test"> Пыщь </VButton>
</template>

<style lang="scss" module></style>
