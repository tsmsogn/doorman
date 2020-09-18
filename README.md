# Doorman

椅子がドアを通るか確認する君

## Build

```shell
git clone https://github.com/tsmsogn/doorman.git
cd doorman/cmd/doorman
go build
```

## Usage

```
$ ./doorman --chair 1 --chair 2 --chair 3 --door 4 --door 5
通れます
$ echo $?
0
```
