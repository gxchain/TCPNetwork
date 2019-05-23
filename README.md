# TCP network 
TCP network is an efficient trusted computing network.


## QuickStart

### Build
set env
```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=\$GOPATH/bin" >> ~/.bash_profile
echo "export PATH=\$PATH:\$GOBIN" >> ~/.bash_profile
echo "export GO111MODULE=on" >> ~/.bash_profile
source ~/.bash_profile
```

build
```bash
# get source code
git clone https://github.com/gxchain/TCPNetwork.git


# Install the app into your $GOBIN
make install

# Now you should be able to run the following commands:
tcpd help
tcpcli help

```


### Run

init
```
# Initialize configuration files and genesis file
tcpd init --chain-id tcp-chain


# Copy the `Address` output here and save it for later use 
tcpcli keys add jack


# Copy the `Address` output here and save it for later use
tcpcli keys add alice

# Add both accounts, with coins to the genesis file
tcpd add-genesis-account $(tcpcli keys show jack -a) 1000nametoken,1000jackcoin
tcpd add-genesis-account $(tcpcli keys show alice -a) 1000nametoken,1000alicecoin

# Configure your CLI to eliminate need for chain-id flag
tcpcli config chain-id tcp-chain
tcpcli config output json
tcpcli config indent true
tcpcli config trust-node true
```

run tcpd

```cassandraql
tcpd start --log_level "*:debug" --trace

```
run tcpcli
```cassandraql
# query account
tcpcli query account $(tcpcli keys show jack -a) 

tcpcli query account $(tcpcli keys show alice -a) 


# transfer asset
 tcpcli transfer --from tcp1vd3afehmxan7tsvqwzk5kwj7k8gh2pghgupv8z --to tcp13skcx3v2kc4p5zkxxy73gpqq75e6g4jgelzc7f --amount 1jackcoin 

```






