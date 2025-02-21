import {
  calculatePriorityVector,
  calculateWeightVector,
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
} from './lab-calculators';
import { Lab1a } from 'src/models/lab/lab1a';

export const solveMatrix = (
  matrix: number[][]
): Lab1a.ReturnTypeSolveMatrix => {
  const priorityVector = calculatePriorityVector(matrix);
  const priorityVectorSum = priorityVector.reduce(
    (acc, value) => acc + value,
    0
  );
  const weightVector = calculateWeightVector(priorityVector, priorityVectorSum);
  const matrixWeightVector = calculateMatrixWeightVector(matrix, weightVector);
  const lambdaVector = calculateLambdaVector(weightVector, matrixWeightVector);
  const eigenvalue = calculateEigenvalue(lambdaVector);
  const consistencyIndex = calculateConsistencyIndex(eigenvalue, lambdaVector);
  const consistencyRatio = calculateConsistencyRatio(matrix, consistencyIndex);

  if (!isConsistencyAcceptable(consistencyRatio)) {
    const weightsRatioMatrix = calculateWeightsRatioMatrix(weightVector);
    const weightsCorrectedRatioMatrix = calculateCorrectedWeightRatioMatrix(
      matrix,
      weightsRatioMatrix
    );
    const maxSumIndex = findMaxSumIndex(weightsCorrectedRatioMatrix);

    const newVariant = calculateNewCorrectedMatrix(
      matrix,
      weightsRatioMatrix,
      maxSumIndex
    );
    return solveMatrix(newVariant);
  }

  return {
    priorityVector,
    priorityVectorSum,
    weightVector,
    matrixWeightVector,
    lambdaVector,
    eigenvalue,
    consistencyIndex,
    consistencyRatio,
  };
};

export const getMatrixWeightVector = (matrix: number[][]) =>
  solveMatrix(matrix).weightVector;
