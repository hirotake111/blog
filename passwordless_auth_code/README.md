## Password-less Authentication Code Sample App in TypeScript

### Set up

```bash
npm install
```

### Test

```bash
# create a test user
bun run seed "<enter email address here>"  # ➜ user <email address> created!

# run dev server
npm run dev  # ➜ http://localhost:3000

# generate authentication code
curl -X POST "http://localhost:3000/auth?auth_type=code" -d '{"email": "user@example.com"}' -H "Content-Type: application/json"
# ➜ {"success":true,"detail":"authentication code generated!"}

# get authentication code via email and then, login using the code
curl -X POST "http://localhost:3000/login" -d '{"email": "user@example.com", "code": "1234"}' -H "Content-Type: application/json"
# ➜ {"success":true,"customToken":"..."}
```
