internet-in-a-box/
├── README.md
├── LICENSE
├── .gitignore

├── docs/
│   ├── architecture.md        <-- General system vision
│   ├── roadmap.md             <-- Milestones and phases
│   ├── design-decisions.md    <-- Why certain decisions were made
│   └── references.md          <-- Useful RFCs, papers, links

├── core/
│   ├── net/                   <-- Network natives (TCP, UDP, sockets)
│   │   ├── connection/
│   │   └── listener/
│   │
│   ├── io/                    <-- Buffers, async, backpressure
│   │   ├── buffer/
│   │   └── scheduler/
│   │
│   ├── protocol/
│   │   ├── http/              <-- HTTP parsing, encoding, state machine
│   │   └── dns/               <-- DNS protocol
│   │
│   ├── config/                <-- Configuration & hot-reload
│   │
│   ├── metrics/               <-- Latencies, counters, tracing
│   │
│   └── chaos/                 <-- Fail injection
│
├── components/
│   ├── http-server/
│   │   ├── server/
│   │   ├── handler/
│   │   └── tests/
│   │
│   ├── reverse-proxy/
│   │   ├── router/
│   │   └── tests/
│   │
│   ├── cache/
│   │   ├── eviction/
│   │   └── tests/
│   │
│   ├── load-balancer/
│   │   ├── algorithms/
│   │   └── healthcheck/
│   │
│   └── dns/
│       └── resolver/
│
├── cmd/
│   ├── http-server/            <-- Executable binary
│   ├── proxy/
│   └── playground/             <-- Experiments and demos
│
├── tests/
│   ├── integration/
│   └── chaos/
│
├── benchmarks/
│   ├── http/
│   ├── proxy/
│   └── cache/
│
└── tools/
    ├── loadgen/                <-- Load generator
    └── packet-simulator/       <-- Network Simulation
