# HTTP requests hashier

Simple pub-sub job doer with a CLI interface.

  Hashier makes a concurent requests to routes given and print route and hash of response line by line.
  Query can be any string (but if it is not meet RFC 7231 requirements response would run an error).

## Usage

Make `go build` in a web-hashier directory

Run binary with an arguments:

`./web-hashier [-parallel=<num>] <routes names>...`
  
## Arguments:

  *parallel*  (optional) number of the concurent requests (default is 3, zero value is prohibited)

  *hosts*  list of hosts to process
