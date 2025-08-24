import { defineStore } from "pinia";
import { apiRequest } from "@/api";
import { ENDPOINTS } from "@/api/endpoints";
import { apiErrorLayer } from "@/lib/api-error-layers";
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

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IUser>({
            method: "GET",
            path: ENDPOINTS.userInfo,
            refetchNoAuth: false,
          }),
        () => {},
      );

      this.userInfo = result?.data ?? null;
      this.getUserInfoLoading = false;
    },
  },
});
