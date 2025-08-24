import { defineStore } from "pinia";
import { apiRequest } from "@/api";
import { ENDPOINTS } from "@/api/endpoints";
import { apiErrorLayer } from "@/lib/api-error-layers";
import { useNotificationsStore } from "../notifications";
import type { IPaymentTemplate, IPaymentsStore } from "./payments.types";

export const usePaymentsStore = defineStore("payments", {
  state: (): IPaymentsStore => {
    return {
      getPaymentTemplateLoading: false,
      paymentTemplates: [],
    };
  },
  getters: {},
  actions: {
    async getPaymentTemplates() {
      if (this.getPaymentTemplateLoading) return;
      this.getPaymentTemplateLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IPaymentTemplate[]>({
            method: "GET",
            path: ENDPOINTS.userInfo,
          }),
        () => {
          notificationsStore.createMessage("Не удалось получить шаблоны расходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.paymentTemplates = result.data;
      }

      this.getPaymentTemplateLoading = false;
    },
  },
});
