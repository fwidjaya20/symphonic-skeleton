# Symphonic Skeleton

## Developer Guide

### Installation

```shell
// Download the skeleton
git clone git@github.com:fwidjaya20/symphonic-skeleton.git [project-name] && rm -rf [project-name]/.git*

// Installation dependencies
cd [project-name] && go mod tidy

// Rename the module path
find . -type f \( -name "*.go" -o -name "go.mod" -o -name "go.sum" \) -exec sed -i '' 's|github.com/fwidjaya20/symphonic-skeleton|[your-module-path]|g' {} +

// Create .env environment configuration file
cp .env.example .env
```

#### Development

Supercharge your development workflow with these handy libraries. Install them effortlessly to streamline your coding process:

```shell
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.15.0
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/segmentio/golines@latest
go install github.com/bombsimon/wsl/v4/cmd...@master
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Enhance your project's code consistency by automating the application of Prettier formatting using GitHooks. Follow these steps to set it up:

```shell
./scripts/install-git-hooks.sh
```

## Architecture

### [Domain-Driven Design](https://github.com/ZilvinasKucinskas/FRP-EventSourcing/blob/master/sources/xx735.Eric.Evans.Domaindriven.Design.Tackling.Complexity.in.the.Heart.of.Software.pdf)

DDD proposes a Layered Architecture as a way to separate concerns and avoid responsibility confusion by splitting the codebase in 4 main layers: **User Interface**, **Application**, **Domain**, and **Infrastructure**.

The main rule here is that components in each layer should depend only on components in the same layer or any layer beneath it. Upper layers can use components of lower ones just by calling their public interfaces and lower layers can only communicate upwards by Inversion of Control (IoC).

- **User Interface Layer**: is responsible for displaying data and capturing user’s commands. (Views, Controllers, CLI, etc.)
- **Application Layer**: serves as an orchestrator of domain work, it does not know domain rules, but organizes and delegates domain objects to do their job. It is also the only layer accessible to other bounded contexts. (Commands, Command Handlers, etc.)
- **Domain Layer**: holds the business logic and rules, as well as the business state. It is where the Domain Model lives. (Interfaces, Entities, Value Objects, Aggregates, Domain Services, etc.)
- **Infrastructure Layer**: implements all the technical functionalities the application needs to support the higher layers, persistence, messaging, communication between layers, etc... (Repositories, Caching, External Integrations, etc.)

Even though not all layers are required in every system, the existence of the Domain Layer is a prerequisite in DDD.

**Resources**:
- [Tackling Complexity in the Heart of Software](https://github.com/ZilvinasKucinskas/FRP-EventSourcing/blob/master/sources/xx735.Eric.Evans.Domaindriven.Design.Tackling.Complexity.in.the.Heart.of.Software.pdf)
- [Domain-Driven Design: Everything You Always Wanted to Know About it, But Were Afraid to Ask](https://medium.com/ssense-tech/domain-driven-design-everything-you-always-wanted-to-know-about-it-but-were-afraid-to-ask-a85e7b74497a)
- [Vertical Software Development](https://medium.com/ssense-tech/vertical-software-development-495b73f7fcdf)
- [DDD Beyond the Basics: Mastering Multi-Bounded Context Integration](https://medium.com/ssense-tech/ddd-beyond-the-basics-mastering-multi-bounded-context-integration-ca0c7cec6561)

**Example Implementation**:
- [Symphonic DDD Example](https://github.com/fwidjaya20/symphonic-example)