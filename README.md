# td [![Go Reference](https://img.shields.io/badge/go-pkg-00ADD8)](https://pkg.go.dev/github.com/gotd/td#section-documentation) [![codecov](https://img.shields.io/codecov/c/github/gotd/td?label=cover)](https://codecov.io/gh/gotd/td) [![beta](https://img.shields.io/badge/-beta-yellow)](https://go-faster.org/docs/projects/status#beta)

A fast Telegram client in pure Go.

* [Examples](examples)
* [Security policy](.github/SECURITY.md)
* [User support and dev chat](.github/SUPPORT.md)
* [Roadmap](ROADMAP.md)
* [Contributing](CONTRIBUTING.md)
* [Architecture](ARCHITECTURE.md)
* [Generated Go Documentation](https://ref.gotd.dev/pkg/github.com/gotd/td/tg.html)

Before using this library, read [How To Not Get Banned](.github/SUPPORT.md#how-to-not-get-banned) guide.

Due to limitations of `pkg.go.dev`, documentation for `tg` package is not shown, but there is [hosted version](https://ref.gotd.dev/pkg/github.com/gotd/td/tg.html).

## Usage

```console
go get github.com/gotd/td
```

```go
package main

import (
	"context"

	"github.com/gotd/td/telegram"
)

func main() {
	// https://core.telegram.org/api/obtaining_api_id
	client := telegram.NewClient(appID, appHash, telegram.Options{})
	if err := client.Run(context.Background(), func(ctx context.Context) error {
		// It is only valid to use client while this function is not returned
		// and ctx is not cancelled.
		api := client.API()

		// Now you can invoke MTProto RPC requests by calling the API.
		// ...

		// Return to close client connection and free up resources.
		return nil
	}); err != nil {
		panic(err)
	}
	// Client is closed.
}
```

See [examples](examples) for more info.

## Features

* Full MTProto 2.0 implementation in Golang, directly access any MTProto method with [telegram.Client.API()](https://pkg.go.dev/github.com/gotd/td@v0.60.0/telegram#Client.API)
* Highly optimized, low memory (150kb per idle client) and CPU overhead, can handle thousands concurrent clients
* Code for Telegram types generated by `./cmd/gotdgen` (based on [gotd/tl](https://github.com/gotd/tl) parser) with embedded [official documentation](https://core.telegram.org/schema)
* Pluggable session storage
* Automatic re-connects with keepalive
* Vendored Telegram public keys that are kept up-to-date
* Rigorously tested
  * End-to-end with real Telegram server in CI
  * End-to-end with gotd Telegram server (in pure Go)
  * Lots of unit testing
  * Fuzzing
  * 24/7 canary bot in production that tests reconnects, update handling, memory leaks and performance
* No runtime reflection overhead
* [Conforms](https://github.com/gotd/td/issues/155) to [Security guidelines](https://core.telegram.org/mtproto/security_guidelines) for Telegram client software developers
  * Secure PRNG used for crypto
  * Replay attack protection
* 2FA support
* MTProxy support
* Various helpers that lighten the complexity of the Telegram API
  * [uploads](https://pkg.go.dev/github.com/gotd/td/telegram/uploader) for big and small files with multiple streams for single file and progress reporting
  * [downloads](https://pkg.go.dev/github.com/gotd/td/telegram/downloader) with CDN support, also multiple streams
  * [messages](https://pkg.go.dev/github.com/gotd/td/telegram/message) with various convenience builders and text styling support
  * [query](https://pkg.go.dev/github.com/gotd/td/telegram/query) with pagination helpers
  * [middleware](https://pkg.go.dev/github.com/gotd/td/telegram#Middleware) for [rate limiting](https://pkg.go.dev/github.com/gotd/contrib/middleware/ratelimit) and [FLOOD_WAIT handling](https://pkg.go.dev/github.com/gotd/contrib/middleware/floodwait)
* Connection pooling
* Automatic datacenter migration and redirects handling
* Graceful [request cancellation](https://core.telegram.org/mtproto/service_messages#cancellation-of-an-rpc-query) via context
* WebSocket transport support (works in WASM)

## Status

The goal of this project is to implement a stable, performant and safe client for Telegram in pure Go while
having a simple and convenient API and a feature parity with TDLib.

This project is fully non-commercial and not affiliated with any commercial organization
(including Telegram LLC).

## Examples

See [examples](examples) directory.

Also take a look at

* [go-faster/bot](https://github.com/go-faster/bot) as example of sophisticated telegram bot integration with GitHub
* [gotd/cli](https://github.com/gotd/cli), command line interface for subset of telegram methods.

### Auth

#### User

You can use `td/telegram/auth.Flow` to simplify user authentications.

```go
codePrompt := func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
    // NB: Use "golang.org/x/crypto/ssh/terminal" to prompt password.
    fmt.Print("Enter code: ")
    code, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(code), nil
}
// This will setup and perform authentication flow.
// If account does not require 2FA password, use telegram.CodeOnlyAuth
// instead of telegram.ConstantAuth.
if err := auth.NewFlow(
    auth.Constant(phone, password, auth.CodeAuthenticatorFunc(codePrompt)),
    auth.SendCodeOptions{},
).Run(ctx, client.Auth()); err != nil {
    panic(err)
}
```

#### Bot

Use bot token from [@BotFather](https://telegram.me/BotFather).

```go
if err := client.Auth().Bot(ctx, "token:12345"); err != nil {
    panic(err)
}
```

### Calling MTProto directly

You can use the generated `tg.Client` that allows calling any MTProto method
directly.

```go
// Grab these from https://my.telegram.org/apps.
// Never share it or hardcode!
client := telegram.NewClient(appID, appHash, telegram.Options{})
client.Run(ctx, func(ctx context.Context) error {
  // Grab token from @BotFather.
  if _, err := client.Auth().Bot(ctx, "token:12345"); err != nil {
    return err
  }
  state, err := client.API().UpdatesGetState(ctx)
  if err != nil {
    return err
  }
  // Got state: &{Pts:197 Qts:0 Date:1606855030 Seq:1 UnreadCount:106}
  // This will close client and cleanup resources.
  return nil
})
```

### Generated code

The code output of `gotdgen` contains references to TL types, examples, URL to
official documentation and [extracted](https://github.com/gotd/getdoc) comments from it.

For example, the [auth.Authorization](https://core.telegram.org/type/auth.Authorization) type in `tg/tl_auth_authorization_gen.go`:

```go
// AuthAuthorizationClass represents auth.Authorization generic type.
//
// See https://core.telegram.org/type/auth.Authorization for reference.
//
// Example:
//  g, err := DecodeAuthAuthorization(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthAuthorization: // auth.authorization#cd050916
//  case *AuthAuthorizationSignUpRequired: // auth.authorizationSignUpRequired#44747e9a
//  default: panic(v)
//  }
type AuthAuthorizationClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthAuthorizationClass
}
```

Also, the corresponding [auth.signIn](https://core.telegram.org/method/auth.signIn) method:

```go
// AuthSignIn invokes method auth.signIn#bcd51581 returning error if any.
// Signs in a user with a validated phone number.
//
// See https://core.telegram.org/method/auth.signIn for reference.
func (c *Client) AuthSignIn(ctx context.Context, request *AuthSignInRequest) (AuthAuthorizationClass, error) {}
```

The generated constructors contain detailed official documentation, including links:

```go
// FoldersDeleteFolderRequest represents TL type `folders.deleteFolder#1c295881`.
// Delete a peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/method/folders.deleteFolder for reference.
type FoldersDeleteFolderRequest struct {
    // Peer folder ID, for more info click here¹
    //
    // Links:
    //  1) https://core.telegram.org/api/folders#peer-folders
    FolderID int
}
```

## Contributions

Huge thanks to all contributors. Dealing with a project of this scale alone is impossible.

Special thanks:

* [tdakkota](https://github.com/tdakkota)
  * Two-factor authentication (SRP)
  * Proxy support
  * Update dispatcher
  * Complete transport support (abridged, padded intermediate and full)
  * Telegram server for end-to-end testing
  * Multiple major refactorings, including critical cryptographical scope reduction
  * Code generation improvements (vector support, multiple modes for pretty-print)
  * And many other cool things and performance improvements
* [shadowspore](https://github.com/shadowspore)
  * Background pings
  * Links in generated documentation
  * Message acknowledgements
  * Retries
  * RPC Engine
  * Gap (Updates) engine

## Reference

The MTProto protocol description is [hosted](https://core.telegram.org/mtproto#general-description) by Telegram.

Most important parts for client implementations:

* [Security guidelines](https://core.telegram.org/mtproto/security_guidelines) for client software developers

Current implementation [mostly conforms](https://github.com/gotd/td/issues/155) to security guidelines, but no
formal security audit were performed.

## Prior art

* [Lonami/grammers](https://github.com/Lonami/grammers) (Great Telegram client in Rust, many test vectors were used as reference)
* [sdidyk/mtproto](https://github.com/sdidyk/mtproto), [cjongseok/mtproto](https://github.com/cjongseok/mtproto), [xelaj/mtproto](https://github.com/xelaj/mtproto) (MTProto 1.0 in go)

## Who is using gotd?

[<img src="https://user-images.githubusercontent.com/43930873/142855897-7091ced0-4fe8-4f8d-ad43-e9db2723bacc.png" width="150">](https://telq.org)

Drop a comment [here](https://github.com/gotd/td/issues/568) to add your project.

## License

MIT License

Created by Aleksandr (ernado) Razumov

2020

## Links

- [Examples](https://github.com/gotd/td/blob/-/examples/README.md)
- [Security policy](https://github.com/gotd/td/blob/-/.github/SECURITY.md)
- [User support and dev chat](https://github.com/gotd/td/blob/-/.github/SUPPORT.md)
- [Roadmap](https://github.com/gotd/td/blob/-/ROADMAP.md)
- [Contributing](https://github.com/gotd/td/blob/-/CONTRIBUTING.md)
- [Architecture](https://github.com/gotd/td/blob/-/ARCHITECTURE.md)
- [Generated Go Documentation](https://ref.gotd.dev/pkg/github.com/gotd/td/tg.html)
