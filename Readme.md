### Overview
Tired of logging work daily/weekly? Then this is the CLI tool for you which automatically log work automatically for you with only a single command

For now, the tool only support for Jira cloud.

### How to install
1. Prerequisites:
- go >= 21

2. Install cli
run: 
```
go install github.com/TranThang-2804/auto-logwork
```

### How to use
1. Configure your tool
run:
```
auto-logwork configure
```

2. Auto log work for the fucking whole week:
run:
```
auto-logwork logwork
```

3. Show help
run:
```
auto-logwork -h
```

### Contribution guide

## Repo structure
```
.
├── cmd
│   ├── configure.go
│   ├── internal
│   │   ├── configure
│   │   │   └── configure.go
│   │   └── logwork
│   │       ├── algorithm.go
│   │       ├── interface.go
│   │       └── jira.go
│   ├── logwork.go
│   └── root.go
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── pkg
│   ├── constant
│   │   └── constant.go
│   └── types
│       ├── configure.go
│       ├── logaction.go
│       ├── logstatus.go
│       └── tickets.go
├── Readme.md
└── Taskfile.yml
```
