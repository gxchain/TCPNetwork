# API Server
使用 tcpcli的RPC API， 需要先开启rest-server
```
tcpcli rest-server --chain-id tcp-chain --trust-node --laddr tcp://0.0.0.0:1317
```


### get account

request:
```cassandraql
curl -s http://127.0.0.1:1317/auth/accounts/tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk
```

response:
```cassandraql
{
  "type": "auth/Account",
  "value": {
    "address": "tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk",
    "coins": [
      {
        "denom": "jackcoin",
        "amount": "1000"
      },
      {
        "denom": "nametoken",
        "amount": "989"
      }
    ],
    "public_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "A2oPOFwPrHrsdi6FEER05E9Cby3uxpOxIGbKnjwwKMUD"
    },
    "account_number": "0",
    "sequence": "2"
  }
}
```

### get account balance

request:
```cassandraql
 curl -s http://127.0.0.1:1317/bank/balances/tcp1zxsps8kjj4eym6su8e34ny9efrd6cgwag8u0k6
 ```

response:
```cassandraql
[
  {
    "denom": "nametoken",
    "amount": "100000000"
  }
]
```

### get contract code
request:
 ```cassandraql
curl -s http://127.0.0.1:1317/tcp/contracts/tcp1rd0cy8wtnxgy40nksjw06refgzzmvelwxdmv86
```

response:
```cassandraql
{
  "value": "code-x"
}
```
