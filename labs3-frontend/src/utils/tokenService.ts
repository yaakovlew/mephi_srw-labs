export abstract class TokenService {
  static readonly #key = 'token';
  static readonly #type = 'type';
  static readonly #labToken = 'lab-token';

  static get token() {
    return localStorage.getItem(this.#key);
  }

  static set token(token) {
    if (token) localStorage.setItem(this.#key, token);
    else localStorage.removeItem(this.#key);
  }

  static get type() {
    return localStorage.getItem(this.#type);
  }

  static set type(type) {
    if (type) localStorage.setItem(this.#type, type);
    else localStorage.removeItem(this.#type);
  }

  static get labToken() {
    return localStorage.getItem(this.#labToken);
  }

  static set labToken(token) {
    if (token) localStorage.setItem(this.#labToken, token);
    else localStorage.removeItem(this.#labToken);
  }
}
