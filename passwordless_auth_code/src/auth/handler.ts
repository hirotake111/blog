import { Context } from "hono";

import { Result } from "../types";

interface AuthCodeService {
  validate(
    email: string,
    code: string
  ): Promise<Result<{ customToken: string }>>;
  generateCode(email: string): Promise<Result<{ code: string }>>;
}

interface EmailService {
  sendEmailWithAuthCode(
    to: string,
    code: string
  ): Promise<Result<{ success: boolean }>>;
}

export class AuthHandler {
  constructor(
    private authService: AuthCodeService,
    private emailService: EmailService
  ) {}

  public async handleAuth(c: Context) {
    const authType = c.req.query("auth_type");
    switch (authType) {
      case "code":
        return this.handleAuthCodeRequest(c);
      default:
        return c.json({ success: false, detail: "bad request" }, 400);
    }
  }

  public async handleLogin(c: Context) {
    const body = await c.req.json();
    // TODO: validation
    const result = await this.authService.validate(body.email, body.code);
    if (!result.success) {
      return c.json({ success: false, detail: result.detail }, 400);
    }
    return c.json({ success: true, customToken: result.data.customToken }, 200);
  }

  public async handleAuthCodeRequest(c: Context) {
    const body = await c.req.json();
    const codeResult = await this.authService.generateCode(body.email);
    if (!codeResult.success) {
      return c.json({ success: false, detail: codeResult.detail }, 400);
    }
    const emailResult = await this.emailService.sendEmailWithAuthCode(
      body.email,
      codeResult.data.code
    );
    if (!emailResult.success) {
      return c.json({ success: false, detail: emailResult.detail }, 400);
    }
    return c.json(
      {
        success: true,
        detail: "authentication code generated!",
      },
      201
    );
  }
}
