/* eslint-disable */

declare namespace NodeJS {
  interface ProcessEnv {
    NODE_ENV: string;
    VUE_ROUTER_MODE: 'hash' | 'history' | 'abstract' | undefined;
    VUE_ROUTER_BASE: string | undefined;
    readonly BASE_URL: string;
    readonly API_PATH: string;
    readonly PUBLIC_PATH: string;
    readonly USER_PATH: string;
    readonly ADMIN_PATH: string;
    readonly STATIC_PATH: string;
  }
}
