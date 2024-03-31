import { AuthCodeDoc } from "./domain";
import { Err, Ok, Result } from "../types";

type UserRecord = { uid: string };

interface DBRepository {
  get(uid: string): Promise<Result<AuthCodeDoc>>;
  create(uid: string, doc: any): Promise<Result<AuthCodeDoc>>;
  update(uid: string, doc: Partial<AuthCodeDoc>): Promise<Result<void>>;
  delete(uid: string): Promise<Result<void>>;
}

interface AuthRepository {
  getUserByEmail(email: string): Promise<UserRecord>;
  createCustomToken(uid: string): Promise<string>;
}

const MAX_ATTEMPTS = 3;
const EXPIRATION_SEC = 60 * 1_000;

export class AuthCodeService {
  constructor(private db: DBRepository, private auth: AuthRepository) {}

  public async validate(
    email: string,
    code: string
  ): Promise<Result<{ customToken: string }>> {
    // get uid
    let userRecord: UserRecord;
    try {
      userRecord = await this.auth.getUserByEmail(email);
    } catch (err) {
      console.log({ err });
      return Err("user not found");
    }
    // get authentication code data from db
    const doc = await this.db.get(userRecord.uid);
    if (!doc.success) {
      return Err(doc.detail);
    }
    const authCodeDoc = doc.data;
    if (authCodeDoc.expiresAt._seconds < Date.now() / 1000) {
      await this.db.delete(userRecord.uid);
      return Err("code expired");
    }
    if (authCodeDoc.code !== code) {
      await this.db.update(userRecord.uid, {
        email: authCodeDoc.email,
        code: authCodeDoc.code,
        attempts: authCodeDoc.attempts + 1,
      });
      if (authCodeDoc.attempts + 1 >= MAX_ATTEMPTS) {
        await this.db.delete(userRecord.uid);
      }
      return Err("invalid code");
    }
    const customToken = await this.auth.createCustomToken(userRecord.uid);
    await this.db.delete(userRecord.uid);
    return Ok({ customToken });
  }

  public async generateCode(email: string): Promise<Result<{ code: string }>> {
    let userRecord: UserRecord;
    try {
      userRecord = await this.auth.getUserByEmail(email);
    } catch (err) {
      console.log({ err });
      return Err("user not found");
    }
    // generate authentication code
    const code = getRandomCode(4);
    const now = Date.now();
    // store code to database
    const result = await this.db.create(userRecord.uid, {
      email,
      code,
      attempts: 0,
      expiresAt: new Date(now + EXPIRATION_SEC),
      createdAt: new Date(now),
    });
    if (!result.success) {
      return Err(result.detail);
    }
    // send code to email
    await sendEmailWithAuthCode(email, code);
    return Ok({ code });
  }
}

function getRandomCode(length: number) {
  let code = "";
  for (let i = 0; i < length; i++) {
    code += Math.floor(Math.random() * 10);
  }
  return code;
}

async function sendEmailWithAuthCode(
  to: string,
  code: string
): Promise<Result<{ success: boolean }>> {
  try {
    console.log(`sending email to ${to} with authentication code: ${code}`);
    return Ok({ success: true });
  } catch (e) {
    console.log(e);
    return Err("failed to send email to user");
  }
}
