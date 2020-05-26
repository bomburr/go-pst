<h1 align="center">
  <br>
  <a href="https://github.com/mooijtech/go-pst"><img src="https://i.imgur.com/qE8QsP6.png" alt="go-pst" width="320"></a>
  <br>
  go-pst
  <br>
</h1>

<h4 align="center">A fast library for reading PFF/OFF/PST/OST/PAB files in Go (Golang).</h4>

<p align="center">
  <a href="https://github.com/mooijtech/go-pst/blob/master/LICENSE.txt">
      <img src="https://img.shields.io/badge/license-MIT-199473?style=flat-square">
  </a>
  <a href="https://github.com/mooijtech/go-pst">
    <img src="https://img.shields.io/badge/version-0.0.1-4D7CFE?style=flat-square">
  </a>
  <a href="https://travis-ci.org/github/mooijtech/go-pst">
    <img src="https://travis-ci.org/mooijtech/go-pst.svg?branch=master">
  </a>
  <a href="https://github.com/mooijtech/go-pst">
      <img src="https://img.shields.io/badge/contributions-welcome-9446ED?style=flat-square">
  </a>
</p>

---

#### This library doesn't work yet.
The PFF (Personal Folder File) and OFF (Offline Folder File) format is used to store Microsoft Outlook e-mails, appointments and contacts. 
The PST (Personal Storage Table), OST (Offline Storage Table) and PAB (Personal Address Book) file format consist of the PFF format.

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

## Documentation

[File format specification](https://github.com/mooijtech/go-pst/tree/master/docs)

## References

- [java-libpst](https://github.com/rjohnsondev/java-libpst)
- [libpff](https://github.com/libyal/libpff)

## License

[MIT](https://github.com/mooijtech/go-pst/blob/master/LICENSE.txt)
