# Use a base image with Python, Bash, and Node.js
FROM python:3.9-slim

# Install Bash, Node.js, and clean up in the same layer to reduce image size
RUN apt-get update && \
    apt-get install -y bash curl && \
    curl -fsSL https://deb.nodesource.com/setup_14.x | bash - && \
    apt-get install -y nodejs && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /code

# Create a non-root user for running scripts
RUN useradd -m runixuser

# Switch to the non-root user
USER runixuser
