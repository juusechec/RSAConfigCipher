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

./rsaconfigcipher -h
./rsaconfigcipher --help

echo "I WANT TO ENCRYPT" | ./rsaconfigcipher

rm -rf ./rsaconfigcipher

echo "Test Complete Successfully"
