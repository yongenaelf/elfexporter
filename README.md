# ELFexporter

A lightweight Prometheus exporter and [Grafana Dashboard](https://grafana.com/dashboards/6970) that will output AElf wallet balances from a list of addresses you specify. ELFexporter attaches to an aelf node to fetch aelf wallet balances for your Grafana dashboards.

## Watch Addresses
The `addresses.txt` file holds all the addresses to fetch balances for. Use the format `name:address` on each new line.

## Running the Exporter

1. Set Environment Variables:

```sh
export AELF_URL="http://127.0.0.1:8000"
export PORT="8080"
export PREFIX="my_aelf"
export SLEEP_DURATION="15"
```

2. Create the `addresses.txt` file:

```sh
a:2tWqvAgJ8Bsw97YQKs6wWEpBsLzbKBhMaYbbAkm5umJcTEY9oy
b:2V5B7EsPTEo7yyTVmzaWRiVtaHfKQzmkFz29QmFtt7vnZk7Rxi
c:VAXhzz3qRNxuweC38cYt26Qw8Ladp7uty6srBUjK6KbNh8BG8
d:2W4cT6Z3WJ2AVehTnZ6psKQEE2ZPc2mhv9Q7K3as2aNvKP7cn7
```

3. Run the Exporter:

```sh
go run main.go
```

This will start the AElf balance exporter, which will check the balances every 15 seconds and expose the data on the specified HTTP port.

## Grafana Dashboard
TODO: Something like [this](https://grafana.com/dashboards/6970)

## Build Docker Image
Clone this repo and then follow the simple steps below!

##### Build Docker Image
`docker build -t yongenaelf/elfexporter:latest .`

The Docker image should be running with the default addresses.

## Prometheus Response
```
aelf_balance{name="a",address="2tWqvAgJ8Bsw97YQKs6wWEpBsLzbKBhMaYbbAkm5umJcTEY9oy"} 24919.37437
aelf_balance{name="b",address="2V5B7EsPTEo7yyTVmzaWRiVtaHfKQzmkFz29QmFtt7vnZk7Rxi"} 687509.5097
aelf_balance{name="c",address="VAXhzz3qRNxuweC38cYt26Qw8Ladp7uty6srBUjK6KbNh8BG8"} 72284.47401
aelf_balance{name="d",address="2W4cT6Z3WJ2AVehTnZ6psKQEE2ZPc2mhv9Q7K3as2aNvKP7cn7"} 159592.0022
aelf_balance_total 944305.360280000022612512
aelf_load_seconds 1.15
aelf_loaded_addresses 4
aelf_total_addresses 4
```
