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
