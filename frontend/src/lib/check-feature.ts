import features from "../features.json";

export function isHasFeature(key: keyof typeof features) {
  return Boolean(features[key]);
}
