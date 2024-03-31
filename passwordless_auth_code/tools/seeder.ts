import { z } from "zod";
import { auth } from "../src/firebase";

const EmailSchema = z.string().email();
const email = EmailSchema.parse(process.argv[2]);

auth.createUser({ email });
console.log(`user ${email} created!`);
