import type { IOauthProvider } from "./tech.types";

export const PAGES = {
  NotFound: "not_found",
  Error: "error",
  Main: "main",
  Login: "login",
} as const;

export const OAUTH_PROVIDERS = {
  Google: "google",
  Yandex: "yandex",
  Github: "github",
  Gitlab: "gitlab",
} as const;

export const ACTIVE_OAUTH_PROVIDERS: IOauthProvider[] = [
  OAUTH_PROVIDERS.Google,
  OAUTH_PROVIDERS.Yandex,
  OAUTH_PROVIDERS.Github,
  OAUTH_PROVIDERS.Gitlab,
];

export const AUTH_CHANNEL_ID = "auth_channel";

export const OAUTH_EXPIRES_KEY = "session_token_expires";
export const OAUTH_REFRESH_KEY = "only_refresh";
export const REQUEST_DELAY = 2000;
export const DATE_FORMAT = "DD.MM.YYYY";
export const DATE_WITH_TIME_FORMAT = "DD.MM.YYYY HH:mm";
