# ELFexporter

A lightweight Prometheus exporter and [Grafana Dashboard](https://grafana.com/dashboards/6970) that will output AElf wallet balances from a list of addresses you specify. ELFexporter attaches to an aelf node to fetch aelf wallet balances for your Grafana dashboards.

## Watch Addresses
The `addresses.txt` file holds all the addresses to fetch balances for. Use the format `name:address` on each new line.

## Running the Exporter
You can easily run this AElf balance prometheus exporter with the docker command:
```
docker run -it -d -p 9015:9015 \
  -e "RPC=https://tdvw-test-node.aelf.io" \ 
  -v /myfolder/addresses.txt:/app/addresses.txt \ 
  yongenaelf/elfexporter
```

## Grafana Dashboard
ELFexporter will include a Grafana Dashboard similar to [this](https://grafana.com/dashboards/6970) so you visualize AElf wallet balances with ease.

## Build Docker Image
Clone this repo and then follow the simple steps below!

##### Build Docker Image
`docker build -t yongenaelf/elfexporter:latest .`

The Docker image should be running with the default addresses.

## Prometheus Response
```
elf_balance{name="a",address="2tWqvAgJ8Bsw97YQKs6wWEpBsLzbKBhMaYbbAkm5umJcTEY9oy"} 24919.37437
elf_balance{name="b",address="2V5B7EsPTEo7yyTVmzaWRiVtaHfKQzmkFz29QmFtt7vnZk7Rxi"} 687509.5097
elf_balance{name="c",address="VAXhzz3qRNxuweC38cYt26Qw8Ladp7uty6srBUjK6KbNh8BG8"} 72284.47401
elf_balance{name="d",address="2W4cT6Z3WJ2AVehTnZ6psKQEE2ZPc2mhv9Q7K3as2aNvKP7cn7"} 159592.0022
elf_balance_total 944305.360280000022612512
elf_load_seconds 1.15
elf_loaded_addresses 4
elf_total_addresses 4
```
