export type ISeminarian = ISeminarian.Seminarian;

export namespace ISeminarian {
  export interface Seminarian {
    email: string;
    name: string;
    seminarian_id: string;
    surname: string;
  }

  export interface FetchData {
    seminarians: ISeminarian[];
  }

  export interface GroupSeminarian {
    discipline_id: string;
    group_id: string;
    seminarian_id: string;
  }
}
