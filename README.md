# Finger Daemon in Go

A simple Finger Daemon written in Go.

Inspired by [Happy Net Box](happynetbox.com)

## Build

Checkout the project and run

```bash
$ go build -v
```

## Run

This application need to bind to port 79 which normally requires root access. You can get round this by running
the following command

```bash
$ sudo setcap CAP_NET_BIND_SERVICE=+eip fingerd
```

This will grant the `fingerd` binary the ability to open port 79 as a normal user

```bash
$ ./fingerd
```