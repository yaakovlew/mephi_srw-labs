export namespace Lab1a {
  export interface Variant {
    matrix_1: number[][];
    matrix_2: number[][];
    matrix_3: number[][];
    matrix_4: number[][];
    matrix_5: number[][];
  }

  export interface SaveVariant {
    number: number;
    data: Variant;
  }

  export interface VariantInfo {
    variant: Variant;
    variantNumber: number;
  }

  export interface Info {
    percentage: number;
    step: number;
    user_id: number;
    variance: SaveVariant;
  }

  export type UpdateInfo = Pick<Info, 'percentage' | 'step'>;

  export type MatrixAnswer = {
    priorityVector: number[];
    weightVector: number[];
    matrixWeightVector: number[];
    lambdaVector: number[];
    eigenvalue: number;
    consistencyIndex: number;
    consistencyRatio: number;
  };

  export type ReturnTypeSolveMatrix = MatrixAnswer & {
    priorityVectorSum: number;
  };

  export interface PrioritiesAlternatives {
    alternatives: number[];
    choosenAlternativeIndex: number;
  }

  export interface VariantAnswer {
    bestAlternative: number;
    multipliedMatrices: number[][];
    resultPriorities: number[];
  }

  export interface Result {
    percentage: number;
  }
}
