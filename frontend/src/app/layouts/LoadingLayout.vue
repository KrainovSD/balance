<script setup lang="ts">
  import { VLoader } from "@krainovsd/vue-ui";
  import { computed } from "vue";

  const loading = computed(() => false);
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
