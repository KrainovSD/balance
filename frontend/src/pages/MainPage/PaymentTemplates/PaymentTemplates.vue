<script setup lang="ts">
  import { VButton, VTooltip } from "@krainovsd/vue-ui";
  import { computed, ref, watch } from "vue";
  import { usePaymentsStore } from "@/entities/payments";
  import TemplateMinusIcon from "@/components/icons/TemplateMinusIcon.vue";
  import PaymentTemplate from "./PaymentTemplate.vue";
  import PaymentTemplateModal from "./PaymentTemplateModal.vue";

  type Props = {
    paymentsAmountMap: Record<string, number | undefined>;
  };

  defineProps<Props>();
  const paymentsStore = usePaymentsStore();
  const openPaymentTemplateModal = ref(false);
  const editedPaymentTemplateId = ref<null | number>(null);
  const editedPaymentTemplate = computed(() =>
    paymentsStore.paymentTemplates.find((temp) => temp.id === editedPaymentTemplateId.value),
  );

  function onPaymentTemplateAction(name: string, amount: number) {
    if (editedPaymentTemplateId.value != undefined) {
      void paymentsStore
        .updatePaymentTemplate(editedPaymentTemplateId.value, name, amount)
        .then((success) => {
          if (success) {
            openPaymentTemplateModal.value = false;
          }
        });
    } else {
      void paymentsStore.createPaymentTemplate(name, amount).then((success) => {
        if (success) {
          openPaymentTemplateModal.value = false;
        }
      });
    }
  }

  function onPaymentTemplateDelete() {
    if (editedPaymentTemplate.value == undefined) return;

    void paymentsStore.deletePaymentTemplates([editedPaymentTemplate.value.id]).then((success) => {
      if (success) {
        openPaymentTemplateModal.value = false;
      }
    });
  }

  watch(
    openPaymentTemplateModal,
    (open) => {
      if (!open) {
        editedPaymentTemplateId.value = null;
      }
    },
    { immediate: true },
  );
</script>

<template>
  <PaymentTemplateModal
    v-model="openPaymentTemplateModal"
    :loading="
      paymentsStore.createPaymentTemplatesLoading || paymentsStore.updatePaymentTemplatesLoading
    "
    :payment-template="editedPaymentTemplate"
    @action="onPaymentTemplateAction"
    @delete="onPaymentTemplateDelete"
  />

  <Teleport to="#app">
    <VTooltip :text="'Создать шаблон расхода'">
      <VButton
        size="large"
        shape="default"
        :class="$style.button"
        @click="openPaymentTemplateModal = true"
        @pointerdown.stop=""
      >
        <template #icon>
          <TemplateMinusIcon :class="$style.icon" />
        </template>
      </VButton>
    </VTooltip>
  </Teleport>
  <PaymentTemplate
    v-for="template in paymentsStore.paymentTemplates"
    :key="template.id"
    :template="template"
    :used="paymentsAmountMap[template.id] ?? 0"
    @open="
      (id) => {
        editedPaymentTemplateId = id;
        openPaymentTemplateModal = true;
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
    right: 20px;
  }
</style>
