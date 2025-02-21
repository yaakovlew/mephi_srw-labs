export namespace ITest {
  export interface Test {
    default_mark: number;
    minutes_duration: number;
    name: string;
    task_description: string;
    test_id: number;
  }

  export interface CreateTest {
    default_mark: number;
    minutes_duration: number;
    name: string;
    task_description: string;
    name_en: string;
    task_description_en: string;
  }

  export interface GetTest {
    en: Test[];
    ru: Test[];
  }

  export interface AddTheme {
    count: number;
    test_id: number;
    theme_id: number;
  }

  export interface DeleteTheme {
    test_id: number;
    theme_id: number;
  }

  export interface OpenTest {
    date: number;
    test_id: number;
    user_id: number;
  }

  export interface CloseTest {
    test_id: number;
    user_id: number;
  }

  export interface GetOpenedTest {
    test_id: number;
    user_id: number;
  }

  export interface GetTestReport {
    test_id: number;
    user_id: number;
  }

  export interface ChangeTestDescription {
    description: string;
    test_id: number;
  }

  export interface ChangeTestDuration {
    minutes_duration: number;
    test_id: number;
  }

  export interface GetStudentMark {
    test_id: number;
    user_id: number;
  }

  export interface ChangeTestMark {
    mark: number;
    test_id: number;
    user_id: number;
  }

  export interface ChangeTestName {
    name: string;
    test_id: number;
  }

  export interface TestAnswer {
    answer: [];
    question_id: number;
    test_id: number;
  }

  export interface TestAnswers {
    answers: TestAnswer[];
  }

  export interface PassTest {
    answers: TestAnswers;
    test_id: number;
  }

  export interface StudentOpenTest {
    group_name: string;
    minutes_duration: number;
    name: string;
    start_date: number;
    student_id: number;
    surname: string;
  }

  export interface GetStudentsOpenTest {
    students: StudentOpenTest[];
  }
}
