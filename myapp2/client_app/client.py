import socket, sys
from fastecdsa import curve, ecdsa, keys
from hashlib import sha384
import time

def signNonce(file, m):
	private_key, _ = keys.import_key(file)
	print("private key: ",private_key)

	# standard signature, returns two integers
	r, s = ecdsa.sign(m, private_key)
	# should return True as the signature we just generated is valid.
	print("r: ",r)
	print("s: ",s)
	return r, s

filePubKey = sys.argv[1]
filePrivKey = sys.argv[2]
HOST = ''    # The remote host
PORT = 50007              # The same port as used by the server
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, PORT))


#Init request with Public Key
with open(filePubKey, 'r') as myfile:
    dataPubKey=myfile.read() #.replace('\n', '')
print(dataPubKey)
s.send(dataPubKey)

r_policy = s.recv(4)
if r_policy != "OK":
	print("Server: Policy doesn't exists")
	s.close()
	sys.argv(0)


#Receive and sign Nonce
nonce = s.recv(32)
print("Received nonce: ",nonce)
f0, f1 = signNonce(filePrivKey,nonce)

#Send signed nonce
print("Sending r: ", str(f0))
s.send(str(f0))
#time.sleep(3)
print("Sending s: ", str(f1))
s.send(str(f1))

#Receive response
res = s.recv(1024)
s.close()

print('Response from server: ',res)
