package main

import (
  "fmt"
  "os"
  "os/exec"
  "syscall"
)


// docker         run image <cmd> <params>
// go run main.go run image <cmd> <params>


func main() {
    switch os.Args[1] {
      case "run":
        run()
      case "child":
        child()
      default:
        panic("bad command")
    }
}


func run() {

  fmt.Printf("Running pai %v as %d\n", os.Args[2:], os.Getpid())


  cmd := exec.Command("/proc/self/exe", append([]string{"child"},os.Args[2:]...)...)

  cmd.SysProcAttr = &syscall.SysProcAttr  {
   	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID ,
  }

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr


  check(cmd.Run())
}


func child() {
  fmt.Printf("Running child %v as %d\n", os.Args[2:], os.Getpid())

  syscall.Sethostname([]byte("container"))


  syscall.Chroot("../images/ubuntu")
  os.Chdir("/")


  os.Mkdir("proc", 0700)
  check(syscall.Mount("proc", "proc", "proc", 0, ""))

  cmd := exec.Command(os.Args[2], os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  check(cmd.Run())
  check(syscall.Unmount("proc", 0))

}


func check(err error) {
  if err != nil {
    fmt.Println("ERROR", err)
    os.Exit(1)
  }
}