export const roundToNDigits = (value: number, digits = 2) => {
  return Math.round((value + Number.EPSILON) * 10 ** digits) / 10 ** digits;
};
