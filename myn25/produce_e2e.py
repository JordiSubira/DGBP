import yaml

'''datapeer0 ={
		    'peer0.org1.example.com' : { 
		    	'container_name': 'peer0.org1.example.com',
		    	
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

peer0.org1.example.com:
    container_name: peer0.org1.example.com
    extends:
      file:  base/docker-compose-base.yaml
      service: peer0.org1.example.com
    networks:
      - byfn
'''
port_base = 7051
for i in range(21,26):
	peer0 = 'peer0.org' + str(i) + '.example.com'
	port0 = port_base + (i-1)*1000
	datapeer0 ={
	    peer0 : { 
	    	'container_name': peer0,
	    	'extends':{ 'file': 'base/docker-compose-base.yaml', 'service': peer0},
	    	'networks' : ['byfn']
	    	}
	}
	
	with open('data_e2e.yml', 'a') as outfile:
		yaml.dump(datapeer0, outfile, default_flow_style=False)