# 1.About miner_key.json

## report chip info uploading to the chain 
for a near super account, he will have to EXECUTE "near extensions register-rsa-keys superAccount use-file /path/to/signer0_key.json with-init-call json-args '{"power": "6000000000000"}' network-config my-private-chain-id sign-with-plaintext-private-key --signer-public-key ed25519:xxx --signer-private-key ed25519:xxx send" TO upload the chip information on the chain \
Here the signer0_key.json is the json file to be read by the near command, '{"power": "6000000000000"}' is defined how much computation a chip possesses.\
--signer-public-key ed25519:xxx is the public key of the super account, --signer-private-key ed25519:xxx is the private key of the super account to do a signature

## challenge_key
this is the public key of miner, obtained by the rpc func GetMinerKeys, with its record of "private.pem" and "public.pem". It is based on ED25519 algorithm generation

## public_key
this is the public key of the chip

## generate signature for claim computation by a miner
"near extensions create-challenge-rsa minerAccount use-file /path/to/miner_key.json without-init-call network-config my-private-chain-id sign-with-plaintext-private-key --signer-public-key ed25519:xxx --signer-private-key ed25519:xxx" \
miner_key.json is the json file that possess challenge_key and public_key, allowing the near node to read the keys message, with --signer-public-key ed25519:xxx as the public key of miner account, --signer-private-key ed25519:xxx as the private key of miner account
it will generate a signature: "Signature: ed25519:...", then display it, and obtain the sign transaction S. Using the ClaimChipComputation rpc, insert the value S into the req.Signature, and send transaction.

# 2. About validator_key.json

## account_id
minerAccount name

## public_key
this is the public key of the miner (public.pem)

## public_key
this is the private key of the miner (private.pem)

## stake by a miner
"near stake minerAccount ed25519:xxx amount --keyPath /path/to/validator_key.json" \
miner will stake certain amount of token in order to activate the mining
