<script setup lang="ts"></script>

<template>
  <VFlex :flex-align="'center'" :justify="'center'" w-full h-full vertical :gap="20">
    <VText strong :size="24">Страница не найдена</VText>
    <RouterLink to="/" replace>
      <VButton>Вернуться на главную</VButton>
    </RouterLink>
  </VFlex>
</template>

<style lang="scss" scoped></style>
