# Navio ![](https://github.com/viniciusbds/navio/workflows/build/badge.svg)

<img src="/cargueiro.png" alt="drawing" width="120"/>

----------------------------

**Navio** is an extremely simple tool that aims to create containers based on linux namespace, cgroups and chroot resources. The Navio goes up containers, that is, processes with namespace isolation (PID, MOUNT ...), possible limitation of the amount of resources used via cgroups and a mini operating system that currently can be:

- alpine
- busybox
- ubuntu

### What are containers?

Containers are a set of linux processes with several layers of isolation.


## [Namespaces](https://en.wikipedia.org/wiki/Linux_namespaces)

what the processes can see

**CLONE_NEWUTS** : The UTS namespace provides isolation of the hostname and domainname system identifiers

**CLONE_NEWPID** : PID namespace isolates the process ID number space. This means that two processes running on the same host can have the same PID!

**CLONE_NEWNS** : The Mount namespace isolate the filesystem mount points

---


- [x] UTS - isolate **hostname and domainname**

- [x] PID - isolate the **PID number space**

- [x] MNT - isolate **filesystem mount points**

- [ ] IPC - isolate **interprocess communication (IPC)** resources

- [ ] NET - isolate **network interfaces**

- [ ] User - isolate **UID/GID number spaces**

- [ ] Cgroup - isolate **cgroup root directory**

- [ ] Time Namespace - allows processes to see **different system times** in a way similar to the UTS namespace.


## [Cgroups](https://en.wikipedia.org/wiki/Cgroups)

What the processes can use

- [ ] Memory

- [ ] CPU

- [ ] I/O

- [ ] Process numbers


## Requirements

- [golang environment](https://golang.org/)
- make

## How to install

This project support static compiling for linux platform (amd64).

#### If you just want use the navio, is very simples: 

```
 git clone https://github.com/viniciusbds/navio.git
 cd navio
 sudo cp ./navio /usr/local/bin
```

#### If you want compile the code before running it:



```
 git clone https://github.com/viniciusbds/navio.git
 cd navio
 make
 sudo cp ./navio /usr/local/bin
```

#### To run all unit tests, type:

```
  sudo su -
  cd /path/to/project/navio
  make unit-tests
```

#### To uninstall:

 ```
 sudo rm /usr/local/bin/navio
```
  
## Example Commands

`navio get images`

`navio remove image alpine`

`navio pull alpine`

`navio run alpine sh`

...

`navio run ubuntu /bin/bash`

`navio run busybox sh`


  
  
## Contributing

You can contribute to the project in any way you want, either by fixing bugs, implementing new features, improving the documentation or proposing new features through issues

See [Contributting](/CONTRIBUTING.md) for more details

## References

  - [Containers from Scratch • Liz Rice](https://www.youtube.com/watch?v=8fi7uSYlOdc)
  
  - [Containers from Scratch](https://ericchiang.github.io/post/containers-from-scratch/)
  
  - [Building a container with less than 100 lines in Go](https://www.infoq.com/br/articles/build-a-container-golang/)

  - [Linux Namespaces](https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a)
  
  - [Namespaces](https://escotilhalivre.wordpress.com/2015/08/12/namespaces/)
  
  - <div><a href="/cargueiro.png" title="Icon">Icon</a> made by <a href="https://www.flaticon.com/br/autores/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/br/" title="Flaticon">www.flaticon.com</a></div>
