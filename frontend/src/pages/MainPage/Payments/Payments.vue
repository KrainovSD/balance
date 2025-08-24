<script setup lang="ts">
  import { VMinusOutlined } from "@krainovsd/vue-icons";
  import { VButton, VTooltip } from "@krainovsd/vue-ui";
  import { computed, ref, watch } from "vue";
  import { usePaymentsStore } from "@/entities/payments";
  import Payment from "./Payment.vue";
  import PaymentModal from "./PaymentModal.vue";

  const paymentsStore = usePaymentsStore();
  const openPaymentModal = ref(false);
  const editedPaymentId = ref<null | number>(null);
  const editedPayment = computed(() =>
    paymentsStore.payments.find((temp) => temp.id === editedPaymentId.value),
  );

  function onPaymentAction(paymentId: number, amount: number, description: string) {
    if (editedPaymentId.value != undefined) {
      void paymentsStore
        .updatePayment(editedPaymentId.value, paymentId, amount, description)
        .then((success) => {
          if (success) {
            openPaymentModal.value = false;
          }
        });
    } else {
      void paymentsStore.createPayment(paymentId, amount, description).then((success) => {
        if (success) {
          openPaymentModal.value = false;
        }
      });
    }
  }

  function onPaymentDelete() {
    if (editedPayment.value == undefined) return;

    void paymentsStore.deletePayments([editedPayment.value.id]).then((success) => {
      if (success) {
        openPaymentModal.value = false;
      }
    });
  }

  watch(
    openPaymentModal,
    (open) => {
      if (!open) {
        editedPaymentId.value = null;
      }
    },
    { immediate: true },
  );
</script>

<template>
  <PaymentModal
    v-model="openPaymentModal"
    :loading="paymentsStore.createPaymentsLoading || paymentsStore.updatePaymentsLoading"
    :payment="editedPayment"
    :templates="paymentsStore.paymentTemplates"
    @action="onPaymentAction"
    @delete="onPaymentDelete"
  />

  <VTooltip :text="'Создать расход'">
    <VButton
      size="large"
      shape="default"
      :class="$style.button"
      :disabled="paymentsStore.paymentTemplates.length === 0"
      @click="openPaymentModal = true"
      @pointerdown.stop=""
    >
      <template #icon>
        <VMinusOutlined :class="$style.icon" />
      </template>
    </VButton>
  </VTooltip>
  <Payment
    v-for="payment in paymentsStore.payments"
    :key="payment.id"
    :payment="payment"
    :used="30"
    @open="
      (id) => {
        editedPaymentId = id;
        openPaymentModal = true;
      }
    "
  />
</template>

<style lang="scss" module>
  .icon {
    font-size: var(--ksd-font-size-lg);
  }

  .button {
    position: absolute;
    bottom: 20px;
    right: 80px;
  }
</style>
