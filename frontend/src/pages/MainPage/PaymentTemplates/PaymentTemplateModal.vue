<script setup lang="ts">
  import { VButton, VInput, VInputNumber, VModal } from "@krainovsd/vue-ui";
  import { computed, ref, watch } from "vue";
  import type { IPaymentTemplate } from "@/entities/payments";
  import Label from "@/components/atoms/Label.vue";

  type Props = {
    paymentTemplate: IPaymentTemplate | undefined;
    loading: boolean;
  };
  type Emits = {
    action: [name: string, amount: number];
    delete: [];
    close: [];
  };

  const props = defineProps<Props>();
  const emit = defineEmits<Emits>();
  const open = defineModel<boolean>();
  const body = computed(() => document.body);
  const name = ref("");
  const amount = ref(0);
  const disabled = computed(() => name.value.trim().length === 0 || props.loading);

  watch(
    () => props.paymentTemplate,
    (template) => {
      if (!template) {
        name.value = "";
        amount.value = 0;
      } else {
        name.value = template.name;
        amount.value = template.amount;
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
    emit("action", name.value, amount.value);
  }
</script>

<template>
  <VModal
    v-model="open"
    :target="body"
    :header="$props.paymentTemplate ? 'Изменение шаблона расходов' : 'Создание шаблона расходов'"
    :class="$style.modal"
  >
    <template #content>
      <Label :label="'Название'">
        <VInput v-model="name" :disabled="$props.loading" :autofocus="true" />
      </Label>
      <Label :label="'Сумма'">
        <VInputNumber v-model="amount" :disabled="$props.loading" />
      </Label>
    </template>
    <template #footer>
      <VButton
        v-if="$props.paymentTemplate"
        type="primary"
        :danger="true"
        :disabled="disabled"
        :loading="$props.loading"
        @click="emit('delete')"
      >
        Удалить
      </VButton>
      <VButton type="primary" :disabled="disabled" :loading="$props.loading" @click="onAction">
        {{ $props.paymentTemplate ? "Изменить" : "Создать" }}
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
