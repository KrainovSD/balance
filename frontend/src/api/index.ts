import { createRequestClientInstance } from "@krainovsd/js-helpers";
import { OAUTH_EXPIRES_KEY, OAUTH_REFRESH_KEY } from "@/entities/tech";

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
  oauthOptions: {
    expiresTokenStorageName: OAUTH_EXPIRES_KEY,
    onlyRefreshTokenWindowQueryName: OAUTH_REFRESH_KEY,
    refreshTokenWindowUrl: `${window.origin}/login`,
    wait: 60 * 1000,
  },
});

export const apiRequest = apiClient.requestApiWithMeta;
export const setRequestMiddlewares = apiClient.setMiddlewares;
