import type { IOauthProvider } from "@/entities/tech";

export const ENDPOINTS = {
  userInfo: `/api/v1/users`,
  auth: (provider: IOauthProvider) =>
    `/api/v1/oauth/${provider}?comebackUrl=${encodeURIComponent(window.location.href)}&frontend_host=${window.location.host}`,
  paymentTemplates: `/api/v1/payment_templates`,
  paymentTemplate: (id: number) => `/api/v1/payment_templates/${id}`,
  payments: `/api/v1/payments`,
  payment: (id: number) => `/api/v1/payments/${id}`,
  receiptTemplates: `/api/v1/receipt_templates`,
  receiptTemplate: (id: number) => `/api/v1/receipt_templates/${id}`,
  receipts: `/api/v1/receipts`,
  receipt: (id: number) => `/api/v1/receipts/${id}`,
};
