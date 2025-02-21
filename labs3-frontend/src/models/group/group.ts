import { ISeminarian } from '../seminarian/seminarian';

export type IGroup = IGroup.Group;

export namespace IGroup {
  export interface Group {
    group_id: string;
    is_archive: boolean;
    name: string;
  }

  export interface GroupsFetchData {
    groups: IGroup.Group[];
  }

  export interface AddGroup {
    group_name: string;
  }

  export interface ChangeGroupName {
    name: string;
    group_id: number;
  }

  export interface DeleteGroup {
    id: string;
  }

  export interface FetchGroupStudents {
    id: string;
  }

  export interface GroupStudent {
    name: string;
    surname: string;
    user_id: string;
  }

  export interface GetGroupStudentStudentId {
    students: GroupStudentStudentId[];
  }

  export interface GroupStudentStudentId {
    name: string;
    surname: string;
    student_id: string;
  }

  export interface FetchGroupSeminarians {
    discipline_id: number;
    group_id: number;
  }

  export interface GroupSeminarians {
    seminarians: ISeminarian.Seminarian[];
  }

  export interface GroupStudents {
    students: GroupStudent[];
  }

  export interface AddSeminarian {
    discipline_id: number;
    group_id: number;
    seminarian_id: number;
  }
}
