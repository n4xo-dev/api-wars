#!/bin/bash

# Define the directories to search for JS files and corresponding Go programs
directories=("rest/tests" "graphql/tests" "grpc/tests")

# Loop through each directory
for dir in "${directories[@]}"; do
  # Get the base directory (rest, graphql, grpc) from the path
  base_dir=$(dirname "$dir")

  # Check if the directory exists
  if [ -d "$dir" ]; then
    # Check if the Go main file exists
    go_file="$base_dir/cmd/main.go"
    if [ -f "$go_file" ]; then
      # Loop through all JS files in the current directory
      for jsfile in "$dir"/*.js; do
        echo "Running Go program: $go_file"
        # Reset the database
        go run "$go_file" -reset
        # Run the Go Server
        go run "$go_file" -server &
        go_pid=$!

        # Allow the Go server some time to start (adjust if needed)
        sleep 3
        # Check if there are any JS files
        if [ -f "$jsfile" ]; then
          # Remove the extension from the file name
          base_name=$(basename "$jsfile" .js)

          echo "Running k6 on $jsfile"
          # Run k6, output JSON to a file, and CLI output to a text file
          k6 run "$jsfile" --out json="tests/$base_name.json" > "tests/$base_name.txt" 2>&1
        fi
        # Kill the Go program after k6 tests are run
        if [ ! -z "$go_pid" ]; then
          echo "Killing Go process with PID $go_pid"
          kill -2 $go_pid
        fi
      done
    else
      echo "Go program not found in $base_dir/cmd"
    fi
  else
    echo "Directory $dir does not exist"
  fi
done
