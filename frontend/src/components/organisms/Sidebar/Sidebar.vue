<script setup lang="ts"></script>

<template>
  <aside :class="$style.root">sidebar</aside>
</template>

<style lang="scss" module>
  .root {
    --width-max: 310px;
    --width-min: 53px;
    height: 100%;
    width: 310px;
    display: flex;
    flex-direction: column;

    &[data-open="true"] {
      animation: appear var(--ksd-transition-slow) ease both;
    }
    &[data-open="false"] {
      animation: disappear var(--ksd-transition-slow) ease both;
    }

    &.appear {
      width: var(--width-max);
    }
    &.disappear {
      width: var(--width-min);
    }
  }

  @keyframes appear {
    from {
      width: var(--width-min);
    }
    to {
      width: var(--width-max);
    }
  }
  @keyframes disappear {
    from {
      width: var(--width-max);
    }
    to {
      width: var(--width-min);
    }
  }
</style>
