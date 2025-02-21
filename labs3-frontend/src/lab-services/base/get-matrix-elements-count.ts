export const getMatrixElementsCount = (matrix: number[][] | number[]) =>
  Array.isArray(matrix[0])
    ? (matrix as number[][]).reduce((acc, arr) => acc + arr.length, 0)
    : matrix.length;
