import { roundToNDigits } from './round-to-n-digints';

export const getPercent = (comparable: number, full: number) =>
  roundToNDigits(comparable / full);
