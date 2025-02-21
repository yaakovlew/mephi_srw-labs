export namespace Lab3b {
  export interface RuleNumber {
    matrix: number[][];
  }

  export interface MatriceXY {
    X: number;
    Y: number;
  }

  export type MatriceXYRow = MatriceXY[];

  export type MatriceXYRowArray = MatriceXYRow[];

  export type AllMatrices = MatriceXYRowArray[];

  export type sendAllMatrices = {
    matrices: AllMatrices;
  };

  export type sendAllMatricesIntersection = {
    matrix: MatriceXYRowArray;
  };

  export type sendAllMatricesIntersectionResult = {
    result: MatriceXYRowArray;
    percentage: number;
  };

  export interface AllMatricesResult {
    result: AllMatrices;
    percentage: number;
  }

  export interface LevelSet {
    set: number[];
    delta: number;
    powerful: number;
  }

  export interface SendLevelSet {
    answer_level_set: LevelSet[][];
  }

  export interface SendLevelSetResult {
    result: SendLevelSet;
    percentage: number;
  }

  export interface MatrixDataCell {
    data: number;
    flag: boolean;
  }

  export interface AlternativeSetAnswer {
    max_mark: number;
    percentage: number;
    result: Record<string, MatrixDataCell[]>;
  }
}
