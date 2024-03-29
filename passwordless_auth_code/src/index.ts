import { serve } from "@hono/node-server";
import { Context, Hono } from "hono";

import { db, auth } from "./firebase";
import { FirestoreAuthCodeRepository } from "./repository";
import { AuthCodeService } from "./service";

const app = new Hono();

const fsAuthCodeRepository = new FirestoreAuthCodeRepository(db);
const authCodeService = new AuthCodeService(fsAuthCodeRepository, auth);

app.get("/", async (c) => {
  return c.text("Hello Hono!");
});

app.post("/auth", async (c) => {
  const authType = c.req.query("auth_type");
  switch (authType) {
    case "code":
      return handleAuthCodeRequest(c);
    default:
      c.status(400);
      return c.json({ success: false, detail: "bad request" });
  }
});

app.post("/login", async (c) => {
  const body = await c.req.json();
  // TODO: validation
  const result = await authCodeService.validate(body.email, body.code);
  if (!result.success) {
    return c.json({ success: false, detail: result.detail });
  }
  return c.json({ success: true, customToken: result.data.customToken });
});

async function handleAuthCodeRequest(c: Context): Promise<Response> {
  const body = await c.req.json();
  const res = await authCodeService.generateCode(body.email);
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

const port = 3000;
console.log(`Server is running on port ${port}`);

serve({
  fetch: app.fetch,
  port,
});
