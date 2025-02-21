import type { IResource } from 'src/services/resource/resource';
import type { AxiosResponse } from 'axios';
import type { IResourceErrorAction } from 'src/utils/service/errorHandler';
import { handleAxiosError } from 'src/utils/service/errorHandler';
import { errorNotify } from './notify';

type IResponce<T> = Promise<AxiosResponse<T>>;
export type IFutureResource<T> = Promise<IResource<T>>;

type ISimpleParams<T> = () => IResponce<T>;
type IParams<T, P> = (params: P) => IResponce<T>;
type ISimpleReturn<T> = () => IFutureResource<T>;
type IReturn<T, P> = (params: P) => IFutureResource<T>;

export function useServiceAction<T>(
  query: ISimpleParams<T>,
  errorCallback?: IResourceErrorAction | null
): ISimpleReturn<T>;
export function useServiceAction<T, P>(
  query: IParams<T, P>,
  errorCallback?: IResourceErrorAction | null
): IReturn<T, P>;
export function useServiceAction<T, P>(
  query: IParams<T, P>,
  errorCallback: IResourceErrorAction | null = errorNotify
) {
  return async (params: P): IFutureResource<T> => {
    try {
      const res = await query(params);
      return { data: res.data };
    } catch (err) {
      return handleAxiosError(err, errorCallback);
    }
  };
}

// export async function withFileName(
//   responsePromise: Promise<AxiosResponse<File>>
// ): Promise<AxiosResponse<File>> {
//   const response = await responsePromise;
//   const contentDisposition: string | undefined =
//     response.headers['content-disposition'];
//   if (contentDisposition) {
//     const fileNameRegex = /filename\*?=(UTF-8''|")([^"]*)"?/;
//     const name = contentDisposition.match(fileNameRegex)?.[2];
//     if (name) {
//       const decodedName = decodeURIComponent(name);
//       (response.data as Writable<File>).name = decodedName;
//     }
//   }
//   return response;
// }
