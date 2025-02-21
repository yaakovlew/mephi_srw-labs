import type { AxiosError } from 'axios';

export interface IResourceError {
  statusCode?: number;
  message: string;
  original: AxiosError<IAxiosError>;
}

export namespace IResourceError {
  export function build(err: AxiosError<IAxiosError>): IResourceError {
    return {
      statusCode: err.response?.status,
      message:
        (err.response?.data as IApiError)?.errorDescription?.ru ||
        'Неизвестная ошибка!',
      original: err,
    };
  }
}

type IApiErrorLocales = 'ru' | 'en';
export interface IApiError {
  message: string;
  status: string;
  timestamp: string;
  errorDescription: Record<IApiErrorLocales, string>;
}

export interface IServerError {
  error: string;
  path: string;
  status: number;
  timestamp: string;
}

export type IAxiosError = IApiError | IServerError;
