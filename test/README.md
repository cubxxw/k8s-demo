# k8s demo 测试参考目录

```bash
project-root-directory/
│
├── cmd/                      # Command applications
│   ├── app1/
│   └── app2/
│
├── pkg/                      # Libraries and packages
│   ├── package1/
│   ├── package2/
│   └── ...
│
├── api/                      # API definitions and protocol files
│   ├── api_definition1/
│   └── api_definition2/
│
├── config/                   # Configuration files and related utilities
│   ├── dev.config.yaml
│   ├── prod.config.yaml
│   └── ...
│
├── logs/                     # Log files (usually gitignored, but structured for tools like Loki)
│   ├── app1.log
│   └── app2.log
│
├── monitoring/               # Monitoring tools and configuration
│   ├── prometheus/           # Prometheus configuration and rules
│   └── grafana/              # Grafana dashboards
│
├── deploy/                   # Deployment configurations and scripts
│   ├── kubernetes/           # Kubernetes manifests
│   ├── helm/                 # Helm charts
│   └── istio/                # Istio configurations and policies
│
├── test/                     # Specialized tests and testing utilities
│   ├── e2e/
│   ├── integration/
│   ├── fuzz/
│   ├── mocks/
│   ├── fixtures/
│   └── helpers/
│
├── scripts/                  # Build, install, and other automation scripts
│
├── .gitignore
├── go.mod                    # Go module file
├── go.sum                    # Go module checksum
├── README.md
└── ...


```