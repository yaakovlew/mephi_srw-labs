export type IUser = IUser.UserType & {
  id: string;
};

export namespace IUser {
  export interface Common {
    email: string;
    name: string;
    surname: string;
    group_name?: string;
    post: UserType;
  }

  export type UserType = 'student' | 'lecturer' | 'seminarian';

  export type Register = Common & {
    password: string;
  };

  export type Login = Pick<Register, 'email' | 'password'>;

  //   export type Update = Omit<Union, 'email'>;

  export function isStudent(user: UserType): user is UserType {
    return user === 'student';
  }

  export interface ForgetPassword {
    email: string;
  }

  export interface RestorePassword {
    new_password: string;
    token: string;
  }
}
