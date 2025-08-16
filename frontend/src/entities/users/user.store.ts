import { defineStore } from "pinia";
import { apiRequest } from "@/api";
import { ENDPOINTS } from "@/api/endpoints";
import { apiErrorLayer } from "@/lib/api-error-layers";
import { useNotificationsStore } from "../notifications";
import type { IUser, IUsersStore } from "./users.types";

export const useUsersStore = defineStore("users", {
  state: (): IUsersStore => {
    return {
      userInfo: undefined,
      getUserInfoLoading: false,
    };
  },
  getters: {},
  actions: {
    async getUsers() {
      if (this.getUserInfoLoading) return;
      this.getUserInfoLoading = true;
      const notificationStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IUser>({
            method: "GET",
            path: ENDPOINTS.userInfo,
          }),
        () => {
          notificationStore.createMessage("Не удалось получить информацию о пользователе");
        },
      );

      this.userInfo = result?.data ?? null;
      this.getUserInfoLoading = false;
    },
  },
});
