# IPInfo
IPInfo was created to aid in getting relevant information about an IP and its subnet. 

## Currently returned by IPInfo:
- Subnet mask
- First available host
- Max hosts
- Broadcast address
- Network address

## Usage
`ipinfo <ip> <# of subnets>`

After running this command:

`ipinfo 214.45.7.102 22`

the CLI will return the following information:
```
Subnet mask: 255.255.252.0
First available host: 214.45.7.103
Max hosts: 1022
Broadcast address: 214.45.7.255
Network address: 214.45.4.0
```