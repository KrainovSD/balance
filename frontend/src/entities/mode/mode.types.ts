import type { ValueOf } from "@krainovsd/js-helpers";
import type { MODES } from "./mode.constants";

export type IMode = ValueOf<typeof MODES>;

export type IModeStore = {
  mode: IMode;
};
