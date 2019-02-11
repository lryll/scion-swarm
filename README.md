# Info

The code is based on [go-ethereum](https://github.com/ethereum/go-ethereum) (commit [54f650](https://github.com/ethereum/go-ethereum/commit/54f650a3be2ccf7cd44e9929e3e132ef93f101ad)) and Aleksandar Vorkapic's [SCION-Ethereum](https://github.com/Aleksandaarrr/ETH-SCION-ONLY) (commit [f97586](https://github.com/Aleksandaarrr/ETH-SCION-ONLY/commit/f975862dda65758ed29eb8b8b6bc9fbd7d471a16)).

Any hints, suggestions or PRs appreciated.

##### Note

In the [swarm docs](https://swarm-guide.readthedocs.io/en/latest/introduction.html), swarm is considered as "experimental code and untested in the wild". This applies to SCION-Swarm as well.

Currently, the transmission rate of files may be reduced. This problem is most likely related to node connectivity (_err="no peer found"_ messages that can occur during chunk transmission) and is under investigation. Rebasing scion-swarm to the latest swarm version may help. Furthermore, reviewing the SCION integration in this project, which slightly changed after the last SCIONLab update, may be beneficial.

# Installation

*Note*: We will be using the legacy ethereum directory structure to allow the source code and all dependencies to compile smoothly. Compile _swarm_ and _geth_ as well. _Geth_ allows access to the swarm node via a javascript console (see [Swarm Documentation](https://swarm-guide.readthedocs.io/en/latest/introduction.html)).

```
mkdir -p $GOPATH/src/github.com/ethereum
cd $GOPATH/src/github.com/ethereum
git clone https://github.com/lryll/scion-swarm go-ethereum

cd go-ethereum/vendor/github.com/lucas-clemente/quic-go/
git apply ../../../../changeQUICVersion.patch
cd ../../../../
make swarm geth
cp build/bin/* $GOPATH/bin/
```

# Tutorial - Send a file via SCION-Swarm

We send a file via SCION-Swarm over *SCIONLab*, which constitutes the SCION testbed. This requires two different SCION ASes, on different machines (or VMs) that are successfully connected to SCIONLab and SCION-Swarm installed on each of them. Refer to the [SCION Tutorials](https://netsec-ethz.github.io/scion-tutorials) in case of SCION related questions.

A swarm node requires an Ethereum account to work properly. Therefore, this repo already contains two subdirs with two different accounts (~scion/n1 and ~scion/n2) with the password _scion_ for both.

### 1. Connect nodes

Replace the AS, the IP address of the _-scion_ flag with your own configuration. The flag _--bzzport_ constitutes the port for uploading/retrieving files through swarms REST API. Start from main go-ethereum directory.

#### Node 1 (serves file), AS A
```
cd /home/$USER/go/src/github.com/ethereum/go-ethereum
swarm --bzzaccount 7dedcee4f25d5100146f8b1a4e9b47f91379a1c4 --datadir /home/$USER/go/src/github.com/ethereum/go-ethereum/~scion/n1 --keystore /home/$USER/go/src/github.com/ethereum/go-ethereum/~scion/n1/keystore --password /home/$USER/go/src/github.com/ethereum/go-ethereum/~scion/password.txt --bzzport 5001 -scion "19-ffaa:1:121,[10.0.8.83]:30306"
```

#### Node 2 (retrieves file), AS B

Replace the AS, the IP address of the _-scion_ flag with your own configuration of your chosen AS B. This is the SCION address of this node. Also substitute all other occurences of ASes, IP and port in the argument --bootnodesSCION , which is the address of the previous node 1 in AS A.
```
swarm --bzzaccount b0b0da8b63fddb44010759d6f241d22895bacee1 --datadir /home/$USER/go/src/github.com/ethereum/go-ethereum/~scion/n2 --keystore /home/$USER/go/src/github.com/ethereum/go-ethereum/~scion/n2/keystore --password /home/$USER/go/src/github.com/ethereum/go-ethereum/~scion/password.txt --bzzport 5002 --bootnodesSCION="enode://78e42c61a89929f422c97fe1b0f3c6eb807e67d77e6becf56d68dcf9df9ee51e5d9f56d961d51badc4a59a2b67bbc2572042b22df58bca068a491322daab72b9@10.0.8.83:30306@SCION@19-ffaa:1:121,[10.0.8.83]:30306" -scion "20-ffaa:1:124,[10.0.8.29]:30309"
```

### 2. Upload file

Open a new console in your AS A. Create a random file, upload it and copy the hash.
```
head -c 1MB </dev/urandom > random1MB.file
swarm --bzzapi http://localhost:5001 up \~scion/files/data/random1MB.file

... 14c81a7e11bcd431534230b73c7f78d1cda5dc16ae2990c77fe1f23a7d99e11a
```
### 3. Download file
Open a new console in your AS B. Replace the hash in the command with yours and download the file from node 2 in AS B.
```
curl http://localhost:5002/bzz:/14c81a7e11bcd431534230b73c7f78d1cda5dc16ae2990c77fe1f23a7d99e11a/ >> /vagrant/random1MB.file
```

# Tipps & functionality

## Path directory

SCION-Swarm allows to select a path when connecting to a node. This can be done by the flag _-path_ followed by the file path that contains the path dictionary.

The path dictionary is a file with the following structure.
```
<SourceAS>
<DestinationAS_1>;<DestinationPath_1_1>
<DestinationAS_2>;<DestinationPath_2_1>
```
For a sample path dictionary:
```
cat /home/xenon3/go/src/github.com/ethereum/go-ethereum/~scion/files/paths/paths.dct
```

## Connecting to the node via web3 console
It's possible to access the node and check it's peers via a web3 console. See [Ethereum Project](https://github.com/ethereum/go-ethereum) for further information.
```
cd $GOPATH/src/github.com/ethereum/go-ethereum/~scion/n2
geth attach bzzd.ipc # opens console
admin.peers
```

## A compilation for ARM is possible

See [go-ethereum wiki](https://github.com/ethereum/go-ethereum/wiki/Cross-compiling-Ethereum) for further information.

## Links

- [SCION Project](https://www.scion-architecture.net/)
- [SCION Tutorials](https://netsec-ethz.github.io/scion-tutorials/)
- [SCIONLab](https://www.scionlab.org/)

- [Ethereum Project](https://github.com/ethereum/go-ethereum)
- [Swarm Project](https://swarm-gateways.net/bzz:/theswarm.eth/)
- [Swarm Documentation](https://swarm-guide.readthedocs.io/en/latest/introduction.html)
