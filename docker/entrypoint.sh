#!/bin/sh

echo "Starting entrypoint.sh"

# Print the content of the /rust_server directory
ls -l /usr/local/bin/

# Start rust server
/usr/local/bin/app &

# Print the process list to check if rust_server is running
ps aux

# Start Nginx 
nginx -g 'daemon off;'
