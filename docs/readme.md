# Code Execution Service

This project implements a secure code execution service that runs user-provided code in isolated Docker containers. It currently supports Python and Bash scripts.

## Features

- Executes code in isolated Docker containers for security
- Supports Python and Bash languages
- Limits resource usage (memory, CPU, processes) for each execution
- Truncates large outputs to prevent excessive memory usage

## Implementation Details

The main functionality is implemented in the `executor` package:

- `ExecuteCode(code, language string)`: Main function to execute code
- `getCommand(language, code string)`: Helper function to prepare the Docker command

### Security Measures

- Uses Docker containers for isolation
- Disables network access for executed code
- Limits memory usage to 100MB
- Disables swap
- Limits CPU usage to 0.5 cores
- Restricts the number of processes/threads to 50

### Output Handling

- Captures both stdout and stderr
- Truncates output to 65533 characters to prevent excessive memory usage

## TODO List

1. Add support for more programming languages (e.g., JavaScript, Ruby, Java) [DONE AND MORE TO BE ADDED]
 x 2. Implement input handling for interactive programs
3. Add a timeout mechanism to prevent long-running executions [DONE]
4. Create an API endpoint to receive code execution requests
5. Implement proper error handling and logging [DONE]
6. Add unit tests for the executor package [DONE]
7. Set up CI/CD pipeline for automated testing and deployment
8. Create a user interface for code submission and result display
9. Implement rate limiting to prevent abuse
10. Add a caching mechanism for frequently executed code snippets
11. Implement a sandboxing solution for additional security (e.g., gVisor)
12. Add support for file I/O within the container (with proper restrictions)
13. Implement a queuing system for handling multiple execution requests
14. Implement a cleanup mechanism to remove unused Docker images and containers

## Getting Started

(Add instructions for setting up and running the project locally)

## Contributing

(Add guidelines for contributing to the project)

## License

(Add license information for your project)