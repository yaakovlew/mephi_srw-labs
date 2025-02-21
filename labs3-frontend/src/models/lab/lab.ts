export namespace Lab {
  export interface FuncMark {
    name: string;
    func: string;
  }

  export interface Criteria {
    definition: string;
    extra_info: string;
    func: string;
    weight: number;
    func_mark: FuncMark[];
  }

  export type MatriciesHeaders = string[][];

  export interface Variant {
    number: number;
    task: string;
    criteria: Criteria[];
    alternative: Alternative[];
    rule: Rule[];
    importance_criteria: ImportanceCriteria[];
    importance_alternative: ImportanceCriteria[];
    matrix?: number[][];
    next_matrices_headers?: MatriciesHeaders;
  }

  export interface LabVariant {
    variant: Variant;
  }

  export interface CriteiaCount {
    count: string | number;
    value: string;
  }

  export interface Alternative {
    description: string;
    criteria_count: CriteiaCount[];
  }

  export interface Rule {
    name: string;
  }

  export interface ImportanceCriteria {
    importance: string;
    short_importance_name: string;
    points: Point[];
  }

  export interface Point {
    X: number;
    Y: number;
  }

  export interface GetVariant {
    user_id: number;
    variant: Variant;
  }

  export interface AlternativeSet {
    sets: number[][];
  }

  export interface NonDominated {
    set: number[];
  }

  export interface Result {
    set: number[];
    chosen_index: number;
  }

  export interface ResultAnswer {
    max_mark: number;
    percentage: number;
    result: number[][];
    index: number;
  }

  export interface AlternativeDiffMatrices {
    matrices: number[][];
    step: number;
  }

  export interface Intersection {
    matrix: number[][];
  }

  export interface Coff {
    matrix: number[][];
  }

  export interface CoffAnswer {
    max_mark: number;
    percentage: number;
    result: number[][];
  }

  export interface IntersectionAnswer {
    max_mark: number;
    percentage: number;
    result: number[][];
  }

  export interface AlternativeDiffMatricesAnswer {
    max_mark: number;
    percentage: number;
    result: number[][];
  }

  export interface NonDominatedAnswer {
    max_mark: number;
    percentage: number;
    result: number[];
  }

  export interface AlternativeSetAnswer {
    max_mark: number;
    percentage: number;
    result: number[][];
  }

  export interface Info {
    max_mark: number;
    percentage: number;
    step: number;
    user_id: number;
  }
}
