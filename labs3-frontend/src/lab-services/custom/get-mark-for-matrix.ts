export const getMarkForStep = (percentage: number, maxMark: number) =>
  Math.ceil(percentage * maxMark);
