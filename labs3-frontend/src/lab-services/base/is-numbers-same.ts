export const isNumbersSame = (a: number, b: number, tolerance = 0.03) =>
  Math.abs(a - b) < tolerance;
