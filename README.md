## GeminiGo

GeminiGo is a CLI tool (REPL) written in Go that allows you to interact with Gemini, a multimodal AI language model developed by Google.
With GeminiGo, you can ask Gemini questions, have conversations, and generate text.

### Installation

To install GeminiGo as an executable, you can use the following steps:

1. Clone the GeminiGo repository:

```
git clone https://github.com/codelikesuraj/gemini-go.git
```

2. Change directory to the GeminiGo directory:

```
cd gemini-go
```

3. Copy .env.example to .env
```
cp .env.example .env
```

4. Update your API key in the .env file
```
GEMINI_API_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXX
```

5. Run the program
```
go run main.go
```
6. Optional: Build the GeminiGo executable:
    1. run the go build command
    
    ```
    go build
    ```
    
    2. Move the GeminiGo executable to a directory that is in your PATH environment variable. For example, on Unix systems, you can move the executable to the `/usr/local/bin` directory:

    ```
    sudo mv gemini-go /usr/local/bin
    ```

    3. Verify that the GeminiGo executable is installed correctly by running the following command:

    ```
    gemini-go
    ```

    You should see the application running on the foreground.

### Usage

Once GeminiGo is installed, you can use it to interact with Gemini by providing a prompts after opening it.

### License

GeminiGo is licensed under the MIT License.