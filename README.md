# Golang microservices assignment

**Congratulations for making it into the next round of our interview!**

This assignment is meant to prove your golang proficiency or how fast can you learn new concepts. :)

Your code structure should follow microservices best practices and our evaluation will focus primarily on your ability to follow good design principles and less on correctness and completeness of algorithms. During the face to face interview you will have the opportunity to explain your design choices and provide justifications for the parts that you omitted.

## Evaluation points in order of importance

- use of clean code, which is self documenting
- use of packages to achieve separation of concerns
- use of domain driven design
- use of golang idiomatic principles
- use of docker
- tests for business logic
- use of code quality checks such as linters and build tools
- use of git with appropriate commit messages
- documentation: README.md and inline code comments
- you must use go modules and a version of go >= 1.16

Please avoid using frameworks such as `go-kit` and `go-micro` since one of the purposes of the assignment is to evaluate the candidate ability of structuring the solution in their own way.
If you have questions about the test, please draw your own conclusions.


## Instructions

- Given a JSON file with ports data (ports.json), write 2 services (Client API, PortDomainService)
- Client API should parse the JSON file and have REST interface
- This JSON file is of unknown size, it can contain several millions of records (i.e. 20Gb+ in size, which can't be read at once)
- Client API has limited resources available (e.g. 200MB ram)
- While reading JSON file, Client API calls PortDomainService, that either creates a new record in a database, or updates the existing one
- PortDomainService should store the data in a database that contains ports, representing the latest version found in the JSON. 
- Database can be Map in memory
- Client API should provide an endpoint to retrieve the data from the PortDomainService
- Each service should be built using Dockerfile
- Use gRPC as a transport between services

Choose the approach that you think is best (i.e. most flexible).

## Bonus points

- Database in docker container
- Domain Driven Design
- Docker-compose file

## Note
We are looking for the ClientAPI (the service reading the JSON) to be written in a way that is easy to reuse, give or take a few customisations.
The services themselves should handle certain signals correctly (e.g. a TERM or KILL signal should result in a graceful shutdown).

**Important:** Please fork this repository, create a branch in your fork and make commits there. The commit messages have to follow our styleguide (see [this commits](https://github.com/bleenco/bproxy/commits/master) for the reference). When you are done, please submit your solution by creating Pull Request into master branch of the main repository. Please make sure to add any installation changes to `README.md` and the amount of time (hours / minutes) spent on the task. It is expected that both solutions contain sufficient code documentation and test coverage.
Good luck! 
