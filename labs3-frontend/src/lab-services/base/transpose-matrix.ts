export const mirrorMatrix = (matrix: number[][]): number[][] => {
  const numRows = matrix.length;
  const numCols = matrix[0].length;

  const mirroredMatrix = new Array(numCols);
  for (let i = 0; i < numCols; i++) {
    mirroredMatrix[i] = new Array(numRows);
  }

  for (let i = 0; i < numRows; i++) {
    for (let j = 0; j < numCols; j++) {
      mirroredMatrix[j][i] = matrix[i][j];
    }
  }

  return mirroredMatrix;
};
