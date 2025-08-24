export function getMonthName(index: number, month: "long" | "short" = "short") {
  const date = new Date();
  date.setMonth(index);

  return date.toLocaleString("ru-RU", { month });
}
