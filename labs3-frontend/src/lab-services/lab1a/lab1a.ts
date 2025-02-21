import { Lab1a } from 'src/models/lab/lab1a';
import { multiplyMatrices } from '../base/multiply-matrices';
import { solveMatrix } from '../custom/solve-matrix';
import {
  createWeightsMatrix,
  createResultPriorities,
  findBestAlternative,
} from '../custom/lab-calculators';

export const getFinalMatrix = (variant: Lab1a.Variant) => {
  const { matrix_1, matrix_2, matrix_3, matrix_4, matrix_5 } = variant;
  const firstWeights = solveMatrix(matrix_1).weightVector;
  const secondWeights = solveMatrix(matrix_2).weightVector;
  const thirdWeights = solveMatrix(matrix_3).weightVector;
  const fourthWeights = solveMatrix(matrix_4).weightVector;
  const fifthWeights = solveMatrix(matrix_5).weightVector;

  const finalMatrix = createWeightsMatrix(
    secondWeights,
    thirdWeights,
    fourthWeights,
    fifthWeights
  );

  const finalWeights = createWeightsMatrix(firstWeights);

  return {
    finalMatrix,
    finalWeights,
  };
};

export const solveVariant = (
  variant: Lab1a.Variant
): Lab1a.VariantAnswer | null => {
  const { finalMatrix, finalWeights } = getFinalMatrix(variant);

  const multipliedMatrices = multiplyMatrices(finalMatrix, finalWeights);

  if (!multipliedMatrices) return null;

  const resultPriorities = createResultPriorities(multipliedMatrices);

  const bestAlternative = findBestAlternative(resultPriorities);

  return {
    bestAlternative,
    multipliedMatrices,
    resultPriorities,
  };
};

export const getMaxMarksForStep = () => {
  const markForStep = 15;
  const marks = Array<number>(6).fill(0);

  const res: number[] = marks.map(() => markForStep);

  res[res.length - 1] = 25;

  return res;
};
