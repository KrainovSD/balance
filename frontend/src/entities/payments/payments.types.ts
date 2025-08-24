export type IPaymentTemplate = {
  id: number;
  name: string;
  amount: number;
};

export type IPaymentsStore = {
  paymentTemplates: IPaymentTemplate[];
  getPaymentTemplateLoading: boolean;
};
