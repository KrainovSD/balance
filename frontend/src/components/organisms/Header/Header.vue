<script setup lang="ts">
  import { VSettingOutlined, VUploadOutlined } from "@krainovsd/vue-icons";
  import { type DropDownMenuItem, VDropDown } from "@krainovsd/vue-ui";
  import { computed, h, useCssModule } from "vue";
  import { useUsersStore } from "@/entities/users";
  import UserInfo from "@/components/atoms/UserInfo.vue";

  const usersStore = useUsersStore();
  const styles = useCssModule();

  const userMenu = computed<DropDownMenuItem[]>(() =>
    usersStore.userInfo
      ? [
          { key: "1", label: usersStore.userInfo.name, noInteractive: true },
          { key: "2", divider: true },
          { key: "3", label: "Настройки", icon: h(VSettingOutlined) },
          { key: "4", label: "Выход", icon: h(VUploadOutlined, { class: styles.exit }) },
        ]
      : [],
  );

  function keyboardProfile(event: KeyboardEvent) {
    if (event.key === " " || event.key === "Enter") {
      const target = event.currentTarget as HTMLElement;
      const mouseEvent = new MouseEvent("click");
      target.dispatchEvent(mouseEvent);
    }
  }
</script>

<template>
  <header :class="$style.root">
    <VDropDown v-if="usersStore.userInfo" :menu="userMenu">
      <UserInfo
        :big-name="usersStore.userInfo.name"
        :tag="usersStore.userInfo.username"
        size="default"
      />
    </VDropDown>
  </header>
</template>

<style lang="scss" module>
  .root {
    height: 85px;
    padding: var(--ksd-padding-sm) var(--ksd-padding);
    border-bottom: var(--ksd-line-width) var(--ksd-line-type) var(--ksd-border-color);
    display: flex;
    align-items: center;
    gap: var(--ksd-margin-xxl);
  }

  .logo {
    font-size: 45px !important;
    padding-top: 5px !important;
    padding-bottom: 5px !important;
    height: fit-content !important;
    line-height: 1 !important;
  }

  .user {
    margin-left: auto;
    display: flex;
    gap: 10px;
    align-items: center;
    border-radius: var(--ksd-border-radius-sm);
    cursor: pointer;

    &:focus-visible {
      outline: var(--ksd-outline-width) var(--ksd-outline-type) var(--ksd-outline-color);
      outline-offset: 1px;
      transition:
        outline-offset 0s,
        outline 0s;
    }
  }

  .avatar {
    width: var(--ksd-control-height-lg);
    height: var(--ksd-control-height-lg);
    min-width: var(--ksd-control-height-lg);
    min-height: var(--ksd-control-height-lg);
    border-radius: 100%;
    border: var(--ksd-line-width) var(--ksd-line-type) var(--ksd-border-color);
  }

  .exit {
    transform: rotate(90deg);
  }
</style>
