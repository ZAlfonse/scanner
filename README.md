# Scanner
Scans a cidr block and attempts to open tcp connections on a given port, gives you a result of alive hosts

Reqs
==
If you don't have go you can follow the getting started guide [here](https://golang.org/doc/install)

Usage
==
```
$ scanner -h
Usage of /code/go/bin/scanner:
  -cidrBlock string
        CIDR block to scan (default "10.0.0.0/8")
  -port int
        Target port (default 23)
  -timeout int
        Time (ms) to wait for established TCP connection (default 10)
  -v    Print everything
  ```
