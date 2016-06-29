# go build sh
genarate a sh file to build go application in multi platform

you can use `go-build-sh <go-excute-name>` to render a build.sh

And then you can do this to build your go application:
```
./build.sh all      # build all ( windows, linux, darwin )
./build.sh windows  # build windows
./build.sh linux    # build linux
./build.sh darwin   # build darwin
./build.sh          # just build default platform
./build.sh help     # for cmd help
``` 