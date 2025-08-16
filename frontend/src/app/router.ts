import { type NavigationGuardWithThis, createRouter, createWebHistory } from "vue-router";
import LoginPage from "@/pages/LoginPage.vue";
import MainPage from "@/pages/MainPage.vue";
import { PAGES } from "@/entities/tech";
import { useUsersStore } from "@/entities/users";
import { extractQueries } from "@/lib/extract-queries";

const authGuard: NavigationGuardWithThis<undefined> = async () => {
  const usersStore = useUsersStore();

  if (usersStore.userInfo === undefined) {
    await usersStore.getUsers();
  }

  if (!usersStore.userInfo) {
    return { name: PAGES.Login };
  }
};
const loginGuard: NavigationGuardWithThis<undefined> = async () => {
  const usersStore = useUsersStore();
  if (usersStore.userInfo === undefined) {
    await usersStore.getUsers();
  }

  if (usersStore.userInfo) {
    return { name: PAGES.Main };
  }
};
const authProcess: NavigationGuardWithThis<undefined> = async () => {
  const usersStore = useUsersStore();

  if (usersStore.userInfo === undefined) {
    await usersStore.getUsers();
  }

  const queries = extractQueries();
};

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: PAGES.Main, component: MainPage, beforeEnter: authGuard },
    { path: "/login", name: PAGES.Login, component: LoginPage, beforeEnter: loginGuard },
    { path: "/auth", component: LoginPage, beforeEnter: authProcess },
    { path: "/error", name: PAGES.Error, component: () => import("@/pages/ErrorPage.vue") },
    {
      path: "/:pathMatch(.*)*",
      name: PAGES.NotFound,
      component: MainPage,
    },
  ],
  scrollBehavior: () => {
    return { top: 0, behavior: "smooth" };
  },
});
