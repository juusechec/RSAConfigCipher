#!/bin/bash -eux
go build -o rsaconfigcipher decrypt_files.go

./rsaconfigcipher *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher *.rsa -v
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher *.rsa --verbose
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher -v *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa -p keys/rsakey.pem.pub
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa -P keys/rsakey.pem
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher --verbose *.rsa -p keys/rsakey.pem.pub -P keys/rsakey.pem
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher -p keys/rsakey.pem.pub *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher -P keys/rsakey.pem *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher -p keys/rsakey.pem.pub -P keys/rsakey.pem *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher --public-key-path keys/rsakey.pem.pub *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher --private-key-path keys/rsakey.pem *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher --public-key-path keys/rsakey.pem.pub --private-key-path keys/rsakey.pem *.rsa
md5sum -c .md5checksum
rm -rf example_config.{yml,php}

./rsaconfigcipher -h
./rsaconfigcipher --help

./rsaconfigcipher --version

echo "I WANT TO ENCRYPT" | ./rsaconfigcipher

echo "I WANT TO ENCRYPT" | ./rsaconfigcipher -i

printf "I WANT TO ENCRYPT" | ./rsaconfigcipher

printf "I WANT TO ENCRYPT" | ./rsaconfigcipher -i

# validate encrypt-decrypt a string
printf "I WANT TO ENCRYPT" | ./rsaconfigcipher -s > constructed_config.txt.rsa
./rsaconfigcipher constructed_config.txt.rsa
CYPHERTEXT=$(cat constructed_config.txt)
if [ ! "$CYPHERTEXT" == "I WANT TO ENCRYPT" ]
then
  echo "Error, the cypher text is not equals"
  exit 1
fi
rm -rf constructed_config.txt{,.rsa}

# validate encrypt-decrypt a long text

while IFS='' read -r line || [[ -n "$line" ]]; do
    echo "$line" | ./rsaconfigcipher -s >> constructed_config.txt.rsa
done < "example_long_value.txt"
./rsaconfigcipher constructed_config.txt.rsa
if [ ! "$(md5sum constructed_config.txt | awk '{ print $1 }')" == "$(md5sum example_long_value.txt | awk '{ print $1 }')" ]
then
  echo "Error, the cypher text of files is not equals"
  exit 1
fi
rm -rf constructed_config.txt{,.rsa}

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
