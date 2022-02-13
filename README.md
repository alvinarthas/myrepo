# MyRepo
An Experiment Golang Repository for handling REST, RPC, etc. This will be my boilerplate for any project comming next!

# Repository Goals
- Simple, not make anyone confuse how to use the structure
- Max 3 flow (rabbit hole) for the process -> to make sure that the process not taking too much path
- Folder organize well
- As a boilerplate to handle some tech tools

# What you need to prepare
``` bash
latest golang version
latest mysql version
```

# How to Run
``` bash
go mod init
make run // Make Run will init swagger documentation and run the server
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
```

### Business:
- Anything that invovle every business process
- The main business logic will be in the controller
- The domain will be processing part of the process
### Handlers:
- Able to handle REST, RPC, Message Broker, Scheduler, ETC
### Utils:
- As the helpers to support how the system work

# Tech Tools
| Tools | Description | Reference |
| ------ | ------ |------|
| Go-Chi | As Http Routing | https://github.com/go-chi/chi |
| Logrus | Handling Logging in every environtment| https://github.com/sirupsen/logrus |
| Swagger | API Documentation | https://github.com/swaggo/http-swagger |
| MySQL | SQL Database | https://github.com/go-sql-driver/mysql |

# Tech Feature Checklist
- [x] Rest API
- [ ] RPC
- [ ] Message Broker Implementation
- [ ] Database Migration
- [x] MySQL Implementation
- [ ] Postgre Implementation
- [ ] Redis Implementation
- [ ] MongoDB Implementation
- [x] Swagger - API Documentation
- [ ] OAuth2 Implementation
- [ ] JSON Web Token Implementation
- [ ] Docker Implementation
