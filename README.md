# Go CLI experiments

## Using Cobra (command)
https://github.com/spf13/cobra

```bash
$ cd prat
$ go build


$ ./prat help
Using config file: /Users/sam/.prat/config.yml  (ENV vars will override config values)
This util is used to help you create a well formed Pull Request

You will be asked a few questions culminating in the construction of your PR

Usage:
  prat [command]

Available Commands:
  hello       hello from Prat
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.prat/config.yaml)
  -d, --dry-run         try out the command without any lasting effect
  -h, --help            help for prat
  -t, --toggle          Help message for toggle

Use "prat [command] --help" for more information about a command.
```

```bash
$ ./prat hello -n Bob
Hello Bob
```

```bash
$ mkdir $HOME/.prat
$ cp config.yml $HOME/.prat/
$ ./prat hello
Using config file: /Users/sam/.prat/config.yml  (ENV vars will override config values)
Hello Name From Config
```

```bash
$ PRAT_NAME="Name from ENV" ./prat hello
Using config file: /Users/sam/.prat/config.yml  (ENV vars will override config values)
Hello Name from ENV
```