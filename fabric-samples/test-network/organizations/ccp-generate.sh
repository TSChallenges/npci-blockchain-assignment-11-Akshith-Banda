#!/usr/bin/env bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=1
P0PORT=7051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/tatamotors.example.com/tlsca/tlsca.tatamotors.example.com-cert.pem
CAPEM=organizations/peerOrganizations/tatamotors.example.com/ca/ca.tatamotors.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/tatamotors.example.com/connection-tatamotors.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/tatamotors.example.com/connection-tatamotors.yaml

ORG=2
P0PORT=9051
CAPORT=9054
PEERPEM=organizations/peerOrganizations/icici.example.com/tlsca/tlsca.icici.example.com-cert.pem
CAPEM=organizations/peerOrganizations/icici.example.com/ca/ca.icici.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/icici.example.com/connection-icici.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/icici.example.com/connection-icici.yaml

ORG=3
P0PORT=8051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/tesla.example.com/tlsca/tlsca.tesla.example.com-cert.pem
CAPEM=organizations/peerOrganizations/tesla.example.com/ca/ca.tesla.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/tesla.example.com/connection-tesla.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/tesla.example.com/connection-tesla.yaml

ORG=4
P0PORT=10051
CAPORT=10054
PEERPEM=organizations/peerOrganizations/chase.example.com/tlsca/tlsca.chase.example.com-cert.pem
CAPEM=organizations/peerOrganizations/chase.example.com/ca/ca.chase.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/chase.example.com/connection-chase.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/chase.example.com/connection-chase.yaml
