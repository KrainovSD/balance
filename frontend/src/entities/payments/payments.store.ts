import { defineStore } from "pinia";
import { apiRequest } from "@/api";
import { ENDPOINTS } from "@/api/endpoints";
import { apiErrorLayer } from "@/lib/api-error-layers";
import { isHasFeature } from "@/lib/check-feature";
import { useNotificationsStore } from "../notifications";
import { REQUEST_DELAY } from "../tech";
import type { IPayment, IPaymentTemplate, IPaymentsStore } from "./payments.types";

export const usePaymentsStore = defineStore("payments", {
  state: (): IPaymentsStore => {
    return {
      getPaymentTemplatesLoading: false,
      createPaymentTemplatesLoading: false,
      deletePaymentTemplatesLoading: false,
      updatePaymentTemplatesLoading: false,
      updatePaymentsLoading: false,
      createPaymentsLoading: false,
      deletePaymentsLoading: false,
      getPaymentsLoading: false,
      payments: [],
      paymentTemplates: [],
    };
  },
  getters: {},
  actions: {
    async getPaymentTemplates() {
      if (this.getPaymentTemplatesLoading) return;
      this.getPaymentTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IPaymentTemplate[]>({
            method: "GET",
            path: ENDPOINTS.paymentTemplates,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
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

      this.getPaymentTemplatesLoading = false;
    },
    async deletePaymentTemplates(ids: number[]) {
      if (this.deletePaymentTemplatesLoading) return;
      this.deletePaymentTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<boolean>({
            method: "DELETE",
            path: ENDPOINTS.paymentTemplates,
            body: ids,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось удалить шаблоны расходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.paymentTemplates = this.paymentTemplates.filter((temp) => !ids.includes(temp.id));
      }

      this.deletePaymentTemplatesLoading = false;

      return !!result;
    },
    async updatePaymentTemplate(id: number, name: string, amount: number) {
      if (this.updatePaymentTemplatesLoading) return;
      this.updatePaymentTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<boolean>({
            method: "PATCH",
            path: ENDPOINTS.paymentTemplate(id),
            body: {
              name,
              amount,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось обновить шаблон расходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.paymentTemplates = this.paymentTemplates.map<IPaymentTemplate>((temp) => {
          if (temp.id === id) {
            return {
              amount,
              id,
              name,
            };
          }

          return temp;
        });
      }

      this.updatePaymentTemplatesLoading = false;

      return !!result;
    },
    async createPaymentTemplate(name: string, amount: number) {
      if (this.createPaymentTemplatesLoading) return;
      this.createPaymentTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<number>({
            method: "POST",
            path: ENDPOINTS.paymentTemplates,
            body: {
              name,
              amount,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось создать шаблон расходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.paymentTemplates = [...this.paymentTemplates, { amount, name, id: result.data }];
      }

      this.createPaymentTemplatesLoading = false;

      return !!result;
    },

    async getPayments() {
      if (this.getPaymentsLoading) return;
      this.getPaymentsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IPayment[]>({
            method: "GET",
            path: ENDPOINTS.payments,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось получить расходы", {
            type: "error",
          });
        },
      );

      if (result) {
        this.payments = result.data;
      }

      this.getPaymentsLoading = false;
    },
    async deletePayments(ids: number[]) {
      if (this.deletePaymentsLoading) return;
      this.deletePaymentsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<boolean>({
            method: "DELETE",
            path: ENDPOINTS.payments,
            body: ids,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось удалить расходы", {
            type: "error",
          });
        },
      );

      if (result) {
        this.payments = this.payments.filter((payment) => !ids.includes(payment.id));
      }

      this.deletePaymentsLoading = false;

      return !!result;
    },
    async updatePayment(id: number, paymentId: number, amount: number, description: string) {
      if (this.updatePaymentsLoading) return;
      this.updatePaymentsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<boolean>({
            method: "PATCH",
            path: ENDPOINTS.payment(id),
            body: {
              paymentId,
              amount,
              description,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось обновить расходы", {
            type: "error",
          });
        },
      );

      if (result) {
        this.payments = this.payments.map<IPayment>((payment) => {
          if (payment.id === id) {
            const name =
              this.paymentTemplates.find((temp) => temp.id === paymentId)?.name ?? payment.name;

            return {
              id,
              amount,
              description,
              paymentId,
              date: payment.date,
              name,
            };
          }

          return payment;
        });
      }

      this.updatePaymentsLoading = false;

      return !!result;
    },
    async createPayment(paymentId: number, amount: number, description: string) {
      if (this.createPaymentsLoading) return;
      this.createPaymentsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<number>({
            method: "POST",
            path: ENDPOINTS.payments,
            body: {
              paymentId,
              amount,
              description,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось создать расходы", {
            type: "error",
          });
        },
      );

      if (result) {
        const name = this.paymentTemplates.find((payment) => payment.id === paymentId)?.name ?? "";
        this.payments = [
          ...this.payments,
          { amount, date: new Date().toISOString(), description, id: result.data, paymentId, name },
        ];
      }

      this.createPaymentsLoading = false;

      return !!result;
    },
  },
});
