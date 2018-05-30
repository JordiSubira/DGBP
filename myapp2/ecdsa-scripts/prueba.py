from fastecdsa import curve, ecdsa, keys
from hashlib import sha384

m = "a message to sign via ECDSA"  # some message

''' use default curve and hash function (P256 and SHA2) '''
private_key = keys.gen_private_key(curve.P256)
public_key = keys.get_public_key(private_key, curve.P256)

print("private key: ",private_key)
print("public key: ",public_key)

# standard signature, returns two integers
r, s = ecdsa.sign(m, private_key)
# should return True as the signature we just generated is valid.
valid = ecdsa.verify((r, s), m, public_key)

print("Is it valid? = ",valid)

# save the private key to disk
keys.export_key(private_key, curve=curve.P256, filepath='./p256.key')
# save the public key to disk
keys.export_key(public_key, curve=curve.P256, filepath='./p256.pub')