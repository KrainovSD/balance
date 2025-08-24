export type IReceiptTemplate = {
  id: number;
  name: string;
  amount: number;
};

export type IReceipt = {
  id: number;
  receiptId: number;
  amount: number;
  name: string;
  date: string;
  description: string;
};

export type IReceiptsStore = {
  receiptTemplates: IReceiptTemplate[];
  receipts: IReceipt[];
  getReceiptTemplatesLoading: boolean;
  createReceiptTemplatesLoading: boolean;
  deleteReceiptTemplatesLoading: boolean;
  updateReceiptTemplatesLoading: boolean;
  getReceiptsLoading: boolean;
  createReceiptsLoading: boolean;
  deleteReceiptsLoading: boolean;
  updateReceiptsLoading: boolean;
};
