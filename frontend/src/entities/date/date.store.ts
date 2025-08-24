import { defineStore } from "pinia";
import type { IDateStore } from "./date.types";

export const useDateStore = defineStore("dates", {
  state: (): IDateStore => ({ date: null }),
});
