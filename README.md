# yaml-parser
The App compares multiple yaml configurations. 

## How to use
go run main.go config1.yml config2.yml config3.yml ...


## Dockerfile
docker build -t yaml-parser .

### Show help
docker run yaml-parser

### Show version
docker run yaml-parser -version

### Compare config files
docker run --rm -v $(pwd)/configs_example:/configs yaml-parser /configs/config1.yaml /configs/config2.yaml /configs/config3.yaml /configs/config4.yaml