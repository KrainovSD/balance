<script setup lang="ts">
  import { VNotification } from "@krainovsd/vue-ui";
  import { computed, useTemplateRef, watch } from "vue";
  import { useNotificationsStore } from "@/entities/notifications";

  const noticeRef = useTemplateRef("notice");
  const createNotification = computed(() => noticeRef.value?.createNotification);
  const store = useNotificationsStore();

  watch(
    createNotification,
    (value) => {
      if (value) {
        store.$patch({ _createNotificationFn: value });
      }
    },
    { immediate: true },
  );
</script>

<template>
  <VNotification
    ref="notice"
    :default-type="'error'"
    :default-duration="8"
    :max-count="3"
    :position="'top-right'"
  >
    <slot></slot>
  </VNotification>
</template>

<style lang="scss" module></style>
