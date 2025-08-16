/* eslint-disable no-console */
import { isObject } from "@krainovsd/js-helpers";
import type { WatchEffectOptions } from "vue";

type DebugEffect = {
  track?: boolean;
  trigger?: boolean;
};

export function debugEffect(
  opts: DebugEffect = { track: true, trigger: true },
): WatchEffectOptions {
  return {
    onTrack: opts.track
      ? (event) => {
          console.log("track", {
            type: event.type,
            key: event.key,
            target:
              isObject(event.target) && "_value" in event.target
                ? event.target._value
                : event.target,
          });
        }
      : undefined,
    onTrigger: opts.trigger
      ? (event) => {
          console.log("trigger", {
            type: event.type,
            key: event.key,
            target:
              isObject(event.target) && "_value" in event.target
                ? event.target._value
                : event.target,
            value: event.newValue,
            oldValue: event.oldValue,
          });
        }
      : undefined,
  };
}
