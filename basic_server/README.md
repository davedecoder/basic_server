# Basic Go Server

This is an example of a basic concurrent server written in GO using http.Server included in go itself.

We have 3 endpoints, `/inc`, `/get`, `/set`.

You can interact with them with the following cURLs.

```bash
# GET
curl -X 'GET' \
  'http://localhost:8000/get?name=i'
```

```bash
# SET
curl -X 'GET' \          
  'http://localhost:8000/set?name=i&val=0'
```

```bash
# INC
curl -X 'GET' \          
  'http://localhost:8000/inc?name=i'
```

## Benchmark it using Apache Benchmark

If you have access to a Mac or Linux system, chances are you already have Apache Bench installed as a program called ab. Run a simple load test with:

Use the following command:
```bash
ab -n 6000 -c 200 'http://localhost:8000/inc?name=i'
```

Try to increase the number of request sent by `ab` and surely the localhost won't accept more connections, to work around 
that issue try to run the server within a Docker container then benchmark it or use the `ab` tool inside a container and 
fire requests against the server running in standalone mode, or run both tools in separate containers and reproduce the test.