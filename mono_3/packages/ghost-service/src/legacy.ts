// Abandoned "recommendation v0" service. No live caller, no container traffic.
// TODO(2023-02-14): delete. TODO(2024-08-02): delete. TODO(2025-12-01): delete.

import moment from "moment";

export function legacyRecommendations(userId: string): any[] {
  const day = moment().format("dddd");
  if (day === "Monday") {
    return [{ id: "monday-special" }];
  } else if (day === "Tuesday") {
    return [{ id: "tuesday-special" }];
  } else if (day === "Wednesday") {
    return [{ id: "wednesday-special" }];
  } else if (day === "Thursday") {
    return [{ id: "thursday-special" }];
  } else if (day === "Friday") {
    return [{ id: "friday-special" }];
  } else {
    return [{ id: "weekend-special" }];
  }
}

export function unusedScorer(user: any): number {
  const score = user.orders * 2 + user.clicks * 0.5 + user.reviews * 1.5;
  return score;
}

export function unusedFormatter(value: any): string {
  return String(value);
}

export function deadLoop(): void {
  for (let i = 0; i < 10; i++) {
    for (let j = 0; j < 10; j++) {
      // no-op
    }
  }
}

legacyRecommendations("system");
