# Generate the service
if [ -z "$1" ]; then
    echo "Usage: bash ./generate.sh \"your service\""
    exit 1
fi
go run main.go "$1"