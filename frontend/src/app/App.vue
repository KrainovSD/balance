<script setup lang="ts">
  import { VThemeProvider } from "@krainovsd/vue-ui";
  import "@krainovsd/vue-ui/styles";
  import { THEME_CONFIG } from "@/entities/preferences";
  import "./global.scss";
  import RootLayout from "./layouts/RootLayout.vue";
  import RootProvider from "./providers/RootProvider.vue";
</script>

<template>
  <VThemeProvider :theme-config="THEME_CONFIG" :theme="'dark'" :font-size="14">
    <RootProvider>
      <RootLayout>
        <RouterView />
      </RootLayout>
    </RootProvider>
  </VThemeProvider>
</template>
