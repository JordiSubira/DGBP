from fastecdsa import curve, ecdsa, keys
from hashlib import sha384

import sys

m = sys.argv[2]  # some message

''' use default curve and hash function (P256 and SHA2) '''

private_key, public_key = keys.import_key(sys.argv[1] + '.key')

print("private key: ",private_key)
print("public key: ",public_key)

# standard signature, returns two integers
r, s = ecdsa.sign(m, private_key)
# should return True as the signature we just generated is valid.

print("r: ",r)
print("s: ",s)