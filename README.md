# CloudMe 1.11.2 BufferOverflow Exploit
CloudMe 1.11.2 - Buffer Overflow (PoC) translated to Golang due the service is scoping at 127.0.0.1, it cannot be accesible from outside and the machine doesnt have Python installed, so to avoid local port forwarding / tunneling and that type of things I translated it to Golang so I can build a .exe and execute it from the machine.
## Building the Exploit to .exe for HackTheBox machine Buff
Execute on CMD:
```
go build poc.go
```
A poc.exe will be generated.
Tested on Windows 10 64bits.
