#/bin/bash
echo "Starting"
sleep 3

export PATH=$PATH:/usr/local/go/bin
go run .

# tail -f /dev/null
