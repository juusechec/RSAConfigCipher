[![Build Status](https://travis-ci.org/juusechec/RSAConfigCipher.svg?branch=master)](https://travis-ci.org/juusechec/RSAConfigCipher)
[![Go Report Card](https://goreportcard.com/badge/github.com/juusechec/RSAConfigCipher?branch=master)](https://goreportcard.com/report/github.com/juusechec/RSAConfigCipher)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

# RSAConfigCipher

With this you can create files and encrypt some values in source code.

Example of encrypted file.

```yml
driver:   pdo_pgsql
host:     "%database_host%"
port:     "%database_port%"
dbname:   "{{%rsa:d1bb75c9660343d5ce620893f82ee55f1e2211a725a0a4967d33d671d478be713d9da391b7ea9820cf9eadb81271638020eedbadfe040719fb87181151f2a2ff756e7b3367bb88cd543c50961ed2c1a1b5de6492ced973ff02eb39b6fc7f065ced0c8c8485bbc862b6ccf23a73e44376b4880b6ef55fdf7a0b14e697ed35f624de70beeaa8c16c219d4b027caf163a58bd15fb7de2073e1905672c52afb07efbed2fe06b04bc1cbe803ae590fd2e384ebefc854a3031e27b50c9497c11d9e67561feaf06c2f5fe2e294ed453cea14aaf8d707b548077a898e34a08aa6d2448b0e82905e2022de98b5925ce8e12e3334dedb26f3653e61e89a16356cbde7d63cf%}}"
user:     "{{%rsa:98cd9f0150862333ac099b3611f3c4b1cc79ca85691fa99f568aed4f9067d4c2a2505b8fbc5b12fc0528cd91160fb0dd94acd7336fd1480908c4dec6b554b864fbbc48a7d27a1fc018e73b5b8629387dc73c7b99db1e166bb912c03b1685c2c516009547be09579e510d87dc2d826a0b778c8a4a8e62b289e9b0e5e5d420ebfee4d869e04b6d9db30bf335bfdc13633fc665d16a6976413ad7b3f80a9d1d2ae4f4e62c9cc8efb1963a5237fe2d9f3ff00f1fec085c2895463edee8651c07d41fa32460f1f2478358802c913b80e3336c4792f987dffc49c736d2550a9e8c9901a0dc4c9f78b7232a5d6730dc6b6ce90c5b1fe58cabeaf5d91e67223a13c6da1c%}}"
password: "{{%rsa:560eb2b9758f0e258a31d2a97c2ee86a60a450d4b5260f2091db0e22e3fbf2e1a399da3032f2836a9bd4e1ae80cf76dc7163901d0f354b76febce9e9fabbbe8d9fdbb5c9f657af74fad8bc619551351c7a14c231e8f2ed82f5fdecf25cee0caf73b3cd037c4c1bd73fab9436af597c663217c6e3b693a076976d9c7bbc1a46a5617b7d50ddbf8cd4ef24932036a87b811ba8f4959edaa2394ab86bac2e4a3fe1487b869ce55ef3a1d9fed90fcf77f78e8a16bb4c3d368d624b7c4a9cb95fd2525339e98575d2aea77a16a82ee682cd5e4f12496a9563576a9728e9fe91ead9b60ae400ca1d7937bba9b6f56717294ffb0fb2df3890174ce4b06f3b4b9ceffbf9%}}"
charset:  UTF8
```

When is decrypted.

```yml
driver:   pdo_pgsql
host:     "%database_host%"
port:     "%database_port%"
dbname:   "nombre_base_datos"
user:     "usuario_base_datos"
password: "More COMPLEX password with %% and $$ ^^&& other chars"
charset:  UTF8
```

## Encrypt values
You can execute ***./rsaconfigcipher*** and paste the desired value, then Intro
key. The result can be copy in configuration file replacing the unencrypted
value.

## Generate pair of private and public key
```
openssl genrsa -out rsakey.pem 2048
openssl rsa -in rsakey.pem -pubout > rsakey.pem.pub
```

## Execute for all files in path
It's recommended rename files with extension "rsa" or whatever, example
***archivo.ext*** would be renamed as ***archivo.ext.rsa*** before
transformation.

```bash
find . -name "*.rsa" -exec ./rsaconfigcipher {} \;
```

## See help
For more options you can use:
```bash
./rsaconfigcipher --help
```

## Build
```bash
go build -o rsaconfigcipher decrypt_files.go
```

## Install as normal user application
For normal user programs not managed by the distribution package.
```bash
wget "https://github.com/juusechec/RSAConfigCipher/releases/download/v1.2.0/rsaconfigcipher" && chmod +x rsaconfigcipher && sudo mv rsaconfigcipher /usr/local/bin/
```

## Alternative install
Its important because sometimes root user haven't */usr/local/bin/* in *PATH*
environment variable.
```bash
wget "https://github.com/juusechec/RSAConfigCipher/releases/download/v1.2.0/rsaconfigcipher" && chmod +x rsaconfigcipher && sudo mv rsaconfigcipher /usr/bin/
```

## Uninstall
```
sudo rm /usr/local/bin/rsaconfigcipher # for normal user
sudo rm /usr/bin/rsaconfigcipher # for alternative install
```
