export type IStudent = IStudent.Student;

export namespace IStudent {
  export interface Student {
    group_name: string;
    name: string;
    student_id: string;
    surname: string;
  }

  export interface FetchData {
    students: Student[];
  }

  export interface ChangeGroup {
    group_id: string;
    user_id: string;
  }

  export interface ChangePassword {
    user_id: string;
    new_password: string;
  }
}
