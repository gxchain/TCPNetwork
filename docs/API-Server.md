# API Server

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