

## query API
### query genesis

request:
```cassandraql
curl http://127.0.0.1:26657/genesis
```

reponse:
```cassandraql
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "genesis": {
      "genesis_time": "2019-06-04T07:30:29.132985Z",
      "chain_id": "tcp-chain",
      "consensus_params": {
        "block_size": {
          "max_bytes": "22020096",
          "max_gas": "-1"
        },
        "evidence": {
          "max_age": "100000"
        },
        "validator": {
          "pub_key_types": [
            "ed25519"
          ]
        }
      },
      "validators": [
        {
          "address": "8A3A62EFB534F2BB08236BC819A26F03009B12E9",
          "pub_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "jQ8tSIVms5HXpwiOZS5Y029XjYLbomDNFNF2yojUidg="
          },
          "power": "10",
          "name": ""
        }
      ],
      "app_hash": "",
      "app_state": {
        "auth": {
          "collected_fees": null,
          "params": {
            "max_memo_characters": "256",
            "tx_sig_limit": "7",
            "tx_size_cost_per_byte": "10",
            "sig_verify_cost_ed25519": "590",
            "sig_verify_cost_secp256k1": "1000"
          }
        },
        "bank": {
          "send_enabled": true
        },
        "accounts": [
          {
            "address": "tcp15juugyz27ldj535ntajp226vwqrnvk0r4w03xk",
            "coins": [
              {
                "denom": "jackcoin",
                "amount": "1000"
              },
              {
                "denom": "nametoken",
                "amount": "1000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          },
          {
            "address": "tcp13x4w47e0n05dmxj5fd5nf7u3jk9chwtpmvjt2v",
            "coins": [
              {
                "denom": "alicecoin",
                "amount": "1000"
              },
              {
                "denom": "nametoken",
                "amount": "1000"
              }
            ],
            "public_key": null,
            "account_number": "0",
            "sequence": "0"
          }
        ]
      }
    }
  }
}
```


### query status
request
```cassandraql
curl http://127.0.0.1:26657/status
```

response:
```cassandraql
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "node_info": {
      "protocol_version": {
        "p2p": "7",
        "block": "10",
        "app": "0"
      },
      "id": "be49b51e208ce9f59c91679b98d7056e92fc502d",
      "listen_addr": "tcp://0.0.0.0:26656",
      "network": "tcp-chain",
      "version": "0.30.1",
      "channels": "4020212223303800",
      "moniker": "zhuliting-2.local",
      "other": {
        "tx_index": "on",
        "rpc_address": "tcp://0.0.0.0:26657"
      }
    },
    "sync_info": {
      "latest_block_hash": "1F878F5DA7AA0F9A706D8B050F182C13D60B57484378C1E2046ED41E9FF15CCA",
      "latest_app_hash": "D44F45F81DFD789EFE27CAC4D62A3499DD2DF24E4F0A5B13C73EA34C04F5D44B",
      "latest_block_height": "114",
      "latest_block_time": "2019-06-04T08:18:13.575539Z",
      "catching_up": false
    },
    "validator_info": {
      "address": "8A3A62EFB534F2BB08236BC819A26F03009B12E9",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "jQ8tSIVms5HXpwiOZS5Y029XjYLbomDNFNF2yojUidg="
      },
      "voting_power": "10"
    }
  }
}
```


### query block
request:
```cassandraql
curl http://127.0.0.1:26657/block?height=51
```

