<!--
SPDX-FileCopyrightText: 2025 Shun Sakai

SPDX-License-Identifier: CC0-1.0
-->

# pagen

[![CI][ci-badge]][ci-url]
[![Go Reference][reference-badge]][reference-url]
![Go version][go-version-badge]

**pagen** is a command-line utility for generating pixel art written in [Go].

The value of each pixel is determined from random numbers.

## Installation

### From source

```sh
go install github.com/sorairolake/pagen@latest
```

### From binaries

The [release page] contains pre-built binaries for Linux, macOS, Windows and
others.

### How to build

Please see [BUILD.adoc].

## Usage

### Basic usage

```sh
pagen output.png
```

## Source code

The upstream repository is available at
<https://github.com/sorairolake/pagen.git>.

## Changelog

Please see [CHANGELOG.adoc].

## Contributing

Please see [CONTRIBUTING.adoc].

## License

Copyright (C) 2025 Shun Sakai (see [AUTHORS.adoc])

This program is distributed under the terms of the _CC0 1.0 Universal_.

This project is compliant with version 3.3 of the [_REUSE Specification_]. See
copyright notices of individual files for more details on copyright and
licensing information.

[ci-badge]: https://img.shields.io/github/actions/workflow/status/sorairolake/pagen/CI.yaml?branch=develop&style=for-the-badge&logo=github&label=CI
[ci-url]: https://github.com/sorairolake/pagen/actions?query=branch%3Adevelop+workflow%3ACI++
[reference-badge]: https://img.shields.io/badge/Go-Reference-steelblue?style=for-the-badge&logo=go
[reference-url]: https://pkg.go.dev/github.com/sorairolake/pagen
[go-version-badge]: https://img.shields.io/github/go-mod/go-version/sorairolake/pagen?style=for-the-badge&logo=go
[Go]: https://go.dev/
[release page]: https://github.com/sorairolake/pagen/releases
[BUILD.adoc]: BUILD.adoc
[CHANGELOG.adoc]: CHANGELOG.adoc
[CONTRIBUTING.adoc]: CONTRIBUTING.adoc
[AUTHORS.adoc]: AUTHORS.adoc
[_REUSE Specification_]: https://reuse.software/spec-3.3/
