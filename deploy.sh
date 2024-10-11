go build -o http_proxy ./cmd/proxy/main.go
cp http_proxy /home/http_proxy
chown ec2-user:ec2-user /home/http_proxy
cp http_proxy.service /etc/systemd/system/http_proxy.service
systemctl restart http_proxy.service