response:
```cassandraql
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "block_meta": {
      "block_id": {
        "hash": "7B62537258F046AC3F358E07BC0B9A6BCB74EA36A455115D91121BD5C4CA9391",
        "parts": {
          "total": "1",
          "hash": "B747E8C9FDA952700DBAAAA42E4159A1DEBCF62C043A987AEAE96255BDE181A8"
        }
      },
      "header": {
        "version": {
          "block": "10",
          "app": "0"
        },
        "chain_id": "tcp-chain",
        "height": "51",
        "time": "2019-06-04T07:34:52.085219Z",
        "num_txs": "1",
        "total_txs": "2",
        "last_block_id": {
          "hash": "594381DDE900FE7D667F6B0F127979EBDF71FE7A25A218BDBB3869A7509F9ADE",
          "parts": {
            "total": "1",
            "hash": "A5CB13BFCF296FDF9BC12B900690187E581C2CBFFC4932C2832E85A61D844DB9"
          }
        },
        "last_commit_hash": "B4612CD8369E6B350764730338414807D3C0AE8216B3A707B9A505AA087F2252",
        "data_hash": "BCD76FC51570812A6FDF1C56FC4DF8F02B7175D67E9310DAD31D5975D78ED55E",
        "validators_hash": "A08D372AF957547441DE1A47BF466EA0C945C86B18BD46EE4855264880077873",
        "next_validators_hash": "A08D372AF957547441DE1A47BF466EA0C945C86B18BD46EE4855264880077873",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "A6FCE6F46AAEB891221B04386C9B0798153200C9B5E27A0865EDB97A84C4A90A",
        "last_results_hash": "",
        "evidence_hash": "",
        "proposer_address": "8A3A62EFB534F2BB08236BC819A26F03009B12E9"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "10",
          "app": "0"
        },
        "chain_id": "tcp-chain",
        "height": "51",
        "time": "2019-06-04T07:34:52.085219Z",
        "num_txs": "1",
        "total_txs": "2",
        "last_block_id": {
          "hash": "594381DDE900FE7D667F6B0F127979EBDF71FE7A25A218BDBB3869A7509F9ADE",
          "parts": {
            "total": "1",
            "hash": "A5CB13BFCF296FDF9BC12B900690187E581C2CBFFC4932C2832E85A61D844DB9"
          }
        },
        "last_commit_hash": "B4612CD8369E6B350764730338414807D3C0AE8216B3A707B9A505AA087F2252",
        "data_hash": "BCD76FC51570812A6FDF1C56FC4DF8F02B7175D67E9310DAD31D5975D78ED55E",
        "validators_hash": "A08D372AF957547441DE1A47BF466EA0C945C86B18BD46EE4855264880077873",
        "next_validators_hash": "A08D372AF957547441DE1A47BF466EA0C945C86B18BD46EE4855264880077873",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "A6FCE6F46AAEB891221B04386C9B0798153200C9B5E27A0865EDB97A84C4A90A",
        "last_results_hash": "",
        "evidence_hash": "",
        "proposer_address": "8A3A62EFB534F2BB08236BC819A26F03009B12E9"
      },
      "data": {
        "txs": [
          "mALwYl3uCp8BHRISWgoUpLnEEEr32ypGk19kFStMcAc2WeMSFBtfgh3LmZBKvnaEnP0PKUCFtmfuGgVzdGF0ZSJCChSkucQQSvfbKkaTX2QVK0xwBzZZ4xIUG1+CHcuZkEq+doSc/Q8pQIW2Z+4aFKS5xBBK99sqRpNfZBUrTHAHNlnjKgVwcm9vZjILcmVzdWx0LWhhc2g6DgoJbmFtZXRva2VuEgExEgQQwJoMGmoKJuta6YchA2oPOFwPrHrsdi6FEER05E9Cby3uxpOxIGbKnjwwKMUDEkBIGRq+mm3CwtCT7xOoKB41h9VbS26x8jCWGbKHGKpavEdxgVfQGLoe/5yfQsPn4pFw5Nw3uLy7yyZ3YHOSjcJ2"
        ]
      },
      "evidence": {
        "evidence": null
      },
      "last_commit": {
        "block_id": {
          "hash": "594381DDE900FE7D667F6B0F127979EBDF71FE7A25A218BDBB3869A7509F9ADE",
          "parts": {
            "total": "1",
            "hash": "A5CB13BFCF296FDF9BC12B900690187E581C2CBFFC4932C2832E85A61D844DB9"
          }
        },
        "precommits": [
          {
            "type": 2,
            "height": "50",
            "round": "0",
            "block_id": {
              "hash": "594381DDE900FE7D667F6B0F127979EBDF71FE7A25A218BDBB3869A7509F9ADE",
              "parts": {
                "total": "1",
                "hash": "A5CB13BFCF296FDF9BC12B900690187E581C2CBFFC4932C2832E85A61D844DB9"
              }
            },
            "timestamp": "2019-06-04T07:34:52.085219Z",
            "validator_address": "8A3A62EFB534F2BB08236BC819A26F03009B12E9",
            "validator_index": "0",
            "signature": "XsRYtVyiawWrMBVAJz7uCHk7FKgQIbh1Q4hyVxPqBUa7uJ2kUy+9Ni59d+pHhnHOTT5VP4cECVJnrKbY0qXTAg=="
          }
        ]
      }
    }
  }
}
```


### query tx
request:
```cassandraql
curl http://127.0.0.1:26657/tx?hash=0x757FB1FFD24D869C8F3F2263373CA49DDB559FFA16C89F8DE11939D7062732AB
```

