from fastecdsa import curve, ecdsa, keys
from hashlib import sha384

import sys

m = sys.argv[2] # some message

''' use default curve and hash function (P256 and SHA2) '''

_ , public_key = keys.import_key(sys.argv[1] + '.pub')

print("public key: ",public_key)

# should return True as the signature we just generated is valid.
valid = ecdsa.verify(( long(sys.argv[3]), long(sys.argv[4]) ), m, public_key)

print("Is it valid? = ",valid)