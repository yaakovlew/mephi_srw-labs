import { Lab1a } from 'src/models/lab/lab1a';
import { getMatrixElementsCount } from '../base/get-matrix-elements-count';
import { getPercent } from '../base/get-percent';
import { isNumbersSame } from '../base/is-numbers-same';

export const getArrayCorrectNumber = (
  checkedMatrix: number[],
  correctMatrix: number[]
) => {
  const checkedMatrixElementsCount = getMatrixElementsCount(checkedMatrix);
  const correctMatrixElementsCount = getMatrixElementsCount(correctMatrix);
  if (checkedMatrixElementsCount !== correctMatrixElementsCount)
    return {
      correct: 0,
      all: 0,
    };
  let correctElements = 0;

  checkedMatrix.forEach((value, firstIndex) => {
    if (isNumbersSame(value, correctMatrix[firstIndex])) correctElements++;
  });

  return { correct: correctElements, all: checkedMatrixElementsCount };
};

export const getMatrixCorrectPercent = (
  answerToCheck: Record<string, number[] | number>,
  correctAnswer: Record<string, number[] | number>
) => {
  let correctElemnts = 0;
  let allElements = 0;
  const toCheckEntries = Object.entries(answerToCheck);

  toCheckEntries.forEach(([key, value]) => {
    const keyOfReturnTypeSolveMatrix = key as keyof Record<
      string,
      number[] | number
    >;
    const correctArr = correctAnswer[keyOfReturnTypeSolveMatrix];
    if (Array.isArray(value) && Array.isArray(correctArr)) {
      const { correct, all } = getArrayCorrectNumber(value, correctArr);
      allElements = allElements + all;
      correctElemnts = correctElemnts + correct;
    } else if (typeof value === 'number' && typeof correctArr === 'number') {
      if (isNumbersSame(correctArr, value)) correctElemnts++;
      allElements++;
    }
  });

  return getPercent(correctElemnts, allElements);
};

export const getResultCorrectPercent = (
  answerToCheck: Lab1a.PrioritiesAlternatives,
  correctAnswer: Lab1a.VariantAnswer
) => {
  let correctElemnts = 0;
  let allElements = 0;

  const prioritiesToCheck = answerToCheck.alternatives;

  const { correct, all } = getArrayCorrectNumber(
    prioritiesToCheck,
    correctAnswer.resultPriorities
  );
  allElements = allElements + all;
  correctElemnts = correctElemnts + correct;

  if (
    isNumbersSame(
      prioritiesToCheck[answerToCheck.choosenAlternativeIndex],
      correctAnswer.resultPriorities[correctAnswer.bestAlternative]
    )
  )
    correctElemnts++;
  allElements++;

  return getPercent(correctElemnts, allElements);
};
