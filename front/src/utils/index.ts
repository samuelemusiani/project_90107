export function rand(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

export function formatTime(date: Date): string {
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  return `${hours}:${minutes}`;
}
