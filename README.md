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
tcpcli tcp deploy --conAddress tcp1rd0cy8wtnxgy40nksjw06refgzzmvelwxdmv86 --code "code-x" --codeHash aaaa  --from tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h --dataSources "tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h,tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h" --targets "tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h,tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h"
```

response:
```
zhuliting-2:TCPNetwork zhuliting$ tcpcli tcp deploy --conAddress tcp1rd0cy8wtnxgy40nksjw06refgzzmvelwxdmv86 --code "code-x" --codeHash aaaa  --from tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h --dataSources "tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h,tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h" --targets "tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h,tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h"
{"chain_id":"tcp-chain","account_number":"0","sequence":"3","fee":{"amount":null,"gas":"200000"},"msgs":[{"type":"tcp/deploy","value":{"From":"tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h","CID":"tcp1rd0cy8wtnxgy40nksjw06refgzzmvelwxdmv86","Targets":["tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h","tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h"],"DataSources":["tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h","tcp16h88dtlvz84ghh6q5kj2pfnjp3epmsx6t3k08h"],"Code":"Y29kZS14","CodeHash":"YWFhYQ==","State":"AA==","Fee":[{"denom":"nametoken","amount":"10"}]}}],"memo":""}

confirm transaction before signing and broadcasting [Y/n]: y
Password to sign with 'jack':
{
 "height": "709",
 "txhash": "D5D766DA1F24721F9E72FAC41E9FDFD855DD3CF38EA3BE800E69C7931CFC995F",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "25327",
 "tags": [
  {
   "key": "action",
   "value": "tcp_deploy"
  }
 ]
}
```

contract exec

```cassandraql
tcpcli tcp exec --conAddress tcp1upg6v5g7vvcdm7uxay2c43hz9k0ap0vmazft5s --callAddress tcp1gp8t6r9znpnqsj7k54t9mpkafxcwslv5ld4u39 --state aaaa --proof aaaa --resultHash aaaa --from tcp1gp8t6r9znpnqsj7k54t9mpkafxcwslv5ld4u39

```
response:
```
zhuliting-2:gxb-core zhuliting$ tcpcli tcp exec --conAddress tcp1rd0cy8wtnxgy40nksjw06refgzzmvelwxdmv86 --callAddress tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk --state "state" --proof "proof" --resultHash "result-hash" --from tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk
{"chain_id":"tcp-chain","account_number":"0","sequence":"1","fee":{"amount":null,"gas":"200000"},"msgs":[{"type":"tcp/exec","value":{"From":"tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk","CID":"tcp1rd0cy8wtnxgy40nksjw06refgzzmvelwxdmv86","State":"c3RhdGU=","RequestParam":{"from":"tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk","cid":"tcp1rd0cy8wtnxgy40nksjw06refgzzmvelwxdmv86","proxy":"tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk","dataSource":null,"fee":null,"signature":null},"Proof":"cHJvb2Y=","ResultHash":"cmVzdWx0LWhhc2g=","Fee":[{"denom":"nametoken","amount":"1"}]}}],"memo":""}

confirm transaction before signing and broadcasting [Y/n]: y
Password to sign with 'jack':
{
 "height": "51",
 "txhash": "757FB1FFD24D869C8F3F2263373CA49DDB559FFA16C89F8DE11939D7062732AB",
 "data": "cmVzdWx0LWhhc2g=",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "41069",
 "tags": [
  {
   "key": "action",
   "value": "tcp_exec"
  }
 ]
}

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

