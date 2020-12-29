# raw_matcher

It is my Go research project.

I want to be able to collect all needed RAW files for photos selected
and backed on NAS drive. Ruby would be not fast enough for scanning
multiple of TB of data and I wanted to learn Go language.

## How to use

```
go run cmd/command.go -path ./data -output copy_script.sh
```

## Tests

`make test`

## TODO

1. Add missing test
2. Convert to internal functions
3. Separate
4. Test for path w/o time
5. Maybe pointers would be better

### Links, notes

https://blog.alexellis.io/5-keys-to-a-killer-go-cli/
