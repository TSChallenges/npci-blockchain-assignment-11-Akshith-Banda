export FABRIC_CFG_PATH=/workspaces/npci-blockchain-assignment-11-Akshith-Banda/fabric-samples/config

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="TataMotorsMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/workspaces/npci-blockchain-assignment-11-Akshith-Banda/fabric-samples/test-network/organizations/peerOrganizations/tatamotors.example.com/peers/peer0.tatamotors.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/workspaces/npci-blockchain-assignment-11-Akshith-Banda/fabric-samples/test-network/organizations/peerOrganizations/tatamotors.example.com/users/Admin@tatamotors.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer chaincode query -C locchannel -n loc -c '{"Args":["GetLOCHistory", "loc123"]}'