<script setup lang="ts">
  import { VMessage } from "@krainovsd/vue-ui";
  import { computed, useTemplateRef, watch } from "vue";
  import { useNotificationsStore } from "@/entities/notifications";

  const messageRef = useTemplateRef("message");
  const createMessage = computed(() => messageRef.value?.createMessage);
  const store = useNotificationsStore();

  watch(
    createMessage,
    (value) => {
      if (value) {
        store.$patch({ _createMessageFn: value });
      }
    },
    { immediate: true },
  );
</script>

<template>
  <VMessage ref="message" :default-duration="8" :default-type="'error'" :max-count="5">
    <slot></slot>
  </VMessage>
</template>

<style lang="scss" module></style>
