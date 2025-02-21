export function copy<O>(val: O, deep = true): O {
  return deep && typeof val === 'object' && val !== null
    ? JSON.parse(JSON.stringify(val))
    : val;
}

export function update<P extends object, C extends P>(context: C, params: P) {
  for (const [key, value] of Object.entries(params)) {
    context[key as keyof C] = value;
  }
}

export function pick<O extends object, K extends readonly (keyof O)[]>(
  obj: O,
  keys: K,
  { deep = true } = {}
): Pick<O, K[number]> {
  const picked = {} as Pick<O, K[number]>;
  for (const key of keys) {
    if (key in obj) picked[key] = copy(obj[key], deep);
  }
  return picked;
}

export function omit<O extends object, K extends readonly (keyof O)[]>(
  obj: O,
  keys: K,
  { deep = true } = {}
): Omit<O, K[number]> {
  const omited = {} as Omit<O, K[number]>;
  for (const key of Object.keys(obj) as (keyof O)[]) {
    if (!keys.includes(key))
      omited[key as keyof typeof omited] = copy(
        obj[key as keyof typeof omited],
        deep
      );
  }
  return omited;
}
