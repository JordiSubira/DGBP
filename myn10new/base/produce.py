import yaml

'''datapeer0 ={
		    'peer0.org1.example.com' : { 
		    	'container_name': 'orderer.example.com',
		    	'image' : 'hyperledger/fabric-orderer:$IMAGE_TAG',
		    	'enviroment' : ['CORE_PEER_ID=peer0.org1.example.com', 
		    					'CORE_PEER_ADDRESS=peer0.org1.example.com:7051',
		    					'CORE_PEER_GOSSIP_BOOTSTRAP=peer1.org1.example.com:7051',
		    					'CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org1.example.com:7051',
		    					'CORE_PEER_LOCALMSPID=Org1MSP'],
		    	'volumes' : ['/var/run/:/host/var/run/', 
		    					'../crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp:/etc/hyperledger/fabric/msp',
		    					'peer0.org1.example.com:/var/hyperledger/production'],
		    	'ports' : ['- 7051:7051', '7053:7053']
		    	}
		}

datapeer1 ={
    'peer1.org1.example.com' : { 
    	'container_name': 'orderer.example.com',
    	'image' : 'hyperledger/fabric-orderer:$IMAGE_TAG',
    	'enviroment' : ['CORE_PEER_ID=peer1.org1.example.com', 
    					'CORE_PEER_ADDRESS=peer1.org1.example.com:7051',
    					'CORE_PEER_GOSSIP_BOOTSTRAP=peer0.org1.example.com:7051',
    					'CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer1.org1.example.com:7051',
    					'CORE_PEER_LOCALMSPID=Org1MSP'],
    	'volumes' : ['/var/run/:/host/var/run/', 
    					'../crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/msp:/etc/hyperledger/fabric/msp',
    					'peer1.org1.example.com:/var/hyperledger/production'],
    	'ports' :	['- 8051:7051', '8053:7053']
    }
}'''
port_base = 7051
for i in range(1,21):
	peer0 = 'peer0.org' + str(i) + '.example.com'
	port0 = port_base + (i-1)*1000
	datapeer0 ={
	    peer0 : { 
	    	'container_name': peer0,
	    	'extends':{ 'file': 'peer-base.yaml', 'service': 'peer-base'},
	    	'environment' : ['CORE_PEER_ID=' + peer0, 
	    					'CORE_PEER_ADDRESS=' + peer0 +':7051',
	    					#'CORE_PEER_GOSSIP_BOOTSTRAP=' + peer1 +':7051',
	    					'CORE_PEER_GOSSIP_EXTERNALENDPOINT=' + peer0 + ':7051',
	    					'CORE_PEER_LOCALMSPID=Org'+ str(i) +'MSP'],
	    	'volumes' : ['/var/run/:/host/var/run/', 
	    					'../crypto-config/peerOrganizations/org'+str(i) +'.example.com/peers/'+ peer0 +'/msp:/etc/hyperledger/fabric/msp',
	    					peer0 +':/var/hyperledger/production'],
	    	'ports' : [ str(port0)+':7051', str(port0+2) +':7053']
	    	}
	}
	
	with open('data.yml', 'a') as outfile:
		yaml.dump(datapeer0, outfile, default_flow_style=False)