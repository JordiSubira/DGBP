'''
elif [ $ORG -eq 10 ] ; then
		CORE_PEER_LOCALMSPID="Org10MSP"
		CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/peers/peer0.org10.example.com/tls/ca.crt
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org10.example.com/users/Admin@org10.example.com/msp
		if [ $PEER -eq 0 ]; then
			CORE_PEER_ADDRESS=peer0.org10.example.com:7051
		else
			CORE_PEER_ADDRESS=peer1.org10.example.com:7051
		fi
'''

for i in range(13,21):
	l1 = 'elif [ $ORG -eq ' + str(i) + ' ] ; then\n'
	l2 = '\t CORE_PEER_LOCALMSPID=\"Org' + str(i) + 'MSP\"\n'
	l3 = '\t CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org'+ str(i) +'.example.com/users/Admin@org'+ str(i) + '.example.com/msp\n'
	l4 = '\t CORE_PEER_ADDRESS=peer0.org' + str(i) +'.example.com:7051\n'
	with open('data_script.sh', 'a') as outfile:
		outfile.write(l1)
		outfile.write(l2)
		outfile.write(l3)
		outfile.write(l4)
