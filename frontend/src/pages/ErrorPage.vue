<script setup lang="ts">
  import { VButton, VFlex, VText } from "@krainovsd/vue-ui";
</script>

<template>
  <VFlex :flex-align="'center'" :justify="'center'" w-full h-full vertical :gap="20">
    <VText strong :size="24">Возникла непредвиденная ошибка</VText>
    <RouterLink to="/" replace>
      <VButton>Вернуться на главную</VButton>
    </RouterLink>
  </VFlex>
</template>

<style lang="scss" module></style>
