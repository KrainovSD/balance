import type { Message, Notification } from "@krainovsd/vue-ui";

export type INotificationsStore = {
  _createMessageFn:
    | ((
        text: string,
        opts?: Pick<Message, "duration" | "type"> & {
          id?: number;
        },
      ) => number)
    | undefined;
  _createNotificationFn:
    | ((
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
      ) => number)
    | undefined;
};
