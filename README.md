# GoEmailReverseShell
### This golng reverse shell fetches the address from your latest email subject, and the subject should be like this:  
`10.1.1.1:1234`  
### This golag file requires go-pop3 lib which you can get it here:  
https://github.com/taknb2nch/go-pop3.git  
### If you want to hide the window when executing it, compile the program like this:  
```go build -ldflags -H=windowsgui email_reverse_shell-repl.go```  
