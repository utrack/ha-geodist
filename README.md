# ha-geodist

This tool reads a list of ID-ed coordinates from the CSV and prints out
five closest and five farthest points to the hardcoded reference point.

## Installation

Using `vgo`:
```
git clone https://github.com/utrack/ha-geodist.git
cd ha-geodist
vgo build
```

Using `go` (dependencies will be unmanaged):
```
go get -u github.com/utrack/ha-geodist
```

## Usage

```
cat coords.csv | ./ha-geodist
# OR
./ha-geodist -csv ./coords.csv
```
