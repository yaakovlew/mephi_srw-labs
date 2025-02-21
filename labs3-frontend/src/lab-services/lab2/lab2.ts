import { Lab2 } from 'src/models/lab/lab2';
import { getMatrixWeightVector } from './calculators';
import { roundToNDigits } from 'src/utils/round-to-n-digits';

export const getWeightsDifference = (variant: Lab2.Variant) => {
  const table1 = variant.tables[0];
  const table2 = variant.tables[1];

  const weights1 = getMatrixWeightVector(table1);
  const weights2 = getMatrixWeightVector(table2);

  return weights1.map((value, index) =>
    roundToNDigits(Math.abs(value - weights2[index]), 3)
  );
};

export const solveVariant = (variant: Lab2.Variant): Lab2.Answer => {
  const weightsDifference = getWeightsDifference(variant);

  const answer = weightsDifference.map((value, index) => ({
    value,
    index: index + 1,
  }));

  return answer.sort((a, b) => b.value - a.value);
};
