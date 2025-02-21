import { mirrorMatrix } from '../base/transpose-matrix';
import { roundToNDigits } from '../../utils/round-to-n-digits';
import {
  calculateMode,
  getAverage,
  getGeoMean,
  getHarMean,
  getDispersion,
  getStandardDeviation,
  getReverseDispersion,
  getBayesMean,
  getDeviation,
} from './math';

export const getMatrixMode = (matrix: number[][]): number[] => {
  const mirroredMatrix = mirrorMatrix(matrix);

  const mode: number[] = [];

  mirroredMatrix.forEach((row) => {
    mode.push(calculateMode(row));
  });

  return mode;
};

export const getMatrixAverage = (matrix: number[][]): number[] => {
  const mirroredMatrix = mirrorMatrix(matrix);

  const average: number[] = [];

  mirroredMatrix.forEach((row) => {
    average.push(getAverage(row));
  });

  return average;
};

export const getMatrixGeoMean = (matrix: number[][]): number[] => {
  const mirroredMatrix = mirrorMatrix(matrix);

  const geoMean: number[] = [];

  mirroredMatrix.forEach((row) => {
    geoMean.push(getGeoMean(row));
  });

  return geoMean;
};

export const getMatrixHarMean = (matrix: number[][]): number[] => {
  const mirroredMatrix = mirrorMatrix(matrix);

  const harMean: number[] = [];

  mirroredMatrix.forEach((row) => {
    harMean.push(getHarMean(row));
  });

  return harMean;
};

export const getMatrixDispersion = (matrix: number[][]): number[] => {
  const average = getMatrixAverage(matrix);

  const dispersion: number[] = [];

  matrix.forEach((row) => {
    dispersion.push(getDispersion(row, average));
  });

  return dispersion;
};

export const getMatrixStandardDeviation = (matrix: number[][]): number[] => {
  const average = getMatrixAverage(matrix);

  const standartDevination: number[] = [];

  matrix.forEach((row) => {
    standartDevination.push(getStandardDeviation(row, average));
  });

  return standartDevination;
};

export const getMatrixReverseDispersion = (matrix: number[][]): number[] => {
  const average = getMatrixAverage(matrix);

  const reverseDispersion: number[] = [];

  matrix.forEach((row) => {
    reverseDispersion.push(getReverseDispersion(row, average));
  });

  return reverseDispersion;
};

export const getMatrixWeightVector = (matrix: number[][]): number[] => {
  const reverseDispersion = getMatrixReverseDispersion(matrix);
  const reverseDispersionSum = reverseDispersion.reduce((a, b) => a + b, 0);
  return reverseDispersion.map((value) =>
    roundToNDigits(value / reverseDispersionSum, 3)
  );
};

export const getMatrixBayesMean = (matrix: number[][]): number[] => {
  const weightVector = getMatrixWeightVector(matrix);
  const mirroredMatrix = mirrorMatrix(matrix);

  const bayesMean: number[] = [];

  mirroredMatrix.forEach((row) => {
    bayesMean.push(getBayesMean(row, weightVector));
  });

  return bayesMean;
};

export const getMatrixAverageDeviation = (matrix: number[][]): number[] => {
  const average = getMatrixAverage(matrix);
  const bayes = getMatrixBayesMean(matrix);

  const averageDeviation: number[] = [];

  average.forEach((value, i) => {
    averageDeviation.push(getDeviation(value, bayes[i]));
  });

  return averageDeviation;
};

export const getMatrixGeoDeviation = (matrix: number[][]): number[] => {
  const geo = getMatrixGeoMean(matrix);
  const bayes = getMatrixBayesMean(matrix);

  const geoDeviation: number[] = [];

  geo.forEach((value, i) => {
    geoDeviation.push(getDeviation(value, bayes[i]));
  });

  return geoDeviation;
};

export const getMatrixHarDeviation = (matrix: number[][]): number[] => {
  const har = getMatrixHarMean(matrix);
  const bayes = getMatrixBayesMean(matrix);

  const harDeviation: number[] = [];

  har.forEach((value, i) => {
    harDeviation.push(getDeviation(value, bayes[i]));
  });

  return harDeviation;
};
