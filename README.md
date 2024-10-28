# PufETH Conversion Rate Tracker

![site](site.png)
![demo](https://youtu.be/UWrf2U2yJho)

# Tech Stack

- Timescale DB (Storage)
- Redis (Communication)
- Golang (Backend)
- NextJS (Frontend)

## Backend (Golang)

- Geth -> ETH Client calls & encoding/decoding calldata
- Chi -> Routing web requests
- Solc -> Compiling ABIs into Go bindings
- Pg -> Driver for postgresql
- sqlc -> Golang-specific generated code using SQL queries

## Frontend (Typescript NextJS)

- Rechart -> Charting library
- ReactJS -> Frontend library
- Axios -> Web requests
- Assets from Puffer Finance

# System Design

![design](sys-design.png)

## Frontend <> Backend
