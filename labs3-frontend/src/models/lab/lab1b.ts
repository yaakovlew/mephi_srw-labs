export namespace Lab1b {
  export interface ImpoortanceValues {
    value: number[];
    definition: string;
  }

  export interface Variant {
    theme: string;
    lpr: string;
    criterias: Record<string, Criteria[]>;
    alternatives: string[];
  }

  export interface Criteria {
    title: string;
    type: 'quantitative' | 'qualitative';
    moreBetter?: boolean;
  }

  export type Varinats = Record<number, Variant>;
}
