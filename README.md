# Adjust-Take-Home

## Solution for Adjust Take Home Task

The tool sends parallel requests to the URLs included as parameters. While there is an available spot for parallel requests the tool will keep launching requests until all the available spots are used. Once a new spot is free again, the tool will add a new request. The tool will keep doing this until all the requests have been sent. When the response arrives the tool calculates its MD5 and prints both the URL and the MD5 for the response. If there is no response, the tool will print: "Error requesting: URL".

There are two files contained in this repository:

    * myhttp.go
    * myhttp_test.go

myhttp.go contains the tool's code, to compile the code run:

```bash
    go build myhttp.go
```

Once the code has been compiled, it can be executed. Try one of the following options:

1. Print application help.

```bash
    myhttp -h 
```

2. Launch the application with the maximum of parallel requests as default (10).

```bash
    myhttp www.google.es www.adjust.com www.google.com www.twitter.com
```

3. Launch the application with the maximum of parallel requests set as 4.

```bash
    myhttp -parallel 4 http://google.es http://www.google.com http://www.facebook.com
```

Finally to run the application test execute in the repository the following command in the repository dir:

```bash
    go test -v
```
