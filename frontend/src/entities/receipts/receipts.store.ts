import { defineStore } from "pinia";
import { apiRequest } from "@/api";
import { ENDPOINTS } from "@/api/endpoints";
import { apiErrorLayer } from "@/lib/api-error-layers";
import { isHasFeature } from "@/lib/check-feature";
import { useNotificationsStore } from "../notifications";
import { REQUEST_DELAY } from "../tech";
import type { IReceipt, IReceiptTemplate, IReceiptsStore } from "./receipts.types";

export const useReceiptsStore = defineStore("receipts", {
  state: (): IReceiptsStore => {
    return {
      getReceiptTemplatesLoading: false,
      createReceiptTemplatesLoading: false,
      deleteReceiptTemplatesLoading: false,
      updateReceiptTemplatesLoading: false,
      updateReceiptsLoading: false,
      createReceiptsLoading: false,
      deleteReceiptsLoading: false,
      getReceiptsLoading: false,
      receipts: [],
      receiptTemplates: [],
    };
  },
  getters: {},
  actions: {
    async getReceiptTemplates() {
      if (this.getReceiptTemplatesLoading) return;
      this.getReceiptTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IReceiptTemplate[]>({
            method: "GET",
            path: ENDPOINTS.receiptTemplates,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось получить шаблоны доходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.receiptTemplates = result.data;
      }

      this.getReceiptTemplatesLoading = false;
    },
    async deleteReceiptTemplates(ids: number[]) {
      if (this.deleteReceiptTemplatesLoading) return;
      this.deleteReceiptTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<boolean>({
            method: "DELETE",
            path: ENDPOINTS.receiptTemplates,
            body: ids,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось удалить шаблоны доходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.receiptTemplates = this.receiptTemplates.filter((temp) => !ids.includes(temp.id));
        this.receipts = this.receipts.filter((rec) => !ids.includes(rec.receiptId));
      }

      this.deleteReceiptTemplatesLoading = false;
    },
    async updateReceiptTemplate(id: number, name: string, amount: number) {
      if (this.updateReceiptTemplatesLoading) return;
      this.updateReceiptTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IReceiptTemplate[]>({
            method: "PATCH",
            path: ENDPOINTS.receiptTemplate(id),
            body: {
              name,
              amount,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось обновить шаблон доходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.receiptTemplates = this.receiptTemplates.map<IReceiptTemplate>((temp) => {
          if (temp.id === id) {
            return {
              amount,
              id,
              name,
            };
          }

          return temp;
        });
        this.receipts = this.receipts.map<IReceipt>((rec) => {
          if (rec.receiptId === id) {
            return {
              ...rec,
              name,
            };
          }

          return rec;
        });
      }

      this.updateReceiptTemplatesLoading = false;
    },
    async createReceiptTemplate(name: string, amount: number) {
      if (this.createReceiptTemplatesLoading) return;
      this.createReceiptTemplatesLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<number>({
            method: "POST",
            path: ENDPOINTS.receiptTemplates,
            body: {
              name,
              amount,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),

        () => {
          notificationsStore.createMessage("Не удалось создать шаблон доходов", {
            type: "error",
          });
        },
      );

      if (result) {
        this.receiptTemplates = [...this.receiptTemplates, { amount, name, id: result.data }];
      }

      this.createReceiptTemplatesLoading = false;
    },

    async getReceipts() {
      if (this.getReceiptsLoading) return;
      this.getReceiptsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IReceiptTemplate[]>({
            method: "GET",
            path: ENDPOINTS.receipts,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),

        () => {
          notificationsStore.createMessage("Не удалось получить доходы", {
            type: "error",
          });
        },
      );

      if (result) {
        this.receiptTemplates = result.data;
      }

      this.getReceiptsLoading = false;
    },
    async deleteReceipts(ids: number[]) {
      if (this.deleteReceiptsLoading) return;
      this.deleteReceiptsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<boolean>({
            method: "DELETE",
            path: ENDPOINTS.receipts,
            body: ids,
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось удалить доходы", {
            type: "error",
          });
        },
      );

      if (result) {
        this.receipts = this.receipts.filter((receipt) => !ids.includes(receipt.id));
      }

      this.deleteReceiptsLoading = false;
    },
    async updateReceipt(id: number, receiptId: number, amount: number, description: string) {
      if (this.updateReceiptsLoading) return;
      this.updateReceiptsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<IReceipt[]>({
            method: "PATCH",
            path: ENDPOINTS.receipt(id),
            body: {
              receiptId,
              amount,
              description,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось обновить доходы", {
            type: "error",
          });
        },
      );

      if (result) {
        this.receipts = this.receipts.map<IReceipt>((receipt) => {
          if (receipt.id === id) {
            const name =
              this.receiptTemplates.find((temp) => temp.id === receiptId)?.name ?? receipt.name;

            return {
              id,
              amount,
              description,
              receiptId,
              date: receipt.date,
              name,
            };
          }

          return receipt;
        });
      }

      this.updateReceiptsLoading = false;
    },
    async createReceipt(receiptId: number, amount: number, description: string) {
      if (this.createReceiptsLoading) return;
      this.createReceiptsLoading = true;
      const notificationsStore = useNotificationsStore();

      const result = await apiErrorLayer(
        async () =>
          apiRequest<number>({
            method: "POST",
            path: ENDPOINTS.receipts,
            body: {
              receiptId,
              amount,
              description,
            },
            delay: isHasFeature("request-delay") ? REQUEST_DELAY : undefined,
          }),
        () => {
          notificationsStore.createMessage("Не удалось создать доходы", {
            type: "error",
          });
        },
      );

      if (result) {
        const name = this.receiptTemplates.find((receipt) => receipt.id === receiptId)?.name ?? "";
        this.receipts = [
          { amount, date: new Date().toISOString(), description, id: result.data, receiptId, name },
          ...this.receipts,
        ];
      }

      this.createReceiptsLoading = false;
    },
  },
});
