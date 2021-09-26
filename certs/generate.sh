openssl req -x509 -sha256 -nodes -days 3650 -newkey rsa:2048 -subj '/O=ABCD Inc./CN=ABCD Root CA' -keyout abcd.com.key -out abcd.com.crt
openssl req -out app.abcd.com.csr -newkey rsa:2048 -nodes -keyout app.abcd.com.key -subj "/CN=*.abcd.com/O=ABCD Inc."
openssl x509 -req -days 365 -CA abcd.com.crt -CAkey abcd.com.key -set_serial 0 -in app.abcd.com.csr -out app.abcd.com.crt


openssl req -newkey rsa:2048 -nodes -keyout gowiki.key -subj "/CN=*.gowiki-svc-headless.app.svc.cluster.local/O=ABCD Inc." -out gowiki.csr
openssl x509 -req -extfile <(printf "subjectAltName=DNS:gowiki-svc.app.svc.cluster.local") -days 365 -in gowiki.csr -CA abcd.com.crt -CAkey abcd.com.key -CAcreateserial -out gowiki.crt