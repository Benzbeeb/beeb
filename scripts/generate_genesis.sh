DIR=`dirname "$0"`

rm -rf ~/.beeb

# initial new node
beebd init validator --chain-id beebchain
echo "lock nasty suffer dirt dream fine fall deal curtain plate husband sound tower mom crew crawl guard rack snake before fragile course bacon range" \
    | beebd keys add validator --recover --keyring-backend test
echo "smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic" \
    | beebd keys add requester --recover --keyring-backend test

# cp ./docker-config/single-validator/priv_validator_key.json ~/.beeb/config/priv_validator_key.json
# cp ./docker-config/single-validator/node_key.json ~/.beeb/config/node_key.json

# add accounts to genesis
beebd add-genesis-account validator 10000000000000stake --keyring-backend test
beebd add-genesis-account requester 10000000000000stake --keyring-backend test


# register initial validators
beebd gentx validator 100000000stake \
    --chain-id beebchain \
    --keyring-backend test

# collect genesis transactions
beebd collect-gentxs


