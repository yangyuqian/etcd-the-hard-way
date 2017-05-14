#!/bin/sh

# generate certificates signed by CA
function gencert(){
  echo '{"CN":"'$1'","hosts":[""],"key":{"algo":"rsa","size":2048}}' | cfssl gencert -config=ca-config.json -ca=ca.pem -ca-key=ca-key.pem -hostname="$2" - | cfssljson -bare $1
}

source ./config
source ./default

# generate self-signed ca
echo '{"CN":"CA","key":{"algo":"rsa","size":2048}}' | cfssl gencert -initca - | cfssljson -bare ca -
echo '{"signing":{"default":{"expiry":"43800h","usages":["signing","key encipherment","server auth","client auth"]}}}' > ca-config.json

NAMES="$MEMBERS $SERVER_NAME $CLIENT_NAME"

for m in $NAMES
do
  gencert $m $MEMBER_ADDR
done

