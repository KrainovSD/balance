import type { IOauthProvider } from "@/entities/tech";

export const ENDPOINTS = {
  userInfo: `/api/v1/users`,
  auth: (provider: IOauthProvider) =>
    `/api/v1/oauth/${provider}?comebackUrl=${encodeURIComponent(window.location.href)}&frontend_host=${window.location.host}`,
};
