export type ILesson = IAttendance.Lesson;
export type ISeminar = IAttendance.Seminar;

export namespace IAttendance {
  export interface Lesson {
    lesson_id: string;
    name: string;
  }

  export interface LessonDate {
    date: number;
    group_id: number;
    lesson_id: string;
    lesson_name: string;
  }

  export interface AddLessonDate {
    date: number;
    group_id: number;
    lesson_id: number;
  }

  export interface GetLessonDate {
    group_id: number;
    lesson_id: number;
  }

  export interface ChangeLessonDate {
    date: number;
    group_id: number;
    lesson_id: number;
  }

  export interface DeleteLessonDate {
    group_id: number;
    lesson_id: number;
  }

  export interface GetLessons {
    lessons: Lesson[];
  }

  export interface GetLessonsDate {
    lessons: LessonDate[];
  }

  export interface Seminar {
    date: number;
    seminar_id: string;
    name: string;
  }

  export interface GetSeminars {
    seminars: Seminar[];
  }

  export interface ChangeSeminarName {
    name: string;
    seminar_id: number;
  }

  export interface ChangeSeminarDate {
    date: number;
    seminar_id: number;
  }

  export interface CreateSeminar {
    date: number;
    discipline_id: number;
    group_id: number;
    name: string;
  }

  export interface CreateLesson {
    discipline_id: number;
    name: string;
  }

  export interface FetchData {
    discipline_id: string;
  }

  export interface FecthSeminars {
    group_id: number;
    discipline_id: number;
  }

  export interface NewLesson {
    discipline_id: string;
    name: string;
  }

  export interface ChangeLessonName {
    lesson_id: string;
    name: string;
  }
}
