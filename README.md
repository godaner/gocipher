# Gocipher
Gocipher is a cross platform command line tool for encryption and decryption, including RSA.
## RSA
#### encode
    ./gocipher rsaenc \
    --textfile ./plaintext \
    -o ./ciphertext \
    --pubkeyfile ./publickey \
    --base64
#### decode
    ./gocipher rsadec \
    --textfile ./ciphertext \
    -o ./plaintext \
    --prikeyfile ./privatekey \
    --base64