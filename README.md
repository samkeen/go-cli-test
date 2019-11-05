# Go CLI experiments

## Using Corba
https://github.com/spf13/cobra

```bash
$ cd hello
$ go build


$ ./hello hello --help
You can define name in an ENV var of HELLO_NAME, 
or a config file with a name: key --config =config.yml
or the -name (-n_ flag)

Usage:
  hello hello [flags]

Flags:
  -h, --help          help for hello
  -n, --name string   Set your name (default "Anonymous")

Global Flags:
      --config string   config file (default is $HOME/.hello.yaml)
```