import { importanceValues } from 'src/mock/lab1b';
import { Lab1b } from 'src/models/lab/lab1b';
import { getRandomInt } from '../base/get-random-number';
import { mirrorMatrix } from '../base/transpose-matrix';

export const fillMatrix = (matrix: number[][]) => {
  const importanceValuesToFill = importanceValues.slice(
    0,
    importanceValues.length - 1
  );

  const filledMatrix: number[][] = matrix.map((row, index) =>
    row.map((_, i) => {
      if (i === index) return 1;
      if (i > index) {
        const values =
          importanceValuesToFill[
            getRandomInt(0, importanceValuesToFill.length - 1)
          ].value;
        return values[getRandomInt(0, values.length - 1)];
      }
      return 1 / matrix[i][index];
    })
  );

  return mirrorMatrix(filledMatrix);
};