response:
```cassandraql
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "hash": "757FB1FFD24D869C8F3F2263373CA49DDB559FFA16C89F8DE11939D7062732AB",
    "height": "51",
    "index": 0,
    "tx_result": {
      "data": "cmVzdWx0LWhhc2g=",
      "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
      "gasWanted": "200000",
      "gasUsed": "41069",
      "tags": [
        {
          "key": "YWN0aW9u",
          "value": "dGNwX2V4ZWM="
        }
      ]
    },
    "tx": "mALwYl3uCp8BHRISWgoUpLnEEEr32ypGk19kFStMcAc2WeMSFBtfgh3LmZBKvnaEnP0PKUCFtmfuGgVzdGF0ZSJCChSkucQQSvfbKkaTX2QVK0xwBzZZ4xIUG1+CHcuZkEq+doSc/Q8pQIW2Z+4aFKS5xBBK99sqRpNfZBUrTHAHNlnjKgVwcm9vZjILcmVzdWx0LWhhc2g6DgoJbmFtZXRva2VuEgExEgQQwJoMGmoKJuta6YchA2oPOFwPrHrsdi6FEER05E9Cby3uxpOxIGbKnjwwKMUDEkBIGRq+mm3CwtCT7xOoKB41h9VbS26x8jCWGbKHGKpavEdxgVfQGLoe/5yfQsPn4pFw5Nw3uLy7yyZ3YHOSjcJ2"
  }
}
```


## tx API
###  BroadcastTxAsync
request:
```cassandraql
curl 'http://127.0.0.1:26657/broadcast_tx_async?tx="0xdb01f0625dee0a63ce6dc0430a14813e4939f1567b219704ffc2ad4df58bde010879122b383133453439333946313536374232313937303446464332414434444635384244453031303837392d34341a0d5a454252412d3136445f424e422002280130c0843d38904e400112700a26eb5ae9872102139bdd95de72c22ac2a2b0f87853b1cca2e8adf9c58a4a689c75d3263013441a124015e99f7a686529c76ccc2d70b404af82ca88dfee27c363439b91ea0280571b2731c03b902193d6a5793baf64b54bcdf3f85e0d7cf657e1a1077f88143a5a65f518d2e518202b"'
```

response:
```cassandraql
zhuliting-2:gxb-core zhuliting$ curl 'http://127.0.0.1:26657/broadcast_tx_async?tx="0xdb01f0625dee0a63ce6dc0430a14813e4939f1567b219704ffc2ad4df58bde010879122b383133453439333946313536374232313937303446464332414434444635384244453031303837392d34341a0d5a454252412d3136445f424e422002280130c0843d38904e400112700a26eb5ae9872102139bdd95de72c22ac2a2b0f87853b1cca2e8adf9c58a4a689c75d3263013441a124015e99f7a686529c76ccc2d70b404af82ca88dfee27c363439b91ea0280571b2731c03b902193d6a5793baf64b54bcdf3f85e0d7cf657e1a1077f88143a5a65f518d2e518202b"'
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "code": 0,
    "data": "",
    "log": "",
    "hash": "4A611BCF9C5F905CF95961CB70DCFEEF237A3E9EB763F477AFA7DB24B7B52139"
  }
}
```

### BroadcastTxSync
request:
```cassandraql
curl 'http://127.0.0.1:26657/broadcast_tx_sync?tx="0xdb01f0625dee0a63ce6dc0430a14813e49x39f1567b219704ffc2ad4df58bde010879122b383133453439333946313536374232313937303446464332414434444635384244453031303837392d34381a0d5a454252412d3136445f424e422002280130c0843d38904e400112700a26eb5ae9872102139bdd95de72c22ac2a2b0f87853b1cca2e8adf9c58a4a689c75d3263013441a12406032dd568bac76ef8231fdf928f663ab6893124465528cc8ac5232afdceceea41640227501847c95dc5307f9bbcd01c82b33093c0eb11af8aef9c70eeb554f9318d2e518202f"'

```

response:
```cassandraql
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "code": 0,
    "data": "7B226F726465725F6964223A22383133453439333946313536374232313937303446464332414434444635384244453031303837392D3438227D",
    "log": "Msg 0: ",
    "hash": "920EA6B3EE38AC9B700AB436DABCA8F3D97F06EA63CBCACA7AD22B2E5CA1DF75"
  }
}
```