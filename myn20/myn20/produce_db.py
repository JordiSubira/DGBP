import yaml

'''
couchdb0:
    container_name: couchdb0
    image: hyperledger/fabric-couchdb
    # Populate the COUCHDB_USER and COUCHDB_PASSWORD to set an admin user and password
    # for CouchDB.  This will prevent CouchDB from operating in an "Admin Party" mode.
    environment:
      - COUCHDB_USER=
      - COUCHDB_PASSWORD=
    # Comment/Uncomment the port mapping if you want to hide/expose the CouchDB service,
    # for example map it to utilize Fauxton User Interface in dev environments.
    ports:
      - "5984:5984"
    networks:
      - byfn

  peer0.org1.example.com:
    environment:
      - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb0:5984
      # The CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME and CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD
      # provide the credentials for ledger to connect to CouchDB.  The username and password must
      # match the username and password set for the associated CouchDB.
      - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=
      - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=
    depends_on:
      - couchdb0

  couchdb1:
    container_name: couchdb1
    image: hyperledger/fabric-couchdb
    # Populate the COUCHDB_USER and COUCHDB_PASSWORD to set an admin user and password
    # for CouchDB.  This will prevent CouchDB from operating in an "Admin Party" mode.
    environment:
      - COUCHDB_USER=
      - COUCHDB_PASSWORD=
    # Comment/Uncomment the port mapping if you want to hide/expose the CouchDB service,
    # for example map it to utilize Fauxton User Interface in dev environments.
    ports:
      - "6984:5984"
    networks:
      - byfn

  peer1.org1.example.com:
    environment:
      - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb1:5984
      # The CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME and CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD
      # provide the credentials for ledger to connect to CouchDB.  The username and password must
      # match the username and password set for the associated CouchDB.
      - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=
      - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=
    depends_on:
      - couchdb1
'''

'''couchdb0 ={
		    'couchdb0' : { 
		    	'container_name': 'hyperledger/fabric-couchdb',
		    	'image' : 'hyperledger/fabric-couchdb',
		    	'enviroment' : ['COUCHDB_USER=', 
		    					'COUCHDB_PASSWORD='],
		    	'networks' : byfn,
		    	'ports' : ['5984:5984']
		    	}
		}
	couchdb1 ={
		    'couchdb1' : { 
		    	'container_name': 'hyperledger/fabric-couchdb',
		    	'image' : 'hyperledger/fabric-couchdb',
		    	'enviroment' : ['COUCHDB_USER=', 
		    					'COUCHDB_PASSWORD='],
		    	'networks' : byfn,
		    	'ports' : ['6984:5984']
		    	}
		}

	datapeer0 ={
	'peer0.org1.example.com' : { 
	'enviroment' : ['CORE_LEDGER_STATE_STATEDATABASE=CouchDB', 
					'CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb0:5984',
					'CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=',
					'CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD='],
	'depends_on' : [couchdb0]
	}
	datapeer1={
	'peer1.org1.example.com' : { 
	'enviroment' : ['CORE_LEDGER_STATE_STATEDATABASE=CouchDB', 
					'CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb1:5984',
					'CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=',
					'CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD='],
	'depends_on' : [couchdb1]
	}

'''
port_base = 5984
for i in range(0,10):
	peer0 = 'peer0.org' + str(i+1) + '.example.com'
	peer1 = 'peer1.org' + str(i+1) + '.example.com'
	port0 = port_base + (i)*2000
	port1 = port_base + (i)*2000 + 1000
	couchdb0 = 'couchdb' + str(i*2)
	couchdb1 = 'couchdb' + str(i*2+1)
	datapeer0 ={
	    peer0 : { 
			'environment' : ['CORE_LEDGER_STATE_STATEDATABASE=CouchDB', 
							'CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb'+ str(i*2) +':5984',
							'CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=',
							'CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD='],
			'depends_on' : [couchdb0]
		}
	}
	datapeer1 ={
	   	peer1 : { 
			'environment' : ['CORE_LEDGER_STATE_STATEDATABASE=CouchDB', 
							'CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb'+ str(i*2+1) +':5984',
							'CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=',
							'CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD='],
			'depends_on' : [couchdb1]
		}
	}

	couchdb0 ={
		    couchdb0: { 
		    	'container_name': couchdb0,
		    	'image' : 'hyperledger/fabric-couchdb',
		    	'environment' : ['COUCHDB_USER=', 
		    					'COUCHDB_PASSWORD='],
		    	'networks' : ['byfn'],
		    	'ports' : [str(port0)+ ':5984']
		    }
	}
	couchdb1 ={
		    couchdb1 : { 
		    	'container_name': couchdb1,
		    	'image' : 'hyperledger/fabric-couchdb',
		    	'environment' : ['COUCHDB_USER=', 
		    					'COUCHDB_PASSWORD='],
		    	'networks' : ['byfn'],
		    	'ports' : [str(port1) + ':5984']
		    	}
		}


	with open('db.yml', 'a') as outfile:
		yaml.dump(couchdb0, outfile, default_flow_style=False)
		yaml.dump(datapeer0, outfile, default_flow_style=False)
		yaml.dump(couchdb1, outfile, default_flow_style=False)
		yaml.dump(datapeer1, outfile, default_flow_style=False)