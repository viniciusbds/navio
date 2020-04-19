# Navio

<img src="/cargueiro.png" alt="drawing" width="120"/>

----------------------------

`Navio` é um **container runtime** extremamente simples que tem por objetivo criar contêineres com base em 
recursos de namespace, cgroups e chroot do linux. O Navio sobe contêineres, ou seja, processos com isolamento 
de namespaces (PID, MOUNT ...), limitação da quantidade de recursos usados via cgroups e um mini sistema operacional 
que atualmente pod ser: ubuntu, alpine e arch linux.


### O que são contêineres?

Contêineres são simplesmente um conjunto de processos linux com diversas camadas de isolamento. 



## Namespaces

Limits what process can see. Created with syscalls


https://www.infoq.com/br/articles/build-a-container-golang/

https://medium.com/@lets00/namespace-14c4e64d0559

https://medium.com/@teddyking/linux-namespaces-850489d3ccf

https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a

https://stackoverflow.com/questions/22889241/linux-understanding-the-mount-namespace-clone-clone-newns-flag


- [ ] Unix Timesharing System

- [ ] Process IDs

- [ ] Mounts

- [ ] Network

- [ ] User IDs

- [ ] InterProcess Comms


## Cgroups

What you can use

- [ ] Memory

- [ ] CPU

- [ ] I/O

- [ ] Process numbers



### Referências

- A espetacular talk feita por [@lizrice](https://github.com/lizrice) que [ensina o que são contêineres](https://www.youtube.com/watch?v=8fi7uSYlOdc) e como criá-los do zero de uma forma extremamente prática 


- <div>Ícone do Navio feito por <a href="https://www.flaticon.com/br/autores/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/br/" title="Flaticon">www.flaticon.com</a></div>
