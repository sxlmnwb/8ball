## get a zclassic raw transaction data by hash from the command line
./zclassic-cli getrawtransaction "76ce9890d1855896ca75f7a3c1f5b3e48b09dc40272418432d68ec9ee22d8d81" 1
{
  "hex": "0400008085202f8901710e10d8b43e047220cb38e7974d8f8b1322fc200728b8a7f2153b84b381bf56010000006a473044022006e190ef2b1830f7fdaf14e65e710dff5965644f3bd8568c13bd78c2753f259402207a9313f74f4dfd5673d15afe52775c735263e18e23765ff77e0c2ef2ea9f7eb701210283f85460b559bdfc43d76cc9ce86a09329a0ca04175d8e950ef47cdfe817a2a7feffffff0200e1f505000000001976a914ad0fb2ce904fb7ac69a2760747650a4b6e201fd288ac0b83d717000000001976a9143eafa1453bdc3543e366969c5401d889220dba6188ac80d21a00fdd21a000000000000000000000000",
  "txid": "76ce9890d1855896ca75f7a3c1f5b3e48b09dc40272418432d68ec9ee22d8d81",
  "overwintered": true,
  "version": 4,
  "versiongroupid": "892f2085",
  "locktime": 1757824,
  "expiryheight": 1757949,
  "vin": [
    {
      "txid": "56bf81b3843b15f2a7b8280720fc22138b8f4d97e738cb2072043eb4d8100e71",
      "vout": 1,
      "scriptSig": {
        "asm": "3044022006e190ef2b1830f7fdaf14e65e710dff5965644f3bd8568c13bd78c2753f259402207a9313f74f4dfd5673d15afe52775c735263e18e23765ff77e0c2ef2ea9f7eb7[ALL] 0283f85460b559bdfc43d76cc9ce86a09329a0ca04175d8e950ef47cdfe817a2a7",
        "hex": "473044022006e190ef2b1830f7fdaf14e65e710dff5965644f3bd8568c13bd78c2753f259402207a9313f74f4dfd5673d15afe52775c735263e18e23765ff77e0c2ef2ea9f7eb701210283f85460b559bdfc43d76cc9ce86a09329a0ca04175d8e950ef47cdfe817a2a7"
      },
      "sequence": 4294967294
    }
  ],
  "vout": [
    {
      "value": 1.00000000,
      "valueZat": 100000000,
      "n": 0,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 ad0fb2ce904fb7ac69a2760747650a4b6e201fd2 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a914ad0fb2ce904fb7ac69a2760747650a4b6e201fd288ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "t1ZefiGenesisBootstrapBURNxxxyfs71k"
        ]
      }
    },
    {
      "value": 3.99999755,
      "valueZat": 399999755,
      "n": 1,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 3eafa1453bdc3543e366969c5401d889220dba61 OP_EQUALVERIFY OP_CHECKSIG",
        "hex": "76a9143eafa1453bdc3543e366969c5401d889220dba6188ac",
        "reqSigs": 1,
        "type": "pubkeyhash",
        "addresses": [
          "t1Pb4JDJ8A2sNTVkSHbVes5kzXYx5RRwzh7"
        ]
      }
    }
  ],
  "vjoinsplit": [
  ],
  "valueBalance": 0.00000000,
  "vShieldedSpend": [
  ],
  "vShieldedOutput": [
  ],
  "blockhash": "000016e4e0ff1bf49d32bd2c6f9b6e29d86c5b05657a9835f7b37b94ca53695a",
  "confirmations": 55005,
  "time": 1666272110,
  "blocktime": 1666272110
}

