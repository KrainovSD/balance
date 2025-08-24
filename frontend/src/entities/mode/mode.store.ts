import { defineStore } from "pinia";
import { MODES } from "./mode.constants";
import type { IModeStore } from "./mode.types";

export const useModeStore = defineStore("mode", {
  state: (): IModeStore => ({ mode: MODES.Payment }),
});
