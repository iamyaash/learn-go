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