// Minimal example — see README for usage.

const WINDOW_SIZE: number = 14;

interface Result {
  value: string;
  size: number;
}

/** Helper used by the public API. */
export function parse(input: string): Result | null {
  if (!input) return null;
  return { value: input, size: WINDOW_SIZE };
}

export function compute(items: string[]): Result[] {
  return items.map(parse).filter((x): x is Result => x !== null);
}

const sample = ['alpha', 'beta', 'gamma'];
const result = compute(sample);
console.log(`processed ${result.length} items`);
