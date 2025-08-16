import type { Message, Notification } from "@krainovsd/vue-ui";
import { defineStore } from "pinia";
import type { INotificationsStore } from "./notifications.types";

export const useNotificationsStore = defineStore("notifications", {
  state: (): INotificationsStore => {
    return {
      _createMessageFn: undefined,
      _createNotificationFn: undefined,
    };
  },
  actions: {
    createMessage(
      text: string,
      opts?: Pick<Message, "duration" | "type"> & {
        id?: number;
      },
    ) {
      return this._createMessageFn?.(text, opts);
    },
    createNotification(
      text: string,
      title: string,
      opts?: Pick<
        Notification,
        | "duration"
        | "type"
        | "cancelButton"
        | "okButton"
        | "cancelButtonHandler"
        | "okButtonHandler"
        | "onClose"
      > & {
        id?: number;
      },
    ) {
      return this._createNotificationFn?.(text, title, opts);
    },
  },
});
