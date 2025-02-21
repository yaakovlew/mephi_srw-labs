export type IQuestion = IQuestion.Question;

export namespace IQuestion {
  export interface Question {
    is_variable: 0 | 1 | 2;
    question: string;
    question_id: number;
  }

  export interface GetQuestions {
    en: Question[];
    ru: Question[];
  }

  export interface GetQuestionsWithotTheme {
    questions: Question[];
  }

  export interface ChangeQuestion {
    name: string;
    question_id: number;
  }

  export interface AddQuestionToTheme {
    question_id: number;
    theme_id: number;
  }

  export interface RemoveQuestionFromTheme {
    question_id: number;
    theme_id: number;
  }

  export interface Answer {
    is_right: boolean;
    answer_id: number;
    answer: string;
  }

  export interface GetAnswers {
    en: Answer[];
    ru: Answer[];
  }

  export interface AddAnswer {
    is_right: boolean;
    name: string;
    name_en: string;
    question_id: number;
  }

  export interface ChangeAnswerName {
    answer_id: number;
    name: string;
  }

  export interface ChangeIsAnswerRight {
    answer_id: number;
    is_right: boolean;
  }

  export interface CreateQuestion {
    is_variable: 0 | 1 | 2;
    question: string;
    question_en: string;
  }

  export interface AnswersStudent {
    answer_id: number;
    name: string;
  }

  export interface SingleQuestion {
    answers: AnswersStudent[];
    questions: Question;
    theme_id: number;
  }

  export type StudentTest = SingleQuestion[];

  export interface GetTest {
    en: StudentTest;
    ru: StudentTest;
  }
}
