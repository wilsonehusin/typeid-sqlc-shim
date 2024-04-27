# typeid-sqlc-shim

This is a generator for [sqlc](https://sqlc.dev)-generated Go code which uses [TypeID](https://github.com/jetify-com/typeid-go).

## Why?

To leverage compile-time enforcement, TypeID requires strict typing for every prefix's ID. On the other hand, sqlc supports overriding one data type (e.g. UUID) to another in one-to-one mapping. This is not enough for TypeID.

To make this work with TypeID, we need to generate overrides for every model field, assuming that each model has its own distinct prefix. Writing that by hand does not sound fun, so this generator is here to ease the process.

## Speedrun!

Define your config file:

```yaml
package:    # Co-locate this with models.go from sqlc.
  name: db  # => "package db"
  path: db/id.go
models:
  - name: User
    prefix: usr
    table: users
  - name: Organization
    prefix: org
    table: organizations
```

```shell-session
$ go run go.husin.dev/typeid-sqlc-shim -config path/to/config.yaml
[INFO] db/id.go has been written
[INFO] Append the following to your sqlc.yaml:

overrides:
  - column: users.id
    go_type:
      type: UserID
  - column: organizations.id
    go_type:
      type: OrganizationID

[INFO] See https://docs.sqlc.dev/en/latest/howto/overrides.html#the-go-type-map for more information.
```
