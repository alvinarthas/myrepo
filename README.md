# MyRepo
An Experiment Golang Repository for handling REST, RPC, etc.

# Repository Goals
```bash
    Simple, not make anyone confuse how to use the structure
    Max 3 flow (rabbit hole) for the process -> to make sure that the process not taking too much path
    Folder organize well
```

# What you need to prepare
``` bash
    latest golang version
    latest mysql version
```

# How to Run
``` bash
    go run main.go
```

# Structure
``` bash
    - business
        - controller
        - domain
        - model
    - config
    - connection
    - docs
    - handlers
    - utils

Business:
    - Anything that invovle every business process
    - The main business logic will be in the controller
    - The domain will be processing part of the process
Handlers:
    - Able to handle REST, RPC, Message Broker, Scheduler, ETC
Utils:
    - As the helpers to support how the system work
```
