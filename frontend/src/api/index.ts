import { createRequestClientInstance } from "@krainovsd/js-helpers";

const apiClient = createRequestClientInstance({
  client: window.fetch.bind(window),
  activePostMiddlewares: ["logger"],
  postMiddlewaresOptions: {
    logger: {
      filterStatus(status) {
        return status >= 400;
      },
    },
  },
  activeMiddlewares: ["oauth"],
  middlewareOptions: {
    oauth: {},
  },
});

export const apiRequest = apiClient.requestApiWithMeta;
export const setRequestMiddlewares = apiClient.setMiddlewares;
