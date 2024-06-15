# Golang 
Go useful materials &amp; links

## Troubleshooting

- Sometimes in `Linux` Server when we want particular type of `go version`, and even though after installing that, and when we check for the go version, it shows the previously installed go versions. In that case, it is better to check for go path locations, and delete all other unimportant go versions:

```
# first check the go path with the following command
$ whereis go
# after that remove the go path with the following command
$ sudo rm -rf /usr/local/go
```

![image](https://github.com/afa-farkhod/Go-Study/assets/24220136/6ac2ec90-6bb1-47b9-9308-dea0ba1fcf7a)

### Useful Go commands:

| Command | Description |
| --- | --- |
| `whereis go` | shows go installation path |
| `sudo rm -rf /usr/local/go/bin` | removes go path |
| `go mod tidy` | removes any unused dependencies and adds any missing dependencies to the `go.mod` file |
| `go get -d "github.com/strangelove-ventures/noble/v4@v4.1.0-rc.0"` | get the particular update version and download to `go.mod` |
| `go list -m -versions "github.com/sei-protocol/sei-chain/"` | lists available git repository versions |
| `go get -u "github.com/sei-protocol/sei-chain@7ad39e4"` | use commit hash value |

## Resources

- [w3schools](https://www.w3schools.com/java/default.asp) - tutorial platform free of charge
- [Go-Packages](https://pkg.go.dev/) - Go official packages page, where you can get any Go supported packages
- [Go-Course-YouTube](https://www.youtube.com/watch?v=yyUHQIec83I) - Golang Tutorial for Beginners | Full Go Course by TechWorld with Nana
- [Go-Book-Pdf](https://www.golang-book.com/public/pdf/gobook.pdf) - This book will teach you how to write computer programs using a programming language designed by Google named Go

