# Go_myNote

## 在 Linux 安裝 go
```
從官網下載壓縮檔

# 解壓縮剛剛下載的壓縮檔
sudo tar -xvf go1.15.linux-amd64.tar.gz
sudo mv go /usr/local

# 最後是設定環境變數，在~/.profile加入以下內容
export PATH=$PATH:/usr/local/go/bin

# 測試環境是否正常，新增一個 main.go檔案 

package main
import "fmt"
func main() {
	fmt.Println("Hello, world")
}

# 接下來在terminal執行
go run main.go

```

## Start a Fiber project

Create a new directory for your project
```
mkdir Project_Name
cd Project_Name
```

initialize your project with Go modules by executing the following command in your terminal
```
go mod init github.com/your/repo
```

After setting up your project, you can install Fiber with the go get command:
```
go get -u github.com/gofiber/fiber/v3
```

Here's a basic example to create a simple web server that responds with "Hello, World 👋!" on the root path.
```
package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
```
run this Go program, and visit http://localhost:3000 in your browser
```
go run .
```

## How to kill the server when failed to listen: address already in use
```
# check the port
lsof -i :3000

# kill PID
kill xxxx
```


