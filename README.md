# HTTP requests hashier
Simple pub-sub job doer with a CLI interface.

## Usage

./whash <route name>... [-limiter <num>]
  
  Hashier makes a concurent requests to routes given and print route and hash of response line by line.
  Query can be any string (but if it is not meet RFC 7231 requirements response would not be run).
 
