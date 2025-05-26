# siphash-mapper
This is a small helper tool that distributes items amongst targets according to the siphash algorithm.

## Usage
Create a `config.toml` file in the same directory as the binary with the following contents:

config.toml

```toml
hashkey = "loremipsumdolors"

items = [
    "abc",
    "xyz123314",
]

targets = [
    "target1",
    "target2",
]

```

Run the program if you have cloned the source code using

```bash
go run main.go
```

Otherwise, run the binary:

```bash
./siphash-mapper
```

The program will calculate the distribution and display the results.

Example:

```console
> go run main.go
abc mapped to target1 at index 0
xyz123314 mapped to target2 at index 1
```
