import { z } from "zod";

const FirestoreTimestampSchema = z.object({
  _seconds: z.number(),
  _nanoseconds: z.number(),
});

export const AuthCodeDocSchema = z.object({
  email: z.string(),
  code: z.string(),
  attempts: z.number(),
  createdAt: FirestoreTimestampSchema,
  expiresAt: FirestoreTimestampSchema,
});

export type AuthCodeDoc = z.infer<typeof AuthCodeDocSchema>;
