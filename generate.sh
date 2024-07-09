# Generate the service
if [ -z "$1" ]; then
    echo "Usage: ./generate.sh <YourService>"
    exit 1
fi
go run main.go "$1"