yapserve - http static server based on YAP
====

`yapserve` start a http server that serves static files.

**Usage**:

```
yapserve [-host <host>] <scheme>
```

Here `scheme` is used to create a `http.FileSystem` object. It can be:

* `<localPath>` (means it is a `local http.FileSystem`)
* `"s3:<bucketUrl>"` (means it is a `s3 http.FileSystem`)
* `"kodo:<bucket>?<accessToken>"` (means it is a `kodo http.FileSystem` provided by `qiniu.com`. How to generate a kodo url? See the `kodourl` tool below)


## kodourl

`kodourl` is a tool to generate a kodo url like `kodo:<bucket>?<accessToken>`.

**Usage**:

```
kodourl <bucket> <accessKey> <secretKey>
```

Before running `kodourl`, you should create a environment variable named `YAPSERVE_KEY`. You can generate a long random text and then assign to it. For example:

```
export YAPSERVE_KEY=<LongRandomText>
```
