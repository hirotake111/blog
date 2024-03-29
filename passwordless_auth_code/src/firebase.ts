import admin from "firebase-admin";
import { getFirestore } from "firebase-admin/firestore";

import serviceAccount from "../secret.json";
import { getAuth } from "firebase-admin/auth";

admin.initializeApp({
  credential: admin.credential.cert(serviceAccount as admin.ServiceAccount),
});

const db = getFirestore();
const auth = getAuth();

export { db, auth };
