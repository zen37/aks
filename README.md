Illustrates some of 'Azure Kubernetes Services' capabilities: scaling out pods and nodes.


while sleep 0.01; do curl http://<external-IP>:8080/intros \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"prefix": "Hello World", "timestamp": "Jan 02, 2022 07:24:00 AM"}'; done