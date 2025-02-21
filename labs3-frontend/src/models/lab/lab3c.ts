export namespace Lab3c {
  export interface AlternativeMatrix {
    matrix: number[][];
  }

  export interface AlternativeMatrixResult {
    percentage: number;
    result: number[][];
  }

  export interface CriteriaMatirx {
    set: number[];
  }

  export interface CriteriaMatirxResult {
    percentage: number;
    result: number[][];
  }

  export interface MatriceXY {
    X: number;
    Y: number;
  }

  export type MatriceXYRow = MatriceXY[];

  export type MatriceXYRowArray = MatriceXYRow[];

  export type AlternativeMatrices = {
    matrices: MatriceXYRowArray[]
  };

  export type AlternativeMatricesResult = {
    percentage: number;
    result: AlternativeMatrices;
  };

  export type CriteriaEstimation = MatriceXYRowArray[];

  export type CriteriaEstimationResult = {
    percentage: number;
    result: CriteriaEstimation;
  };

  export interface Area {
    set: number[];
  }

  export interface AreaResult {
    percentage: number;
    result: number[][];
  }

  export interface Line {
    k: number;
    b: number;
  }

  export interface LineParameters {
    parameters: Line[];
  }

  export interface LineResult {
    percentage: number;
    result: LineParameters;
  }

  export interface KvLine {
    a1: number;
    a2: number;
    a3: number;
  }

  export interface KvLineParameters {
    parameters: KvLine[];
  }

  export interface KvLineResult {
    percentage: number;
    result: KvLineParameters;
  }

  export interface Result {
    set: number[];
    index: number;
  }

  export interface ResultResult {
    percentage: number;
    result: Result;
  }
}
