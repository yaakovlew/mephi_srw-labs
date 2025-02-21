import type { IResourceError } from 'src/services/resource/error';

import { Notify } from 'quasar';

export function errorNotify(err: IResourceError) {
  if (
    !err.original.config?.baseURL?.includes(
      /* process.env.API_STATIC_PATH */ '#'
    )
  ) {
    Notify.create({
      type: 'negative',
      message: err.message,
    });
  }
}
