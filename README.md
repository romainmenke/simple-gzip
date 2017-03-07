# Simple GZIP

A lazy man's compression

`go get github.com/romainmenke/simple-gzip`

---

I needed a tool to gzip css and js files for golang web projects.

- it loops over all files in a directory
- applies gzip

I use it with `//go:generate simple-gzip` and [modd](https://github.com/cortesi/modd).

---

### Options

- `-h`            : help
- `-source`       : source directory
- `-out`          : output directory
- `-level`        : compression level
- `trailing args` : exclusion -> simple `must not contain` logic

---

### Simple

- [simple-mini](https://github.com/romainmenke/simple-mini)
- [simple-bundle](https://github.com/romainmenke/simple-bundle)
- [simple-gzip](https://github.com/romainmenke/simple-gzip)
- [simple-template](https://github.com/romainmenke/simple-template)
