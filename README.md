# InternetBox 🌐

> An open-source laboratory to build, understand, and break the Internet — in miniature.

**InternetBox** is an educational and experimental project whose goal is to implement, from scratch and in a modular way, the fundamental building blocks that make the Internet work: HTTP servers, reverse proxies, caches, load balancers, DNS, observability, and fault tolerance.

This project is:

*  Not a framework
*  Not a production-ready server
*  Not a clone of nginx, Envoy, or HAProxy

It **is** a systems engineering laboratory.

---

## Motivation

Most developers *use* Internet infrastructure, but very few have actually **built it**.

This project exists to answer questions like:

* What really happens after you press Enter in a browser?
* Where do latency and failures come from?
* How do retries, timeouts, and caches interact?
* Why are these systems so hard to get right?

Instead of reading about these problems, **you implement them**.

---

## Core Philosophy

* Build things **correctly, not quickly**
* Prefer clarity over cleverness
* Measure everything
* Break things on purpose
* Learn from real constraints (latency, concurrency, failure)

Each component is small, focused, and testable on its own, but becomes more interesting when composed with others.

---

## Components (Planned)

Each component lives as an independent module and can be run standalone or composed into a larger system.

### 1. HTTP Server

* Manual HTTP parsing
* Keep-alive connections
* Chunked transfer encoding
* Real timeouts and limits
* Benchmarks

### 2. Reverse Proxy

* Request routing
* Header manipulation
* Retry policies
* Timeout handling

### 3. Cache Layer

* LRU / LFU strategies
* TTL expiration
* Cache invalidation
* Hit/miss metrics

### 4. Load Balancer

* Round-robin
* Least-connections
* Health checks
* Sticky sessions

### 5. DNS (Simplified)

* Basic record types
* Recursive resolution
* DNS caching
* Failure scenarios

### 6. Observability

* Structured logging
* Latency tracking
* Distributed tracing
* Metrics export

### 7. Chaos Engine

* Artificial latency
* Packet loss simulation
* Random node failures

---

## What This Project Optimizes For

* Learning value
* Architectural clarity
* Benchmark-driven decisions
* Reproducible experiments
* Open-source collaboration

It does **not** optimize for:

* Production readiness
* Feature completeness
* Backward compatibility

---

## Tech Stack

The exact stack may evolve, but the project prioritizes:

* A systems-oriented language for the core (e.g. Go, Rust, Java, C)
* Minimal dependencies
* Clear module boundaries
* Simple tooling

UI components (dashboards, visualizations) may use JavaScript, but the core logic does not depend on it.

---

## Roadmap (High Level)

* Phase 1: Minimal HTTP server + benchmarks
* Phase 2: Reverse proxy
* Phase 3: Cache + load balancer
* Phase 4: DNS
* Phase 5: Observability
* Phase 6: Chaos and failure injection

This roadmap is intentionally flexible and driven by learning goals rather than deadlines.

---

## Contributing

Contributions are welcome, especially:

* New components or experiments
* Performance improvements
* Benchmarks and measurements
* Documentation and explanations

If you are interested in systems programming, networking, or infrastructure, this project is meant for you.

---

## Disclaimer

This project is for **educational and experimental purposes only**.

Do not use it in production.

---

## License

GPL-3

---

## Contributors
- **Perseo** (Creator of the repository)

## Final Note

If this project makes you read RFCs, stare at latency graphs, or question simple design decisions — it is working as intended.