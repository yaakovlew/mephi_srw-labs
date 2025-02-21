export namespace IMark {
  export interface Mark {
    mark: number;
    user_id: number;
    user_name: string;
    user_surname: string;
  }

  export interface GetMarkExam {
    group_id: number;
    discipline_id: number;
  }

  export interface Marks {
    marks: Mark[];
  }

  export interface TestMark {
    mark: number;
    test_id: number;
    user_id: number;
  }

  export interface PostMarkExam {
    mark: number;
    discipline_id: number;
    user_id: number;
  }

  export interface GetMarkLabaratory {
    group_id: number;
    labaratory_id: number;
  }

  export interface PostMarkLabaratory {
    mark: number;
    labaratory_id: number;
    user_id: number;
  }

  export interface GetMarkTest {
    group_id: number;
    test_id: number;
  }

  export interface PostMarkTest {
    mark: number;
    test_id: number;
    user_id: number;
  }
}
