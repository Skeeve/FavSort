# Readme

`FavSort` allows you to easily sort your favorites on your Anbernic RG35XX H.

I'm pretty confident it will run as well on the "Plus" and "SP", but couldn't test that.

For more information please check the [Readme packaged with the binary](favsort/artifacts/Readme.md).

## Contributing

To get a development enviroment, `docker-compose.yaml` is provided.
To start, execute

```shell
docker-compose run --rm godev
```

This will link the required directories into the container.

### Compiling

```shell
# Inside the container
cd PORTS
go build ../cmd/favsort
```

This will create the `favsort` binary which is linked to `++ sort Favorites ++`.

### Packaging

```shell
# Outside the container
./pkg-FavSort
```

This will create `FavSort.zip`.
It requires a proper `go.mod` which you should have created in the development container.

## Credits

Many thanks to the reddit user [AnonymousTokenus](https://www.reddit.com/user/AnonymousTokenus/), for

- giving me the idea to this tools
- beta-testing
- helping me figuring out the meaning of the fields in the favorites file
- and for challenging my programming skills