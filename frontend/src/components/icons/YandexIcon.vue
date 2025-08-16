<script setup lang="ts">
  import type { IconProps } from "@krainovsd/vue-icons";

  const { size, color = "currentColor" } = defineProps<IconProps>();
</script>

<template>
  <svg
    width="1em"
    height="1em"
    fill="none"
    focusable="false"
    class="ksd-icon"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 1080 1080"
    :style="{ fontSize: size != undefined ? `${size}px` : size }"
  >
    <circle :fill="color" cx="540" cy="540" r="538.9" />
    <path
      fill="black"
      d="M735,878.3H616.7V293.2H564c-96.6,0-147.2,48.3-147.2,120.4c0,81.8,34.9,119.7,107.1,168l59.5,40.1
	l-171,256.5H285.2l153.9-229c-88.5-63.2-138.3-124.9-138.3-229c0-130.1,90.7-218.6,262.4-218.6h171v676.5H735z"
    />
  </svg>
</template>

<style lang="scss" module></style>
