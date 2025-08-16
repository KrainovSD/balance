export type ApiErrorInterface<T> = {
  title: string;
  status: number;
  info?: T;
};
