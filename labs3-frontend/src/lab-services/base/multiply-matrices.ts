import { roundToNDigits } from './round-to-n-digints';

export const multiplyMatrices = (matrix1: number[][], matrix2: number[][]) => {
  if (matrix1[0]?.length !== matrix2.length) {
    throw Error(
      'Умножение невозможно. Количество столбцов первой матрицы не равно количеству строк второй матрицы.'
    );
  }

  const result = [];
  for (let i = 0; i < matrix1.length; i++) {
    let sum = 0;
    if (matrix1[i] && matrix1[i].length && matrix2[i]?.length) {
      for (let k = 0; k < matrix1[i].length; k++) {
        sum += roundToNDigits(matrix1[i][k] * matrix2[k][0]);
      }
    }
    result.push([roundToNDigits(sum)]);
  }
  return result;
};
