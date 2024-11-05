# PufETH Conversion Rate Tracker

### UI

![site](site.png)

### Demo

https://youtu.be/UWrf2U2yJho

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

## Overview

The system uses an Event Driven Architecture (EDA). Specifically, reads on the data can often be more than the writes on the data. Thus, the Command Query Responsibility Segregation (CQRS) is used whereby services are used to write to redis and timescaledb while the web server reads from db and listens to new pubsub events.

By using events, it's now possible to also "attach" new processors that listen for specific events. For example a points processor that calculates certain scores based on the tvl of the vault can now be done async as calculations are often I/O-bound and can be expensive in both latency and throughput to the core systems.

## Fetching Data

Fetching data for puffer can be separated into 2 different types:

1. Real-time Market Data
2. Backfilled events -> missing blocks in time series data or backfill till earliest block x

### Real-time Market Data

The RTMD service subscribes to new block header events from the rpc node (alternatively polling can be done if subscription is not possible). On new block header events, it then requests for the relevant data using a multicall. This service is largely network-bound as most of the time it is doing rpc calls, publishing events and storing to db.

### Backfiller

The backfiller fulfills 2 responsibilities. First to backfill missing blocks in the time-series data. Second is to get more historical data. The backfiller maintains 1 producer to a group of consumers. The producer uses a db cursor to run through the db to find out missing blocks, pushing missing blocks to a queue for consumers to fetch via rpc. A different event can be emitted in order to differentiate between such events.

## Reading data

The web server itself is built to be horizontally scalable in order to accommodate for a large number of requests. In order to reduce the latency of requests, historical time series is first fetched from the db and loaded in-memory into a ring buffer. Following real-time events are then subscribed to in real-time and added into the buffer.

# To run

Unforunately, due to time constraints, there is no all-in-one docker compose and no auto-migration setup yet...

Do note that the following ports are used:

- 5432 (Database)
- 6379 (Redis)
- 3001 (Web Server)
- 3000 (UI)

## Start Backend

```bash
# cd into backend folder
cd backend

# create .env file (feel free to edit if needed)
cp .env.example .env
```

### Start TimescaleDB and Redis

```bash
docker compose up -d
```

Do a migration to create the first table

```sql
CREATE EXTENSION IF NOT EXISTS timescaledb;

CREATE TABLE IF NOT EXISTS puffer_conversion_rates (
	chain TEXT NOT NULL,
	block_number BIGINT NOT NULL,
	block_timestamp_ms BIGINT NOT NULL,
	total_asset NUMERIC NOT NULL,
	total_supply NUMERIC NOT NULL,
	conversion_rate NUMERIC NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY (chain, block_number, block_timestamp_ms)
);

SELECT create_hypertable(
    'puffer_conversion_rates',
    by_range('block_timestamp_ms',86400000), -- 1 day in milliseconds
    if_not_exists => TRUE,
    migrate_data => TRUE
);
```

### Start Services

```bash
# start rtmd
go run ./apps/rtmd/main.go

# start backfiller
go run ./apps/backfiller/main.go

# start web server
go run ./apps/front/main.go
```

## Start Frontend

```bash
# cd into frontend folder
cd frontend

# start UI
yarn build
yarn start
```

In a browser, you can enter http://localhost:3000
