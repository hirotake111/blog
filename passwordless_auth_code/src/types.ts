export type Ok<T> = {
  success: true;
  data: T;
};

export function Ok<T>(data: T): Ok<T> {
  return {
    success: true,
    data,
  };
}

export type Err = {
  success: false;
  detail: string;
};

export function Err(detail: string): Err {
  return {
    success: false,
    detail,
  };
}

export type Result<T> = Ok<T> | Err;
