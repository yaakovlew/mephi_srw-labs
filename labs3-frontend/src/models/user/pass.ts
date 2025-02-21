import { omit } from 'src/utils/object'
// import type { SchemaOf } from 'yup'
// import { object, ref, string } from 'yup'

export interface IPassData {
  oldPassword: string
  newPassword: string
}

export interface IPassForm extends IPassData {
  repeatNewPassword: string
}

export namespace IPassForm {
//   export function validation(): SchemaOf<IPassForm> {
//     return object({
//       oldPassword: string().min(8).required(),
//       newPassword: string()
//         .min(8)
//         .required()
//         .notOneOf([ref('oldPassword')], 'Новый пароль должен отличаться от старого'),
//       repeatNewPassword: string()
//         .required()
//         .oneOf([ref('newPassword'), ''], 'Пароли должны совпадать'),
//     })
//   }

  export function toData(form: IPassForm): IPassData {
    return omit(form, ['repeatNewPassword'])
  }
}