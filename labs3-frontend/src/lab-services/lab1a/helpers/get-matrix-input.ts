import { Lab1a } from 'src/models/lab/lab1a';

export const getMatrixInputs = (
  matrix: number[][],
  defaultValues?: Lab1a.MatrixAnswer
): Lab1a.MatrixAnswer => {
  if (defaultValues) {
    return defaultValues;
  }
  const priorityVector = matrix.map(() => 0);
  const weightVector = matrix.map(() => 0);
  const matrixWeightVector = matrix.map(() => 0);
  const lambdaVector = matrix.map(() => 0);
  const eigenvalue = 0;
  const consistencyIndex = 0;
  const consistencyRatio = 0;

  const res = {
    priorityVector,
    weightVector,
    matrixWeightVector,
    lambdaVector,
    eigenvalue,
    consistencyIndex,
    consistencyRatio,
  };

  return res;
};
