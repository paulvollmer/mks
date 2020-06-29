<p align="center">
  <h3 align="center">mks</h3>
  <p align="center">mks is a tool to create files and the intermediate directories.</p>
  <p align="center">
    <a href="https://github.com/paulvollmer/mks/actions"><img alt="Github Actions" src="https://github.com/paulvollmer/mks/workflows/CI/badge.svg?style=flat-square"> </a>
    <a href="https://github.com/paulvollmer/mks/releases"><img alt="Software Release" src="https://img.shields.io/github/v/release/paulvollmer/mks.svg?include_prereleases&style=flat-square&color=blue"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square"></a>
  </p>
</p>

---

## Installation

Install with brew

```sh
brew install paulvollmer/tap/mks
```

Install with go

```sh
go get -u github.com/paulvollmer/mks
```

## Usage

The philosophy behind `mks` is simplicity. With the following command you can create a file `foo.txt` inside the `path/to` directory.

```sh
mks path/to/foo.txt
```

## Development

```sh
git clone git@github.com:paulvollmer/mks.git
cd mks
make
# is equivalent to
# make fmt
# make build
```

## License

[MIT License](LICENSE)
