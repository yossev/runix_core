# Use a base image with Python and Bash
FROM python:3.9-slim

# Install necessary tools and Node.js
RUN apt-get update && \
    apt-get install -y --no-install-recommends bash curl g++ && \
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
