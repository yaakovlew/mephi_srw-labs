export namespace IVisiting {
  export interface StudentVisiting {
    is_absent: boolean;
    student_id: string;
    student_name: string;
  }

  export interface GetLessonsVisiting {
    students: StudentVisiting[];
  }

  export interface GetSeminarsVisiting {
    students: StudentVisiting[];
  }

  export interface FetchStudentVisiting {
    is_absent: boolean;
    lesson_id: number;
    user_id: number;
  }

  export interface FetchStudentSeminarVisiting {
    is_absent: boolean;
    seminar_id: number;
    user_id: number;
  }

  export interface FetchData {
    group_id: number;
    lesson_id: number;
  }

  export interface GetSeminarVisiting {
    seminar_id: string;
  }
}
