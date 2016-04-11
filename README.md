# flux

Event fuzzler for testing and capacity testing.

## Usage

### Flags

```
   -t, --template 		template to use as generator
   -d, --duration "1m0s"	test duration in seconds
   --help, -h			show help
   --version, -v		print the version
```

### Output

To send to stdout

```flux -t <template> -d <duration> stdout```

To send to kafka:

```flux -t <template> -d <duration> kafka --zk <zookeeper> --topic <topic>```

### Templates

Templates are specified in go template syntax. Some helpers were added:

`rand int1 int2`: random number between int1 and int2

`randFloat float1 float2`: random number between float1 and float2

`firstName`: random first names

`lastName`: random last names

`choice a b c`: random select between a, b or c

`sequence 10`: loop helper

```json
{
    "data": [
        {{- range $i, $_ := sequence 2 }}
        {{- if $i}},{{end}}
        {
            "id": "{{ rand 10 1000000 }}",
            "name": "{{ firstName }} {{ lastName }}",
            "gender": "{{ choice "male" "female" }}",
            "type": "{{ choice 1 2 3 }}"
        }
        {{- end}}
    ]
}
```

