---
title: "crypto Package"
description: "Hashes, HMAC and secure random numbers"
sidebar:
  order: 5
status: done
lastReviewed: 2026-11-03
---

The `crypto/...` packages provide vetted primitives. Use them as building blocks; don't invent your own schemes.

## Example

```go title="crypto.go"
sum := sha256.Sum256([]byte("hello"))
hex.EncodeToString(sum[:]) // 2cf24dba5fb0a30e...

mac := hmac.New(sha256.New, []byte("secret")) // sign webhooks, tokens
mac.Write([]byte("msg"))
sig := mac.Sum(nil)
hmac.Equal(sig, expected) // constant-time compare

token := make([]byte, 16)
rand.Read(token) // crypto/rand
```

| Need | Use |
| --- | --- |
| Checksum / fingerprint | `crypto/sha256` |
| Verify a signature with a shared secret | `crypto/hmac` |
| Tokens, IDs, keys | `crypto/rand` |
| Password storage | `golang.org/x/crypto/bcrypt` or `argon2` |
| Encrypt data | `crypto/aes` + `cipher.NewGCM` |
| TLS | `crypto/tls` (used by `net/http`) |

:::caution[Pitfalls]
- Never hash passwords with SHA-256; it's too fast. Use bcrypt or argon2.
- Never use `math/rand` for secrets.
- Compare MACs and tokens with `hmac.Equal`/`subtle.ConstantTimeCompare`, not `==`.
:::
