# Go Alethio API client (golang wrapper for API)

| WARNING: This is currently a WIP, not ready for production |
| --- |

This means that breaking changes will happen in terms of how these wrappers call Alethio's endpoints, use at your own risk.

This repo is an unofficial set of API wrappers for the [Alethio API definition](https://api.aleth.io/v1/docs) written in Golang

Testing and light use is free. If heavier use, get an API key [here](https://developers.aleth.io/)

The key is imported via environment variables, so be sure to set it via the command:

`export ALETHIO_APIKEY="your_api_key"`

Structure is modeled after:
- [go-github API client](https://github.com/google/go-github)
- [godo Digital Ocean API client](https://github.com/digitalocean/godo)

Initial pass for completion will create naive structs for each type.  Second pass will unify redundant work on structs so that they are more readable and understandable on how the API structure works.