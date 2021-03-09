# Gocipher
Gocipher is a cross platform command line tool for encryption and decryption, including RSA, DES, BASE64, MD5, SHA256.
# Install
To install the library, follow the classical:

    $ go get github.com/godaner/gocipher
    
Or get it from the released version: 

    https://github.com/godaner/gocipher/releases
    
> Note: curl -LJO https://github.com/godaner/gocipher/releases/download/v1.0.3/gocipher_linux-amd64.tar.gz, tar -zxvf gocipher_linux-amd64.tar.gz

# Supported platforms

This library works (and is tested) on the following platforms:

<table>
  <thead>
    <tr>
      <th>Platform</th>
      <th>Architecture</th>
      <th>Status</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="2">Linux</td>
      <td><code>amd64</code></td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>386</code></td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="2">Windows</td>
      <td><code>amd64</code></td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>386</code></td>
      <td>✅</td>
    </tr>
    <tr>
      <td>Others</td>
      <td><code>Others</code></td>
      <td>⏳</td>
    </tr>
  </tbody>
</table>

# Usage
```
$ ./gocipher
NAME:
   gocipher - Gocipher is a cross platform command line tool for encryption and decryption, including RSA, DES, BASE64, MD5, SHA256.

USAGE:
   gocipher [global options] command [command options] [arguments...]

VERSION:
   v1.0.3

COMMANDS:
   rsaenc     encrypt by rsa
   rsadec     decrypt by rsa
   desenc     encrypt by des
   desdec     decrypt by des
   base64enc  encrypt by base64
   base64dec  decrypt by base64
   md5        md5
   sha256     sha256
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help     show help
   --version  print the version
```
## RSA
#### encrypt
    ./gocipher rsaenc \
    --textfile ./plaintext \
    -o ./ciphertext \
    --pubkeyfile ./publickey \
    --base64 rawurl
#### decrypt
    ./gocipher rsadec \
    --textfile ./ciphertext \
    -o ./plaintext \
    --prikeyfile ./privatekey \
    --base64 rawurl
    
## DES
#### encrypt
    ./gocipher desenc \
    --textfile ./desplaintext \
    -o ./desciphertext \
    --key 12345678 \
    --base64 rawurl
#### decrypt
    ./gocipher desdec \
    --textfile ./desciphertext \
    -o ./desplaintext \
    --key 12345678 \
    --base64 rawurl
## BASE64
#### encrypt
    ./gocipher base64enc --text godaner
#### decrypt
    ./gocipher base64dec --text Z29kYW5lcg==
## MD5
    ./gocipher md5 --text godaner
## SHA256
    ./gocipher sha256 --text godaner