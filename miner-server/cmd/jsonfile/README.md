# 1.About chips.json

## report chip info uploading to the chain
its purpose is to upload all the real burned chip with its information of serial number, bus id, computation, public key and p2 key, uploaded to the chain, rpc method "reportChip" is used. It requires json file provided by the input ReportChipRequest.ChipFilePath. \
For a near super account, he will have to EXECUTE the rpc "reportChip" with its private key and public key provided in a json file "unc.json".\
--signer-public-key ed25519:xxx is the public key of the super account, --signer-private-key ed25519:xxx is the private key of the super account to do a signature

# 2.About miner_key.json
its purpose is to create a file for containing miner public key and chip public key, so that a miner binds with a chip to claim its computation, rpc method "ClaimChipComputation" is used

## challenge_key
this is the public key of miner, obtained by the rpc func GetMinerKeys, with its record of "private.pem" and "public.pem". It is based on ED25519 algorithm generation

## public_key
this is the public key of the chip

## generate signature for claim computation by a miner (near js)
"unc extensions create-challenge-rsa minerAccount use-file /path/to/miner_key.json without-init-call network-config my-private-chain-id sign-with-plaintext-private-key --signer-public-key ed25519:xxx --signer-private-key ed25519:xxx" \
minerAccount is the account id(name) for the miner, miner_key.json is the json file that possess challenge_key and public_key, allowing the near node to read the keys message, with --signer-public-key ed25519:xxx as the public key of miner account, --signer-private-key ed25519:xxx as the private key of miner account
it will generate a signature: "Signature: ed25519:...", then display it, and obtain the sign transaction S. Using the ClaimChipComputation rpc, insert the value S into the req.Signature, and send transaction.

# 3. About validator_key.json

## account_id
minerAccount name

## public_key
this is the public key of the miner (public.pem)

## public_key
this is the private key of the miner (private.pem)

## stake by a miner (near js)
"near stake minerAccount ed25519:xxx amount --keyPath /path/to/validator_key.json" \
miner will stake certain amount of token in order to activate the mining
