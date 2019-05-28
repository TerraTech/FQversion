[![GoDoc](https://godoc.org/github.com/TerraTech/FQversion?status.svg)](https://godoc.org/github.com/TerraTech/FQversion)

## FQversion

Quick 'Version' helper library that I use in various projects that was written
for a larger private enterprise library, but has been made standalone and public.
This removes the need to make this vendored and helps those that use go modules.

This library supports a versioning catalog where multiple imported libraries can
register their Name, Version, Build information and have it emitted in various ways.
