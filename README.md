# PufETH Conversion Rate Tracker

![site](site.png)
<video width="640" height="360" controls>

  <source src="site.mov" type="video/quicktime">
  Your browser does not support the video tag.
</video>

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
