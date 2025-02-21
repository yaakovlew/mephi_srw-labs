export namespace Lab2 {
  export interface Variant {
    variant: string;
    tables: [Table, Table];
    description: [string, string];
  }

  export type Table = number[][];

  export type AnswerElement = {
    value: number;
    index: number;
  };

  export type Answer = AnswerElement[];

  export interface VariantInfo {
    variant: Variant;
    variantNumber: number;
  }

  export interface SaveVariant {
    number: number;
    data: Variant;
  }

  export interface Info {
    percentage: number;
    step: number;
    user_id: number;
    variance: SaveVariant;
  }
}
