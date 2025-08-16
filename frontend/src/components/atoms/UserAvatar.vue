<script setup lang="ts">
  import { VText } from "@krainovsd/vue-ui";
  import { computed, useCssModule } from "vue";

  type Props = {
    bigName: string;
    label?: string;
    url?: string;
    size?: "lg" | "sm" | "default";
  };

  const props = withDefaults(defineProps<Props>(), {
    size: "default",
    url: undefined,
    label: undefined,
  });
  const styles = useCssModule();
  const initials = computed(() => {
    return (
      props.label ??
      props.bigName
        .split(" ")
        .map((name) => name.charAt(0)?.toUpperCase?.())
        ?.toSpliced?.(2)
        .join("")
    );
  });
  const classes = computed(() => [styles[`size-${props.size}`]]);
</script>

<template>
  <div
    :class="[$style.avatar, classes]"
    :style="{ backgroundImage: $props.url ? `url(${$props.url})` : undefined }"
  >
    <VText v-if="!$props.url" :size="$props.size"> {{ initials }} </VText>
  </div>
</template>

<style lang="scss" module>
  .avatar {
    min-width: var(--ksd-control-height);
    min-height: var(--ksd-control-height);
    width: var(--ksd-control-height);
    height: var(--ksd-control-height);
    border: var(--ksd-line-width) var(--ksd-line-type) var(--ksd-border-color);
    border-radius: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
    background-color: var(--ksd-bg-modal-color);

    &.size-sm {
      width: var(--ksd-control-height-sm);
      height: var(--ksd-control-height-sm);
    }
    &.size-lg {
      width: var(--ksd-control-height-lg);
      height: var(--ksd-control-height-lg);
    }
  }
</style>
