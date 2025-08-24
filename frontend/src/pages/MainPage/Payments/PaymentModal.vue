<script setup lang="ts">
  import {
    type SelectItem,
    VButton,
    VInputNumber,
    VModal,
    VSelect,
    VTextArea,
  } from "@krainovsd/vue-ui";
  import { computed, ref, watch } from "vue";
  import type { IPayment, IPaymentTemplate } from "@/entities/payments";
  import Label from "@/components/atoms/Label.vue";

  type Props = {
    payment: IPayment | undefined;
    templates: IPaymentTemplate[];
    loading: boolean;
  };
  type Emits = {
    action: [paymentId: number, amount: number, description: string];
    delete: [];
    close: [];
  };

  const props = defineProps<Props>();
  const emit = defineEmits<Emits>();
  const open = defineModel<boolean>();
  const body = computed(() => document.body);
  const paymentId = ref<number | null>(null);
  const amount = ref(0);
  const description = ref("");
  const disabled = computed(() => props.loading || paymentId.value == undefined);
  const options = computed<SelectItem[]>(() =>
    props.templates.map((template) => ({ label: template.name, value: template.id })),
  );

  watch(
    () => props.payment,
    (payment) => {
      if (!payment) {
        paymentId.value = props.templates[0]?.id ?? null;
        description.value = "";
        amount.value = 0;
      } else {
        paymentId.value = payment.paymentId;
        amount.value = payment.amount;
        description.value = payment.description;
      }
    },
    { immediate: true },
  );

  watch(open, (open) => {
    if (!open) {
      emit("close");
    }
  });

  function onAction() {
    if (paymentId.value == undefined) return;

    emit("action", paymentId.value, amount.value, description.value);
  }
</script>

<template>
  <VModal
    v-model="open"
    :target="body"
    :header="$props.payment ? 'Изменение шаблона расходов' : 'Создание шаблона расходов'"
    :class="$style.modal"
  >
    <template #content>
      <Label :label="'Шаблон'">
        <VSelect
          v-model="paymentId"
          :options="options"
          :disabled="$props.loading"
          :autofocus="true"
        />
      </Label>
      <Label :label="'Сумма'">
        <VInputNumber v-model="amount" :disabled="$props.loading" />
      </Label>
      <Label :label="'Описание'">
        <VTextArea v-model="description" :disabled="$props.loading" :autofocus="true" />
      </Label>
    </template>
    <template #footer>
      <VButton
        v-if="$props.payment"
        type="primary"
        :danger="true"
        :disabled="disabled"
        :loading="$props.loading"
        @click="emit('delete')"
      >
        Удалить
      </VButton>
      <VButton type="primary" :disabled="disabled" :loading="$props.loading" @click="onAction">
        {{ $props.payment ? "Изменить" : "Создать" }}
      </VButton>
    </template>
  </VModal>
</template>

<style lang="scss" module>
  .modal {
    :global(.ksd-modal__header) {
      gap: var(--ksd-padding-sm);
    }
    :global(.ksd-modal__body) {
      display: flex;
      flex-direction: column;
      gap: var(--ksd-padding-sm);
    }
    :global(.ksd-modal__footer) {
      justify-content: space-between;
    }
  }
</style>
