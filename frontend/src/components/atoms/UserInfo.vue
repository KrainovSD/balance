<script setup lang="ts">
  import { VText } from "@krainovsd/vue-ui";
  import UserAvatar from "./UserAvatar.vue";

  type Props = {
    bigName: string;
    tag: string;
    url?: string;
    size?: "lg" | "sm" | "default";
  };

  defineProps<Props>();
</script>

<template>
  <div :class="$style.user">
    <UserAvatar
      :big-name="$props.bigName"
      :size="$props.size === 'sm' ? 'default' : $props.size"
      :url="$props.url"
    />
    <div :class="$style.info">
      <VText :class="$style.text" :size="$props.size" :weight="800">{{ $props.bigName }}</VText>
      <VText :class="$style.text" :size="$props.size" :type="'secondary'">{{
        `@${$props.tag}`
      }}</VText>
    </div>
  </div>
</template>

<style lang="scss" module>
  .user {
    display: flex;
    align-items: center;
    gap: var(--ksd-padding-xs);
  }

  .info {
    display: flex;
    flex-direction: column;
    align-items: start;
    justify-content: start;
    gap: var(--ksd-padding-xxs);
  }

  .text {
    line-height: 1;
  }
</style>
