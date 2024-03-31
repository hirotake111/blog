import { Firestore } from "firebase-admin/firestore";
import { AuthCodeDoc, AuthCodeDocSchema } from "./domain";
import { Err, Ok, Result } from "../types";

const AUTH_CODE_COLLECTION = "auth_code";

export class FirestoreAuthCodeRepository {
  private readonly db: Firestore;
  constructor(db: Firestore) {
    this.db = db;
  }

  public async get(uid: string): Promise<Result<AuthCodeDoc>> {
    const snapshot = await this.db
      .collection(AUTH_CODE_COLLECTION)
      .doc(uid)
      .get();
    if (!snapshot.exists) {
      return Err("data not found in db");
    }
    const parsedData = AuthCodeDocSchema.safeParse(snapshot.data());
    if (!parsedData.success) {
      await this.delete(uid);
      return Err("failed to parse data in db");
    }
    const authCodeDoc = parsedData.data;
    return Ok(authCodeDoc);
  }

  public async create(uid: string, newDoc: any): Promise<Result<any>> {
    try {
      await this.db.collection(AUTH_CODE_COLLECTION).doc(uid).set(newDoc);
      return Ok(newDoc);
    } catch (err) {
      console.log(err);
      return Err("failed to create doc");
    }
  }

  public async update(
    uid: string,
    newDoc: Partial<AuthCodeDoc>
  ): Promise<Result<void>> {
    try {
      await this.db.collection(AUTH_CODE_COLLECTION).doc(uid).update(newDoc);
      return Ok(undefined);
    } catch (err) {
      console.log(err);
      return Err("failed to update doc");
    }
  }

  public async delete(uid: string): Promise<Result<void>> {
    try {
      await this.db.collection(AUTH_CODE_COLLECTION).doc(uid).delete();
      return Ok(undefined);
    } catch (e) {
      console.log(e);
      return Err("failed to delete doc");
    }
  }
}
