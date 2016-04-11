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


