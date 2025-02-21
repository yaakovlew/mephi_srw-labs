import { II18n } from 'src/services/resource/i18nTypes';

export type IDiscipline = IDiscipline.Discipline;

export namespace IDiscipline {
  export interface Discipline {
    name: string;
    discipline_id: string;
  }

  export interface DisciplineFullInfo {
    discipline_name: string;
    discipline_name_en: string;
    exam_mark: number;
    id: number;
    lesson_visiting_mark: number;
    seminar_visiting_mark: number;
  }

  export type DisciplineFetchData = II18n<IDiscipline[]>;

  export interface FetchData {
    id: string;
  }

  export interface AddGroup {
    discipline_id: number;
    group_id: number;
  }

  export interface GetGroups {
    id: number;
  }

  export interface CreateDiscipline {
    discipline_name: string;
    discipline_name_en: string;
    exam_mark: number;
    lesson_visiting_mark: number;
    seminar_visiting_mark: number;
  }

  export interface ChangeDisciplineName {
    discipline_id: number;
    name: string;
  }

  export interface ChangeDisciplineExamMark {
    discipline_id: number;
    lesson_mark: number;
  }

  export interface ChangeDisciplineLessonMark {
    discipline_id: number;
    lesson_mark: number;
  }

  export interface ChangeDisciplineSeminarMark {
    discipline_id: number;
    seminar_mark: number;
  }
}
