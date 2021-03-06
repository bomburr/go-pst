<h1 align="center">
  <br>
  <a href="https://github.com/mooijtech/go-pst"><img src="https://www.go-pst.org/assets/images/logo.svg" alt="go-pst" width="369"></a>
  <br>
  go-pst
  <br>
</h1>

<h4 align="center">A fast library for reading PST, OST and PAB files in Go (Golang).</h4>

<p align="center">
  <a href="https://github.com/mooijtech/go-pst/blob/master/LICENSE.txt" target="_blank">
      <img src="https://img.shields.io/badge/license-MIT-199473?style=flat-square" alt="License">
  </a>
  <a href="https://github.com/mooijtech/go-pst" target="_blank">
    <img src="https://img.shields.io/badge/version-0.0.1-4D7CFE?style=flat-square" alt="Version">
  </a>
  <a href="https://travis-ci.org/github/mooijtech/go-pst" target="_blank">
    <img src="https://img.shields.io/travis/mooijtech/go-pst/master?style=flat-square" alt="Build">
  </a>
  <a href="https://github.com/mooijtech/go-pst" target="_blank">
      <img src="https://img.shields.io/badge/contributions-welcome-9446ED?style=flat-square" alt="Contributions">
  </a>
</p>

---

#### This library doesn't work yet.

The PFF (Personal Folder File) and OFF (Offline Folder File) format is used to store Microsoft Outlook e-mails, appointments and contacts. 
The PST (Personal Storage Table), OST (Offline Storage Table) and PAB (Personal Address Book) file format consist of the PFF format.

## Goals

- Support PST, OST and PAB files
- Provide a reliable built-in parser 
- Be asynchronous and event-driven
- Have a well documented API
- Understand encrypted files 
  - On by default with newer versions of Outlook
- Extract as much data as possible
- Quickly retrieve messages by their identifiers
- Flexible storage system
- Learn a lot and have fun

## Usage

```bash
```

## Example

```go
```

## Documentation

- [go-pst API]()
- [File format specification](https://github.com/mooijtech/go-pst/tree/master/docs)

## Motivation

[Mooij Forensics](https://www.mooijforensics.com/) allows you to analyse and present digital evidence.

## Versioning

go-pst will be maintained under the Semantic Versioning guidelines as much as possible. <br/>
API releases will be numbered with the following format:
```
<major>.<minor>.<patch>
```

And constructed with the following guidelines:
- Breaking backward compatibility bumps the major
- New additions without breaking backward compatibility bumps the minor
- Bug fixes and misc changes bump the patch

For more information on SemVer, please visit: https://semver.org/

## Issues

Feel free to submit any issues or feature requests [here](https://github.com/mooijtech/go-pst/issues).

## References

- [java-libpst](https://github.com/rjohnsondev/java-libpst)
- [libpff](https://github.com/libyal/libpff)

## License

[MIT](https://github.com/mooijtech/go-pst/blob/master/LICENSE.txt)
