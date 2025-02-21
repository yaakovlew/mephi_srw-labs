import { roundToNDigits } from 'src/utils/round-to-n-digits';

export const calculateMode = (data: number[]) => {
  const map = data.map((a) => data.filter((b) => a === b).length);
  return data[map.indexOf(Math.max(...map))];
};

export const getAverage = (data: number[]) =>
  roundToNDigits(data.reduce((a, b) => a + b, 0) / data.length, 3);

export const getGeoMean = (data: number[]) =>
  roundToNDigits(data.reduce((a, b) => a * b, 1) ** (1 / data.length), 3);

export const getHarMean = (data: number[]) => {
  return roundToNDigits(data.length / data.reduce((a, b) => a + 1 / b, 0), 3);
};

export const getDispersion = (data: number[], average: number[]) => {
  return roundToNDigits(
    data.reduce((a, b, i) => a + (b - average[i]) ** 2, 0) / data.length,
    3
  );
};

export const getStandardDeviation = (data: number[], average: number[]) =>
  roundToNDigits(Math.sqrt(getDispersion(data, average)), 3);

export const getReverseDispersion = (data: number[], average: number[]) =>
  roundToNDigits(1 / getDispersion(data, average), 3);

export const getBayesMean = (data: number[], weightVector: number[]) =>
  roundToNDigits(
    data.reduce((a, b, i) => a + b * weightVector[i], 0),
    3
  );

export const getDeviation = (value: number, bayesValue: number) =>
  roundToNDigits((Math.abs(value - bayesValue) / bayesValue) * 100, 1);
