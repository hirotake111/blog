## Password-less Authentication Code Sample App in TypeScript

### Set up

1. Create your Firebase project.
2. Download service account JSON file from the project ([LINK](https://firebase.google.com/docs/admin/setup#initialize_the_sdk_in_non-google_environments)).
3. Rename the file as "serviceAccount.json"
4. Install packages.

```bash
npm install
```

### Test

```bash
# create a test user
npm run seed "<enter email address here>"  # ➜ user <email address> created!

# run dev server
npm run dev  # ➜ http://localhost:3000

# generate authentication code
curl -X POST "http://localhost:3000/auth?auth_type=code" -d '{"email": "user@example.com"}' -H "Content-Type: application/json"
# ➜ {"success":true,"detail":"authentication code generated!"}

# get authentication code via email and then, login using the code
curl -X POST "http://localhost:3000/login" -d '{"email": "user@example.com", "code": "1234"}' -H "Content-Type: application/json"
# ➜ {"success":true,"customToken":"..."}
```
