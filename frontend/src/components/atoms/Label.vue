<script setup lang="ts">
  import { VText } from "@krainovsd/vue-ui";

  type Props = { label: string };

  defineProps<Props>();
</script>

<template>
  <div :class="$style.base">
    <VText :type="'secondary'" :ellipsis="true" size="sm">{{ $props.label }}</VText>
    <slot></slot>
  </div>
</template>

<style lang="scss" module>
  .base {
    display: flex;
    flex-direction: column;
    gap: var(--ksd-padding-xxs);
  }
</style>
