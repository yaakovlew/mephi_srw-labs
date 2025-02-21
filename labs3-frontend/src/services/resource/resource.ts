import type { IResourceError } from 'src/services/resource/error';

export type IResource<T = unknown> =
  | {
      data: T;
      error?: undefined;
    }
  | {
      data?: undefined;
      error: IResourceError;
    };
