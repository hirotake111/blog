import { Context } from "hono";

import { Result } from "../types";

interface AuthCodeService {
  validate(
    email: string,
    code: string
  ): Promise<Result<{ customToken: string }>>;
  generateCode(email: string): Promise<Result<{ code: string }>>;
}

export class AuthHandler {
  constructor(private service: AuthCodeService) {}

  public async handleAuth(c: Context) {
    const authType = c.req.query("auth_type");
    switch (authType) {
      case "code":
        return this.handleAuthCodeRequest(c);
      default:
        c.status(400);
        return c.json({ success: false, detail: "bad request" });
    }
  }

  public async handleLogin(c: Context) {
    const body = await c.req.json();
    // TODO: validation
    const result = await this.service.validate(body.email, body.code);
    if (!result.success) {
      c.status(400);
      return c.json({ success: false, detail: result.detail });
    }
    return c.json({ success: true, customToken: result.data.customToken });
  }

  public async handleAuthCodeRequest(c: Context) {
    const body = await c.req.json();
    const res = await this.service.generateCode(body.email);
    if (!res.success) {
      c.status(400);
      return c.json({ success: false, detail: res.detail });
    }
    c.status(201);
    return c.json({
      success: true,
      detail: "authentication code generated!",
    });
  }
}
