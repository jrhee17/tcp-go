# A starting point for implementing tcp in GO (OS X only)

### Running
```
sudo ./run.sh
```

You may need to set `GOPATH` and install `water` separately
(running with `sudo` doesn't resolve dependencies automatically)

```bash
go get -u github.com/songgao/water
go get -u github.com/songgao/water/waterutil
```

You should see the following tunnel when inputting ifconfig
```
> ifconfig
...
utun1: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1500
	inet 10.1.0.10 --> 10.1.0.20 netmask 0xff000000
```

### Some useful commands

`GOPATH` may not work under sudo, so we may have to specify `GOPATH` explicitly

```
sudo GOPATH=/Users/john/Projects/go-workspace go run awesomeProject/main/main.go
```

activate utun1
```
sudo ifconfig utun1 10.1.0.10 10.1.0.20 up
```

ping the generated tun

```
ping 10.1.0.20
```

Tada!!
```
Johnui-MacBook-Pro:src john$ sudo GOPATH=/Users/john/Projects/go-workspace go run awesomeProject/main/main.go 
2019/10/13 15:49:19 Interface Name: utun1
2019/10/13 15:51:25 Packet Received: 45 00 00 54 b5 a0 00 00 40 01 b0 e9 0a 01 00 0a 0a 01 00 14 08 00 60 de 70 37 00 00 5d a2 c9 6d 00 09 14 ce 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f 20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f 30 31 32 33 34 35 36 37
```
