<script setup lang="ts">
  import { ref } from "vue";
  import { useRoute } from "vue-router";
  import { PAGES } from "@/entities/tech";
  import Header from "@/components/organisms/Header/Header.vue";
  import Sidebar from "@/components/organisms/Sidebar/Sidebar.vue";
  import LoadingLayout from "./LoadingLayout.vue";

  const siderOpen = ref(true);
  const route = useRoute();
</script>

<template>
  <Header v-if="route.name === PAGES.Main" @logo="siderOpen = !siderOpen" />
  <div :class="$style.wrap">
    <Sidebar v-if="route.name === PAGES.Main" :open="siderOpen" />
    <main :class="$style.root">
      <slot></slot>
    </main>
  </div>
  <LoadingLayout />
</template>

<style lang="scss" module>
  .root {
    width: 100%;
    height: 100%;
    padding: var(--ksd-padding);
    flex: 1;
    overflow: hidden;
  }

  .wrap {
    display: flex;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }
</style>
