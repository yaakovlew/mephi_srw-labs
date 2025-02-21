import { roundToNDigits } from 'src/utils/round-to-n-digits';
import { mathExpectationsOfConsistencyIndex } from '../constants/mathExpectationsOfConsistencyIndex';

const calculatePriorityElement = (arr: number[]) =>
  roundToNDigits(
    Math.pow(
      arr.reduce((acc, value) => acc * value, 1),
      1 / arr.length
    )
  );

const calculatePriorityVector = (matrix: number[][]) =>
  matrix.map((arr) => calculatePriorityElement(arr));

const calculateWeightElement = (vectorElement: number, vectorSum: number) =>
  roundToNDigits(vectorElement / vectorSum);

const calculateWeightVector = (vector: number[], priorityVectorSum: number) =>
  vector.map((element) => calculateWeightElement(element, priorityVectorSum));

const calculateMatrixWeightElement = (
  matrixRow: number[],
  weightVector: number[]
) =>
  roundToNDigits(
    matrixRow.reduce(
      (acc, value, index) => roundToNDigits(acc + value * weightVector[index]),
      0
    )
  );

const calculateMatrixWeightVector = (
  matrix: number[][],
  weightVector: number[]
) =>
  matrix.map((matrixRow) =>
    calculateMatrixWeightElement(matrixRow, weightVector)
  );

const calculateLambdaVector = (
  weightVector: number[],
  matrixWeightVector: number[]
) =>
  matrixWeightVector.map((value, index) =>
    roundToNDigits(value / weightVector[index])
  );

const calculateEigenvalue = (lambdaVector: number[]) =>
  roundToNDigits(
    lambdaVector.reduce((acc, value) => acc + value, 0) / lambdaVector.length
  );

const calculateConsistencyIndex = (
  eigenvalue: number,
  lambdaVector: number[]
) =>
  roundToNDigits(
    (eigenvalue - lambdaVector.length) / (lambdaVector.length - 1)
  );

const calculateConsistencyRatio = (
  matrix: number[][],
  consistencyIndex: number
) =>
  roundToNDigits(
    consistencyIndex / mathExpectationsOfConsistencyIndex[matrix.length]
  );

const isConsistencyAcceptable = (consistencyRatio: number) =>
  Math.abs(consistencyRatio) < 0.1;

const calculateWeightsRatioMatrix = (weightVector: number[]) => {
  const matrix: number[][] = [];
  weightVector.forEach((weight) => {
    matrix.push(weightVector.map((value) => roundToNDigits(weight / value)));
  });

  return matrix;
};

const calculateCorrectedWeightRatioMatrix = (
  matrix: number[][],
  weightsRatioMatrix: number[][]
) => {
  const newMatrix = matrix.map((arr, index) =>
    arr.map((value, index2) =>
      Math.abs(roundToNDigits(value - weightsRatioMatrix[index][index2]))
    )
  );
  return newMatrix;
};

const findMaxSumIndex = (weightRatioMatrix: number[][]) => {
  const sums = weightRatioMatrix.map((arr) =>
    arr.reduce((acc, value) => acc + value, 0)
  );
  return sums.indexOf(Math.max(...sums));
};

const calculateNewCorrectedMatrix = (
  matrix: number[][],
  weightRatioMatrix: number[][],
  indexToChange: number
) =>
  matrix.map((arr, index) => {
    if (index === indexToChange) return weightRatioMatrix[index];
    return arr.map((value, index2) =>
      index2 === indexToChange
        ? roundToNDigits(1 / weightRatioMatrix[indexToChange][index])
        : value
    );
  });

const createWeightsMatrix = (...weights: number[][]) => {
  const matrix: number[][] = [];
  if (weights[0] && weights[0]?.length) {
    for (let i = 0; i < weights[0]?.length; i++) {
      matrix[i] = [];
      for (let j = 0; j < weights.length; j++) {
        matrix[i].push(weights[j][i]);
      }
    }
  }

  return matrix;
};

const createResultPriorities = (multipliedMatrices: number[][]) => {
  const res: number[] = [];

  multipliedMatrices?.forEach((arr) => {
    res.push(...arr);
  });

  return res;
};

const findBestAlternative = (priorities: number[]) => {
  const maxElement = Math.max(...priorities);
  return priorities.findIndex((pr) => pr === maxElement);
};

export {
  calculatePriorityElement,
  calculatePriorityVector,
  calculateWeightElement,
  calculateWeightVector,
  calculateMatrixWeightElement,
  calculateMatrixWeightVector,
  calculateLambdaVector,
  calculateEigenvalue,
  calculateConsistencyIndex,
  calculateConsistencyRatio,
  isConsistencyAcceptable,
  calculateWeightsRatioMatrix,
  calculateCorrectedWeightRatioMatrix,
  findMaxSumIndex,
  calculateNewCorrectedMatrix,
  createWeightsMatrix,
  createResultPriorities,
  findBestAlternative,
};
