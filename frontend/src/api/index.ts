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
  //   activeMiddlewares: ["authNoRefresh"],
  //   middlewareOptions: {
  //     authNoRefresh: {
  //       authUrl: ENDPOINTS_CONFIG.auth,
  //       errorUrl: "/error",
  //       storageTokenExpiresName: STORAGE_TOKEN_EXPIRES_NAME,
  //       queryIsRefreshTokenName: QUERY_IS_REFRESH_TOKEN_NAME,
  //       queryTokenExpiresName: QUERY_TOKEN_EXPIRES_NAME,
  //     },
  //   },
});

export const apiRequest = apiClient.requestApiWithMeta;
export const setRequestMiddlewares = apiClient.setMiddlewares;
