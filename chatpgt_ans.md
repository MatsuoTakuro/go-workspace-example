Sure, let's create a simple project that contains multiple modules. We'll use the `go.work` file to manage these modules. The project structure is as follows:

```shell
/myproject
  go.work
  /module1
    go.mod
    main.go
  /module2
    go.mod
    main.go
```

 1, **Create the directories**:

```shell
mkdir myproject
mkdir myproject/module1
mkdir myproject/module2
```

 2, **Create `go.mod` in each module**:

Navigate to each directory and initialize a new module:

```shell
cd myproject/module1
go mod init github.com/yourusername/module1

cd ../module2
go mod init github.com/yourusername/module2
```

 3, **Create a `main.go` file in each module**:

For `module1`:

```shell
cat > main.go <<EOF
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello from module1!")
}
EOF
```

For `module2`:

```shell
cat > main.go <<EOF
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello from module2!")
}
EOF
```

 4, **Create the `go.work` file**:

Go back to the root directory `myproject` and create a `go.work` file:

```shell
cd ..
cat > go.work <<EOF
go 1.18

use (
    ./module1
    ./module2
)
EOF
```

 5, **Running the code**:

Now, you can build and run your modules:

```shell
go run ./module1
go run ./module2
```

You should see the output `Hello from module1!` and `Hello from module2!`, respectively.

The `go.work` file lets you treat multiple modules as if they were one, so you don't need to adjust `go.mod` files or use `replace` directives. In this example, `module1` and `module2` could easily import each other, or a third module could import both, and `go build`, `go test`, and other commands would work as expected.

Remember to replace `github.com/yourusername` with your actual GitHub username or another module path of your choice.
