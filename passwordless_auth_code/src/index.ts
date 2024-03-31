import { serve } from "@hono/node-server";
import { Hono } from "hono";
import { logger } from "hono/logger";

import { addAuthRoutes } from "./auth/route";

const app = new Hono();

app.use(logger());

app.get("/", async (c) => c.text("Hello Hono!"));

addAuthRoutes(app);

const port = 3000;
console.log(`Server is running on port ${port}`);

serve({
  fetch: app.fetch,
  port,
});
