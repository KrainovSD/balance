export type IPaymentTemplate = {
  id: number;
  name: string;
  amount: number;
};

export type IPayment = {
  id: number;
  paymentId: number;
  amount: number;
  name: string;
  date: string;
  description: string;
};

export type IPaymentsStore = {
  paymentTemplates: IPaymentTemplate[];
  payments: IPayment[];
  getPaymentTemplatesLoading: boolean;
  createPaymentTemplatesLoading: boolean;
  deletePaymentTemplatesLoading: boolean;
  updatePaymentTemplatesLoading: boolean;
  getPaymentsLoading: boolean;
  createPaymentsLoading: boolean;
  deletePaymentsLoading: boolean;
  updatePaymentsLoading: boolean;
};
