import type { AxiosError, AxiosInstance, AxiosRequestConfig } from 'axios';
import { AuthService } from 'src/services/auth';
import { TokenService } from './tokenService';

export function useTokenInterceptors($api: AxiosInstance) {
  useAccessInterceptor($api);
  // useRefreshInterceptor($api);
}

function useAccessInterceptor($api: AxiosInstance) {
  $api.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${TokenService.token}`;
    config.headers['lab-token'] = TokenService.labToken;
    return config;
  });
}

function useRefreshInterceptor($api: AxiosInstance) {
  $api.interceptors.response.use(
    undefined,
    async (error: AuthInterceptorError) => {
      const originalRequest = error.config;
      if (error.response?.status === 401 && !originalRequest._isRetry) {
        originalRequest._isRetry = true;
        const res = await AuthService.refresh();
        if (!res.error) return $api.request(originalRequest);
        TokenService.token = null;
        throw error;
      }
      throw error;
    }
  );
}

//eslint-disable-next-line
//@ts-ignore
interface AuthInterceptorError extends AxiosError {
  config: AxiosRequestConfig & { _isRetry: boolean | undefined };
}
