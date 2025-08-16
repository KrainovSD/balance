export type IUser = {
  id: number;
  name: string;
  username: string;
  email: string;
};

export type IUsersStore = {
  userInfo: IUser | null | undefined;
  getUserInfoLoading: boolean;
};
