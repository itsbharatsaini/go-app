## Go Application with GolangCI-Lint

This project is a Go web application with GolangCI-Lint integration for enforcing coding standards.

## Prerequisites
 - Go: https://go.dev/doc/install
 - GolangCI-Lint: https://github.com/golangci/golangci-lint

## Getting Started

Follow these steps to run the Go application locally:

1. Clone this repository to your local machine:

    ```bash
    git clone <repository-url>
    ```

2. Navigate to the project directory:

    ```bash
    cd <project-directory>
    ```

3. Build the Go application:

    ```bash
    go build -o <output-file-name>
    ```

4. Run the Go application:

    ```bash
    ./<output-file-name>
    ```

## Running GolangCI-Lint

```bash
golangci-lint run
```
This command will execute all the configured linters and display any linting errors or warnings.

## Docker image

### Build the image
 ```bash
docker build -t <image-name> .
```
Replace `<image-name>` with the desired name for your Docker image.

### Run docker container:
```bash
docker run -p 8080:8080 <image-name>
```

Once the Docker container is running, you can access the application at http://localhost:8080 in your web browser.