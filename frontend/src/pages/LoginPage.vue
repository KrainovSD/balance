<script setup lang="ts">
  import { VGithubFilled, VGitlabFilled, VGoogleOutlined } from "@krainovsd/vue-icons";
  import { VButton, VText } from "@krainovsd/vue-ui";
  import { ACTIVE_OAUTH_PROVIDERS, type IOauthProvider, OAUTH_PROVIDERS } from "@/entities/tech";
  import { ENDPOINTS } from "@/api/endpoints";
  import YandexIcon from "@/components/icons/YandexIcon.vue";

  function login(provider: IOauthProvider) {
    void window.location.replace(ENDPOINTS.auth(provider));
  }
</script>

<template>
  <div :class="$style.root">
    <VText size="lg" :strong="true" type="secondary" :class="$style.text"> Войти с помощью: </VText>
    <div :class="$style.buttonWrap">
      <VButton
        v-for="provider in ACTIVE_OAUTH_PROVIDERS"
        :key="provider"
        size="large"
        @click="login(provider)"
      >
        <template #icon>
          <VGoogleOutlined v-if="provider === OAUTH_PROVIDERS.Google" :class="$style.icon" />
          <YandexIcon v-if="provider === OAUTH_PROVIDERS.Yandex" :class="$style.icon" />
          <VGithubFilled v-if="provider === OAUTH_PROVIDERS.Github" :class="$style.icon" />
          <VGitlabFilled v-if="provider === OAUTH_PROVIDERS.Gitlab" :class="$style.icon" />
        </template>
      </VButton>
    </div>
  </div>
</template>

<style lang="scss" module>
  .root {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    width: 100%;
    gap: var(--ksd-padding);
  }

  .text {
    text-align: center;
    display: flex;
  }

  .buttonWrap {
    display: flex;
    align-items: center;
    gap: var(--ksd-padding-sm);
  }

  .icon {
    font-size: 20px;
  }
</style>
