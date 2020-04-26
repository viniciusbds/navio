# Navio

<img src="/cargueiro.png" alt="drawing" width="120"/>

----------------------------

**Navio** is an extremely simple runtime container that aims to create containers based on linux namespace, cgroups and chroot resources. The ship goes up containers, that is, processes with namespace isolation (PID, MOUNT ...), possible limitation of the amount of resources used via cgroups and a mini operating system that currently can be ubuntu, alpine or arch linux.



## Namespaces

a way to limit what a process can see

**CLONE_NEWPID** : PID namespace isolates the process ID number space. This means that two processes running on the same host can have the same PID!

**CLONE_NEWUTS** : The UTS namespace provides isolation of the hostname and domainname system identifiers



https://www.infoq.com/br/articles/build-a-container-golang/

https://medium.com/@lets00/namespace-14c4e64d0559

https://medium.com/@teddyking/linux-namespaces-850489d3ccf

https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a

https://stackoverflow.com/questions/22889241/linux-understanding-the-mount-namespace-clone-clone-newns-flag


- [x] Unix Timesharing System

- [x] Process IDs

- [x] Mounts

- [ ] Network

- [ ] User IDs

- [ ] InterProcess Comms


## Cgroups

What you can use

- [ ] Memory

- [ ] CPU

- [ ] I/O

- [ ] Process numbers



<div>Ícones feitos por <a href="https://www.flaticon.com/br/autores/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/br/" title="Flaticon">www.flaticon.com</a></div>
