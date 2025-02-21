export type IPersonalInfo = IPersonalInfo.Student;

export namespace IPersonalInfo {
  export interface Student {
    email: string;
    name: string;
    surname: string;
    group_name: string;
  }

  export interface IChangeName {
    name: string;
  }
  export interface IChangeSurname {
    surname: string;
  }

  export interface IChangePassword {
    new_password: string;
    old_password: string;
  }
}
