import { ResponseError } from "@krainovsd/js-helpers";
import type { ApiErrorInterface } from "../types";

export async function apiErrorLayer<ErrorT = unknown, T = unknown>(
  request: () => Promise<T>,
  onError?: (error: ApiErrorInterface<ErrorT>) => void,
): Promise<T | undefined> {
  try {
    const result = await request();

    return result;
  } catch (error) {
    if (error instanceof Error && error.name === "AbortError") return;

    if (error instanceof ResponseError) {
      onError?.({
        title: error.message,
        status: error.status,
        info: error.description as ErrorT,
      });
    } else if (error instanceof Error) {
      onError?.({
        title: error.message,
        status: 500,
        info: null as ErrorT,
      });
    } else {
      onError?.({
        title: "Unknown Error",
        status: 500,
        info: null as ErrorT,
      });
    }
  }
}
