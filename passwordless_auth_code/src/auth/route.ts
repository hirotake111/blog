import { Hono } from "hono";

import { db, auth } from "./firebase";
import { FirestoreAuthCodeRepository } from "./repository";
import { AuthCodeService } from "./service";
import { AuthHandler } from "./handler";

const fsAuthCodeRepository = new FirestoreAuthCodeRepository(db);
const authCodeService = new AuthCodeService(fsAuthCodeRepository, auth);
const authHandler = new AuthHandler(authCodeService);

export const addAuthRoutes = (app: Hono): void => {
  app.post("/auth", (c) => authHandler.handleAuth(c));
  app.post("/login", (c) => authHandler.handleLogin(c));
};
