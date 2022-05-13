openssl genrsa -out myCA.key 2048
openssl req -x509 -new -nodes -key myCA.key -sha256 -days 36500 -out myCA.pem



openssl genrsa -out ec2.amazonaws.com.key 2048
openssl req -new -key ec2.amazonaws.com.key -out ec2.amazonaws.com.csr
openssl x509 -req -in ec2.amazonaws.com.csr -CA myCA.pem -CAkey myCA.key -CAcreateserial -out ec2.amazonaws.com.crt -days 36500 -sha256 