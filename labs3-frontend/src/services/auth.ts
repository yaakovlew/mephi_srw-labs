import { $apiPublic, $apiStudent } from 'boot/axios';
import { TokenService } from 'src/utils/tokenService';
import type { IResource } from 'src/services/resource/resource';
import { handleAxiosError } from 'src/utils/service/errorHandler';
import { useServiceAction } from 'src/utils/service/action';
import { errorNotify } from 'src/utils/service/notify';
import type { IUser } from 'src/models/user/user';
import type { IPassData } from 'src/models/user/pass';

export const AuthService = {
  async login(data: IUser.Login): Promise<IResource<IUser.UserType>> {
    try {
      const res = await $apiPublic.post<ILoginResponse>('/auth/sign-in', {
        ...data,
      });
      TokenService.token = res.data.token;
      TokenService.type = res.data.post;
      return { data: res.data.post };
    } catch (err) {
      return handleAxiosError(err, errorNotify);
    }
  },

  register: useServiceAction((data: IUser.Register) =>
    $apiPublic.post<IUser>('auth/sign-up', data)
  ),

  async refresh(): Promise<IResource> {
    try {
      const res = await $apiPublic.post<ILoginResponse>('/auth/refresh');
      TokenService.token = res.data.token;
      return { data: null };
    } catch (err) {
      return handleAxiosError(err);
    }
  },

  async logout(): Promise<IResource> {
    try {
      await $apiPublic.post<never>('/auth/logout');
      return { data: null };
    } catch (err) {
      return handleAxiosError(err);
    } finally {
      TokenService.token = null;
    }
  },

  updatePassword: useServiceAction((data: IPassData) =>
    $apiStudent.post<unknown>('/me/password/new', data)
  ),

  forgetPassword: useServiceAction((data: IUser.ForgetPassword) =>
    $apiPublic.post('/forget-password', data)
  ),

  restorePassword: useServiceAction((data: IUser.RestorePassword) =>
    $apiPublic.post(
      '/restore-password',
      { new_password: data.new_password },
      {
        headers: {
          Authorization: 'Bearer ' + data.token,
        },
      }
    )
  ),
};

export interface ILoginResponse {
  token: string;
  post: IUser.UserType;
}
