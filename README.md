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


deploy contract

```cassandraql
tcpcli tcp deploy --conAddress tcp1upg6v5g7vvcdm7uxay2c43hz9k0ap0vmazft5s --code aaaa --codeHash aaaa --from tcp1gp8t6r9znpnqsj7k54t9mpkafxcwslv5ld4u39

```

contract exec

```cassandraql
tcpcli tcp exec --conAddress tcp1upg6v5g7vvcdm7uxay2c43hz9k0ap0vmazft5s --callAddress tcp1gp8t6r9znpnqsj7k54t9mpkafxcwslv5ld4u39 --state aaaa --proof aaaa --resultHash aaaa --from tcp1gp8t6r9znpnqsj7k54t9mpkafxcwslv5ld4u39

```



start the rest-server in another terminal window:
```cassandraql
tcpcli rest-server --chain-id tcp-chain --trust-node
```

query account
```cassandraql
curl -s http://127.0.0.1:1317/auth/accounts/$(tcpcli keys show jack -a)
```


query contract code
```cassandraql

curl -s http://127.0.0.1:1317/tcp/contracts/tcp1upg6v5g7vvcdm7uxay2c43hz9k0ap0vmazft5s
```

