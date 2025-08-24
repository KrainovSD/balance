import { extractQueries } from "@krainovsd/js-helpers";
import { type NavigationGuardWithThis, createRouter, createWebHistory } from "vue-router";
import LoginPage from "@/pages/LoginPage.vue";
import MainPage from "@/pages/MainPage/MainPage.vue";
import { OAUTH_REFRESH_KEY, PAGES } from "@/entities/tech";
import { useUsersStore } from "@/entities/users";

const authGuard: NavigationGuardWithThis<undefined> = async () => {
  const queries = extractQueries();
  if (queries[OAUTH_REFRESH_KEY] != undefined) return { name: PAGES.Login };

  const usersStore = useUsersStore();

  if (usersStore.userInfo === undefined) {
    await usersStore.getUsers();
  }

  if (!usersStore.userInfo) {
    return { name: PAGES.Login };
  }
};
const loginGuard: NavigationGuardWithThis<undefined> = async () => {
  const queries = extractQueries();
  if (queries[OAUTH_REFRESH_KEY] != undefined) return;

  const usersStore = useUsersStore();
  if (usersStore.userInfo === undefined) {
    await usersStore.getUsers();
  }

  if (usersStore.userInfo) {
    return { name: PAGES.Main };
  }
};

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: PAGES.Main, component: MainPage, beforeEnter: authGuard },
    { path: "/login", name: PAGES.Login, component: LoginPage, beforeEnter: loginGuard },
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
