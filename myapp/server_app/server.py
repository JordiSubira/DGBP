import sys
import subprocess as sp
import string
import json
#import time

RET_KO_POLICY=255
RET_KO_USER=254
RET_OK_POLICY=0

toOrg = 'Org1MSP'

def queryUserDepChaincodeByEID(eid):
	#llamar a queryUser.js con certificado
	proc = sp.Popen( ['node', 'queryUser.js' , eid], stdout=sp.PIPE, stderr=sp.PIPE)
	output, error = proc.communicate()
	print "Output: ",output
	print "Error: ", error
	#if error!='' or output=='':
	#	return None, None

	try:
		j = json.loads(output)
		msp = j['dep']['msp']
		dep = j['dep']['name']
		print dep, msp
		return str(dep), str(msp)
	except ValueError:
		print "Error decoding JSON from HyperLedger"
		return None, None

def queryPolicyChaincode(from_dep,from_msp,to_eid,to_msp):
	proc = sp.Popen( ['node', 'queryPolicy.js' , to_msp, to_eid, from_msp, from_dep], stdout=sp.PIPE, stderr=sp.PIPE)
	output, error = proc.communicate()
	print "Output: ",output
	print "Error: ", error
	return error == ''

if __name__ == "__main__":

	#check len(argv)
	eid_a = sys.argv[1]
	eid_b = sys.argv[2]
 
	#queryUser to Chaincode
	dep, msp = queryUserDepChaincodeByEID(str(eid_a))
	if dep == None:
		print "User EID not found = ",eid_a 
		sys.exit(RET_KO_USER)

	#query if policy exists between departments
	if queryPolicyChaincode(dep,msp,eid_b,toOrg) == False:
		print "Policy doesn't exist"
		sys.exit(RET_KO_POLICY)

	print "Policy does exist"
	sys.exit(RET_OK_POLICY)
