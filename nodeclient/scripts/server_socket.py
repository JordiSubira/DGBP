import socket
import sys
import subprocess as sp
from fastecdsa import curve, ecdsa, keys
from hashlib import sha384
import string
import random
import json
import time

toOrg = 'Org1MSP'
toDep = 'Dep1'

def checkNonce(file, m,r,s):
	_ , public_key = keys.import_key(file)
	print("public key: ",public_key)

	# should return True as the signature we just generated is valid.
	valid = ecdsa.verify( ( r, s ), m, public_key)
	print("Is it valid? = ",valid)
	return valid

def generateNonce(length):
	return ''.join(random.choice(string.ascii_letters + string.digits + '-' + '.' + '_' + '~') for i in range(length))

def queryUserDepChaincodeByPKI(pub_key):
	#llamar a queryUser.js con certificado
	proc = sp.Popen( ['node', '../queryUser.js' , pub_key], stdout=sp.PIPE, stderr=sp.PIPE)
	output, error = proc.communicate()
	print("Output: ",output)
	print("Error: ", error)
	j = json.loads(output)
	msp = j['dep']['msp']
	dep = j['dep']['name']
	print(dep, msp)
	return str(dep), str(msp)

def queryUserDepChaincodeByEID(eid):
	#llamar a queryUser.js con certificado
	proc = sp.Popen( ['node', '../queryUser.js' , eid], stdout=sp.PIPE, stderr=sp.PIPE)
	output, error = proc.communicate()
	print("Output: ",output)
	print("Error: ", error)
	j = json.loads(output)
	msp = j['dep']['msp']
	dep = j['dep']['name']
	print(dep, msp)
	return str(dep), str(msp)

def queryPolicyChaincode(from_dep,from_msp,to_eid,to_msp):
	proc = sp.Popen( ['node', '../queryPolicy.js' , to_msp, to_eid, from_msp, from_dep], stdout=sp.PIPE, stderr=sp.PIPE)
	output, error = proc.communicate()
	print("Output: ",output)
	print("Error: ", error)
	return error == ''




HOST = ''                 # Symbolic name meaning the local host
PORT = 50007              # Arbitrary non-privileged port
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
print('Def timeout:',s.gettimeout())
s.bind((HOST, PORT))
s.listen(1)
print("Socket listen to someone")
conn, addr = s.accept()
print('Connected by', addr)

#converting to mysock
#ct = mysocket(conn)

#Receiving Public key certificate 
#data = ct.myreceive()
data = conn.recv(512)
print("Certificate received", data)
temp_file = "../pub-key-store/User1Org2.pub"

#queryUser to Chaincode
dep, name = queryUserDepChaincode(str(data))

#query if policy exists between departments
if queryPolicyChaincode(dep,name,toDep,toOrg) == False:
	print("Policy doesn't exist")
	conn.send("KO")
	conn.close()
	s.close()
	sys.exit(0)

print("Policy does exist")
conn.send("OK")

mynonce = generateNonce(16)
print(mynonce)

#Enviar nonce
conn.send(mynonce)


#Received signed nonce
dataf0 = conn.recv(256)
f0 = int(dataf0)
print("Received r: ",f0)

dataf1 = conn.recv(256)
f1 = int(dataf1)
print("Received s: ",f1)

res = checkNonce(temp_file,mynonce,f0,f1)

if res == False:
	print("Bad signature")
	conn.send("Bad signature")
else:
	print("Good signature")
	conn.send("Good handshake")
conn.close()
s.close()