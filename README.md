# coolbox
A self-hosted option for common dev tools, such as markup converters, linting or code generation. Designed to be easily extensible, and for companies or individuals who don't want to paste their data online.

# todo
- [X] create a JSON to YAML converter with a web UI. Do not worry about structure, just make it work.
- [ ] create a landing page to link tools
- [ ] break down the converter into components, and attempt to modulise
- [ ] build a second converter using the new modulised system

# running
If you have air installed for hot reloading:
```sh
./air
```

otherwise:
```sh
go run cmd/run_server/main.go
```
