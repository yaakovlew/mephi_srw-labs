import { boot } from 'quasar/wrappers';
import type { AxiosInstance, AxiosRequestConfig } from 'axios';
import axios from 'axios';
import { useTokenInterceptors } from 'src/utils/token';
// import { useDatesSerializer } from 'src/utils/dates/dates';

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
  }
}

// export const baseURL = process.env.BASE_URL
// export const apiURL = baseURL + process.env.API_PATH
// export const apiPublicURL = apiURL + process.env.PUBLIC_PATH
// export const apiUserURL = apiURL + process.env.USER_PATH
// export const apiAdminURL = apiURL + process.env.ADMIN_PATH
// export const staticURL = apiPublicURL + process.env.STATIC_PATH
// export const staticUserURL = apiUserURL + process.env.STATIC_PATH
// export const staticAdminURL = apiAdminURL + process.env.STATIC_PATH

// export const baseURL = 'http://localhost:8000';
export const baseURL = process.env.BASE_URL;
export const apiURL = baseURL + '/api';
export const apiPublicURL = baseURL;
export const apiStudentURL = apiURL + '/student';
export const apiLecturerURL = apiURL + '/lecturer';
export const apiSeminarianURL = apiURL + '/seminarian';
export const staticURL = apiPublicURL + process.env.STATIC_PATH;
export const staticUserURL = apiStudentURL + process.env.STATIC_PATH;
export const staticAdminURL = apiLecturerURL + process.env.STATIC_PATH;

export const axiosCommonConfig: AxiosRequestConfig = {
  withCredentials: true,
};
export const axiosPublicConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiPublicURL,
};
export const axiosStudentConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiStudentURL,
};
export const axiosLecturerConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiLecturerURL,
};
export const axiosSeminarianConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiSeminarianURL,
};

export const $apiPublic = axios.create(axiosPublicConfig);
// useDatesSerializer($apiPublic);

export const $apiStudent = axios.create(axiosStudentConfig);
// useDatesSerializer($apiUser);
useTokenInterceptors($apiStudent);

export const $apiLecturer = axios.create(axiosLecturerConfig);
// useDatesSerializer($apiAdmin);
useTokenInterceptors($apiLecturer);

export const $apiSemianrian = axios.create(axiosSeminarianConfig);
// useDatesSerializer($apiAdmin);
useTokenInterceptors($apiSemianrian);

export const $fdHeaders = { 'Content-Type': 'multipart/form-data' };

export default boot(({ app }) => {
  app.config.globalProperties.$axios = axios;
});

export const baseURLLab = process.env.LAB_URL;
export const apiURLLab = baseURLLab;
export const apiPublicURLLab = baseURLLab;
export const apiStudentURLLab = apiURLLab;
export const apiLecturerURLLab = apiURLLab;
export const apiSeminarianURLLab = apiURLLab;
export const staticURLLab = apiPublicURLLab + process.env.STATIC_PATH;
export const staticUserURLLab = apiStudentURLLab + process.env.STATIC_PATH;
export const staticAdminURLLab = apiLecturerURLLab + process.env.STATIC_PATH;

export const axiosCommonConfigLab: AxiosRequestConfig = {
  withCredentials: true,
};
export const axiosPublicConfigLab: AxiosRequestConfig = {
  ...axiosCommonConfigLab,
  baseURL: apiPublicURLLab,
};
export const axiosStudentConfigLab: AxiosRequestConfig = {
  ...axiosCommonConfigLab,
  baseURL: apiStudentURLLab,
};
export const axiosLecturerConfigLab: AxiosRequestConfig = {
  ...axiosCommonConfigLab,
  baseURL: apiLecturerURLLab,
};
export const axiosSeminarianConfigLab: AxiosRequestConfig = {
  ...axiosCommonConfigLab,
  baseURL: apiSeminarianURLLab,
};

export const $apiPublicLab = axios.create(axiosPublicConfigLab);
// useDatesSerializer($apiPublic);

export const $apiStudentLab = axios.create(axiosStudentConfigLab);
// useDatesSerializer($apiUser);
useTokenInterceptors($apiStudentLab);

export const $apiLecturerLab = axios.create(axiosLecturerConfigLab);
// useDatesSerializer($apiAdmin);
useTokenInterceptors($apiLecturer);

export const $apiSemianrianLab = axios.create(axiosSeminarianConfigLab);
// useDatesSerializer($apiAdmin);
useTokenInterceptors($apiSemianrian);

export const baseURLLab12 = process.env.LAB12_URL;

export const apiURLLab12 = baseURLLab12;

export const apiStudentURLLab12 = apiURLLab12;

export const axiosCommonConfigLab12: AxiosRequestConfig = {
  withCredentials: true,
};

export const axiosStudentConfigLab12: AxiosRequestConfig = {
  ...axiosCommonConfigLab12,
  baseURL: apiStudentURLLab12,
};

export const $apiStudentLab12 = axios.create(axiosStudentConfigLab12);
// useDatesSerializer($apiUser);
useTokenInterceptors($apiStudentLab12);
