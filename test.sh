#!/bin/bash -eux
go build -o rsaconfigcipher decrypt_files.go

./rsaconfigcipher *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher *.rsa -v
rm -rf example_config.{yml,php}

./rsaconfigcipher *.rsa --verbose
rm -rf example_config.{yml,php}

./rsaconfigcipher -v *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa -p keys/rsakey.pem.pub
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa -P keys/rsakey.pem
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa -p keys/rsakey.pem.pub -P keys/rsakey.pem
rm -rf example_config.{yml,php}

./rsaconfigcipher -p keys/rsakey.pem.pub *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher -P keys/rsakey.pem *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher -p keys/rsakey.pem.pub -P keys/rsakey.pem *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher --public-key-path keys/rsakey.pem.pub *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher --private-key-path keys/rsakey.pem *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher --public-key-path keys/rsakey.pem.pub --private-key-path keys/rsakey.pem *.rsa
rm -rf example_config.{yml,php}

./rsaconfigcipher -h
./rsaconfigcipher --help

./rsaconfigcipher --version

echo "I WANT TO ENCRYPT" | ./rsaconfigcipher

echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -i

CYPHERTEXT=$(echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -i -s)
if [ ! $(echo "$CYPHERTEXT" | egrep "^{{%rsa:.*%}}") ]
then
  echo "Error, the cypher text is incorrect"
  exit 1
fi

CYPHERTEXT=$(echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -is)
if [ ! $(echo "$CYPHERTEXT" | egrep "^{{%rsa:.*%}}") ]
then
  echo "Error, the cypher text is incorrect"
  exit 1
fi

CYPHERTEXT=$(echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -si)
if [ ! $(echo "$CYPHERTEXT" | egrep "^{{%rsa:.*%}}") ]
then
  echo "Error, the cypher text is incorrect"
  exit 1
fi

CYPHERTEXT=$(echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -s)
if [ ! $(echo "$CYPHERTEXT" | egrep "^{{%rsa:.*%}}") ]
then
  echo "Error, the cypher text is incorrect"
  exit 1
fi

CYPHERTEXT=$(echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -s -p keys/rsakey.pem.pub)
if [ ! $(echo "$CYPHERTEXT" | egrep "^{{%rsa:.*%}}") ]
then
  echo "Error, the cypher text is incorrect"
  exit 1
fi

CYPHERTEXT=$(echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -is -p keys/rsakey.pem.pub)
if [ ! $(echo "$CYPHERTEXT" | egrep "^{{%rsa:.*%}}") ]
then
  echo "Error, the cypher text is incorrect"
  exit 1
fi

rm -rf ./rsaconfigcipher

echo "Test Complete Successfully"
