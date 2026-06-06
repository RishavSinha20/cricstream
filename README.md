# 🏏 CricStream

**Real-Time Cricket Analytics Platform built using Golang, Kafka, Redis, Docker, and Event-Driven Architecture**

CricStream is a distributed event streaming platform that simulates live cricket matches, processes ball-by-ball events through Kafka, computes real-time analytics using concurrent worker pools, stores aggregated statistics in Redis, and serves them through REST APIs and a live dashboard.

The project was built to explore backend engineering concepts including event-driven systems, concurrency, distributed messaging, caching, and scalable system design.

---

# Architecture

```text
+------------------+
| Match Simulator  |
+------------------+
          |
          v

+------------------+
| Kafka Producer   |
+------------------+
          |
          v

+------------------+
| Kafka Topic      |
| match-events     |
+------------------+
          |
          v

+------------------+
| Consumer Group   |
+------------------+
          |
          v

+------------------+
| Worker Pool      |
| (Goroutines)     |
+------------------+
          |
          v

+------------------+
| Analytics Engine |
+------------------+
          |
          v

+------------------+
| Redis            |
+------------------+
          |
          v

+------------------+
| Gin REST API     |
+------------------+
          |
          v

+------------------+
| Dashboard        |
+------------------+
```

---

# Features

## Match Event Simulation

Simulates live cricket matches by generating ball-by-ball events such as:

* Singles
* Doubles
* Fours
* Sixes
* Wickets
* Dot Balls

These events are streamed into Kafka.

---

## Kafka Event Streaming

Uses Apache Kafka to decouple producers and consumers.

Benefits:

* Scalable event processing
* Replayable event streams
* Loose coupling
* Support for multiple consumers

---

## Concurrent Worker Pool

Events are processed by a bounded worker pool implemented using Go goroutines and channels.

Benefits:

* Controlled concurrency
* Improved throughput
* Avoids unlimited goroutine creation
* Production-style event processing

---

## Real-Time Analytics

Computes:

* Current Score
* Wickets
* Fours
* Sixes
* Balls Faced
* Run Rate

Analytics are continuously updated as events arrive.

---

## Redis Analytics Cache

Aggregated match statistics are stored in Redis.

Benefits:

* Fast reads
* Reduced recomputation
* Low-latency API responses

---

## REST API

Provides real-time access to analytics.

### Get Match Statistics

```http
GET /matches/:id
```

Example:

```json
{
  "MatchID": "ipl_final_2026",
  "Score": 182,
  "Wickets": 4,
  "Balls": 108,
  "Fours": 16,
  "Sixes": 8,
  "RunRate": 10.11
}
```

---

## Live Dashboard

A lightweight dashboard displays:

* Live score
* Run rate
* Boundaries
* Match statistics
* Analytics trends

The dashboard fetches updates continuously from the API.

---

# Tech Stack

## Backend

* Golang
* Gin

## Messaging

* Apache Kafka
* Sarama

## Cache

* Redis

## Infrastructure

* Docker
* Docker Compose

## Frontend

* HTML
* CSS
* JavaScript
* Chart.js

---

# Distributed Systems Concepts Demonstrated

### Event-Driven Architecture

Producers and consumers communicate through Kafka topics.

---

### Consumer Groups

Allows multiple consumers to process events in parallel.

---

### Worker Pools

Bounded concurrency using goroutines and channels.

---

### Asynchronous Processing

Analytics computation is separated from event generation.

---

### Caching Layer

Redis acts as a serving layer for real-time analytics.

---

### Scalability

The system can scale horizontally by increasing:

* Kafka partitions
* Consumer instances
* Worker pool size

---

# Project Structure

```text
cricstream/

├── cmd/
│   ├── simulator/
│   ├── consumer/
│   └── api/
│
├── internal/
│   ├── analytics/
│   ├── kafka/
│   ├── models/
│   ├── redisstore/
│   └── worker/
│
├── dashboard/
│   └── dashboard.html
│
├── docker-compose.yml
│
└── README.md
```

---

# Running the Project

## Start Infrastructure

```bash
docker compose up -d
```

---

## Start Consumer

```bash
go run cmd/consumer/main.go
```

---

## Start Match Simulator

```bash
go run cmd/simulator/main.go
```

---

## Start API

```bash
go run cmd/api/main.go
```

---

## Launch Dashboard

Open:

```text
dashboard.html
```

or

```bash
python -m http.server 5500
```

Then visit:

```text
http://localhost:5500/dashboard.html
```

---

# Sample Data Flow

```text
Match Simulator

Ball Event:
FOUR

      ↓

Kafka Producer

      ↓

Kafka Topic

      ↓

Consumer Group

      ↓

Worker Pool

      ↓

Analytics Engine

      ↓

Redis

      ↓

REST API

      ↓

Dashboard
```

---

# What I Learned

* Building event-driven systems in Go
* Kafka producer/consumer architecture
* Consumer groups and message processing
* Concurrent worker pools using goroutines
* Redis-backed analytics serving
* Designing scalable backend systems
* Real-time dashboard integration
* Dockerized local development environments

---

# Future Improvements

* Dead Letter Queues (DLQ)
* Retry mechanisms
* Prometheus metrics
* Grafana dashboards
* Historical match replay
* Kubernetes deployment
* Multi-match analytics support
