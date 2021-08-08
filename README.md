# Everything fallen to swap stays in swap
Here is my experimental proof that swapped-out pages stay in swap forever.

## Description
I launched qemu VM with the Debian buster. It was launched with the next command line arguments:

```
qemu-system-x86_64 -cdrom ~/Downloads/mini.iso -drive file=debian.img -m 1G,slots=3,maxmem=8G -smp 4
```

In the VM I launched the binary built out of [the code](./main.go). When the process starting to use swap, it slows down significantly. Then I add the RAM DIMM with command in qemu monitor:

```
object_add memory-backend-ram,id=mem1,size=1G
device_add pc-dimm,id=dimm1,memdev=mem1
```

## Results
The swapped-out pages indeed stay in swap forever until the process died. They never swapped-in even they are under active read/write operations.

### vm.swappiness=0
[![asciicast](https://asciinema.org/a/429536.svg)](https://asciinema.org/a/429536)

### vm.swappiness=60
[![asciicast](https://asciinema.org/a/429497.svg)](https://asciinema.org/a/429497)

### vm.swappiness=100
[![asciicast](https://asciinema.org/a/429537.svg)](https://asciinema.org/a/429537)
