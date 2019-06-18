# Simple echo server
This is a simple echo server with auth header on constant

#### Run
- first ```git clone github.com/xsikor/echo-server```
- next do ```GO111MODULE=on go run main.go```

#### How to connect
After run - you will get running http/ws server on port 8080 or what you
set to -addr option before run command

For connect to echo server you need set custom header ```secretKey=54686973206973206d7920626f6f6d737469636b```\
in other case you will get error before connect upgrade on WS

##### Task

Create simple CLI for this echo-server using [github.com/spf13/cobra](github.com/spf13/cobra) and websocket client (gorilla/x/any)

CLI must be able to make auth before first command.\
You need add `-auth` command that will be save given secret to file in tmp directory (windows/linux/mac) what will be used in next request for auth\
if program can't find this file with secret - print help msg how to auth

CLI must accept command like `-user-create myUser` then send command to echo server and print result to console
If you will get bad status (no auth, or bad secret) print to console with error
