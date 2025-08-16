import type { ValueOf } from "@krainovsd/js-helpers";
import type { OAUTH_PROVIDERS } from "./tech.constants";

export type IOauthProvider = ValueOf<typeof OAUTH_PROVIDERS>;
