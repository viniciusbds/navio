# Navio

<img src="/cargueiro.png" alt="drawing" width="120"/>

----------------------------

**Navio** is an extremely simple runtime container that aims to create containers based on linux namespace, cgroups and chroot resources. The ship goes up containers, that is, processes with namespace isolation (PID, MOUNT ...), possible limitation of the amount of resources used via cgroups and a mini operating system that currently can be ubuntu, alpine or arch linux.



## Namespaces

a way to limit what a process can see

**CLONE_NEWPID** : PID namespace isolates the process ID number space. This means that two processes running on the same host can have the same PID!

**CLONE_NEWUTS** : The UTS namespace provides isolation of the hostname and domainname system identifiers

**CLONE_NEWNS** : The Mount namespace isolate the ...



https://www.infoq.com/br/articles/build-a-container-golang/


https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a
https://escotilhalivre.wordpress.com/2015/08/12/namespaces/

https://stackoverflow.com/questions/22889241/linux-understanding-the-mount-namespace-clone-clone-newns-flag




- [x] UTS - isolate hostname and domainname

- [x] PID - isolate the PID number space

- [x] Mount - isolate filesystem mount points

- [ ] IPC - isolate interprocess communication (IPC) resources

- [ ] Network - isolate network interfaces

- [ ] User - isolate UID/GID number spaces

- [ ] Cgroup - isolate cgroup root directory


## Cgroups

What you can use

- [ ] Memory

- [ ] CPU

- [ ] I/O

- [ ] Process numbers



### Source

  - [Linux Namespaces](https://medium.com/@teddyking/linux-namespaces-850489d3ccf)
  
<div>Ícones feitos por <a href="https://www.flaticon.com/br/autores/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/br/" title="Flaticon">www.flaticon.com</a></div>
