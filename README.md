# certchecker

## Summary

Certificate checker is a command-line application written in Go that can perform
an ad-hoc check to web servers. This check will return basic information about
the web servers certificate. More importanly, if it's expired. For now this is
the beginning stage of the application. But, the overall vision is to have an
app that will iterate through a list of servers and provide this data in various
reporting formats. Such as email, pdf, txt, csv, etc.

This can be accomplished with Nagios or OpenSSL. But, this is intended for use
outside of standard monitoring platforms or just to provide output that's
simple. Plus, it's been a desired project of mine for learning Go and to see
where Go can be used.

## Build

The current build will be undergoing a refactoring at some point.

Latest Stable Build: v1.0.0

Build process is to build it and move it to the /usr/local/bin. Which seems to
work on the Linux systems I've tested.

I have not tested this on Mac or Windows. Upon testing, I will have more
results.

Build the application.

```sh
go build
```

Move it to the /usr/local/bin directory or any other path you desire.

```sh
sudo mv ./certchecker /usr/local/bin
```

## Usage

Applciation can be run by issuing the command with a target host. Not URL. The
application will take care of the url portion for you for now.

```sh
certchecker -s www.n3s0.tech
```

The output for this command can be found below.

```sh
|---------------------------------------|
|-- Certificate Checker (certchecker) --|
|---------------------------------------|
╭─ Server: www.n3s0.tech
│  ├─ TLS Version: TLS 1.3
│  ├─ Cipher Suite: TLS_AES_128_GCM_SHA256
│  ╰─ Subject: CN=www.n3s0.tech
├─ Certificate Dates:
│  ├─ Not Before: Sun, 17 Nov 2024 19:44:48 CST
│  ╰─ Not After: Sat, 15 Feb 2025 19:44:47 CST
│     ╰─ Status: Valid
╰─ Client Information:
   ╰─ Local Date/Time: Sun, 15 Dec 2024 03:36:48 CST
```
