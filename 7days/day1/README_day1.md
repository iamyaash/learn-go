# Day 1
> **Everything is mentioned clearly, so read it carefully.**
## Download the Installation File
1. Download the installation file for your operating system from the [official Go website](https://go.dev/doc/install).

## 1. Install `Go` on Linux
1. Change the directory to the location of the downloaded file in your terminal.

2. Execute the following command to remove any existing Go installation and extract the new files:
   ```bash
   sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf <filename>.tar.gz
   ```
> Replace **\<filename\>.tar.gz** with the actual filename of the downloaded tarball (e.g., go1.22.4.linux-amd64.tar.gz).

>Note: The above command will remove existing Go files in the specified directory and proceed to extract the files to the /usr/local directory.

## 2. Add `PATH` in environment variable

- Open your `.bash_profile` or `.profile`
```bash
nano ~/.bash_profile
```

- Add the following line under `# User specific environment and startup programs
` in the `./bash_profile` file.
```bash
export PATH=$PATH:/usr/local/go/bin
```

- Save & Close the file, then enter the following command:
```bash
source /.bash_profile
```

## 3. Set-up `Go` Workspace

- First create a folder wherever you want.
```bash
mkdir go-learn
```

- Initialize go in the directory using this command:
```bash
go init mod example/hello
```

- After entering the above command, the directory will be workspace to work on `go`.

- Create a file with `.go` as extension & open in any editor. As it supports almost every IDEs.
## 4. Run your first program
- Enter this code,
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```
- Right after saving the file, enter this command to run your first `Go` program:
```bash
go run .
```
or
```bash
go run <filename>
```
## 5. Tools & Features
- To look for packages, visit this [pck.go.dev](https://pkg.go.dev/)
- To download & install packages,
```bash
go get <pkg.go.dev package name>
```

- To install all the package that are used in the `.go` file:
```bash
go mod tidy
```
> The command will download the packages used in go files that are not yet download. 