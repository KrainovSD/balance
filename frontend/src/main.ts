import { createPinia } from "pinia";
import { createApp } from "vue";
import App from "./app/App.vue";
import { router } from "./app/router";

createApp(App).use(router).use(createPinia()).mount("#app");
