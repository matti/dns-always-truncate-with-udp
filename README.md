# dns-always-truncate-with-udp

a dns proxy which always responds truncated (TC) with udp, but not with tcp

    go run main.go udp
    go run main.go tcp

and

    $ dig @127.0.0.1 www.microsoft.com
    ;; Truncated, retrying in TCP mode.
    ...
