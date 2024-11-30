# Utah Go Meetup - December 3, 2024

Demo application for Utah Go Meetup that contains a modified version of the
[Order Management System reference application](https://github.com/temporalio/reference-app-orders-go).

## Temporal Reference Application: Order Management System (Go)

![OMS logo](docs/images/oms-logo.png)

The Order Management System (OMS) is a reference application that 
demonstrates one way to approach the design and implementation of 
an order processing system based on Temporal Workflows. You can run 
this application locally (directly on a laptop) or in a Kubernetes 
cluster. In addition, the required Temporal Service can be run locally, 
or be provided by a remote self-hosted deployment, or be provided by 
Temporal Cloud. 

### Required Software
You will need [Docker](https://www.docker.com/) to run the core OMS application, 
and the [Temporal CLI](https://docs.temporal.io/cli#install) to run the 
Temporal Service locally. For local development, this repository uses [omni](https://omnicli.dev/).


### Start the Temporal Service
Run the following command in your terminal:

```command
temporal server start-dev  \
        --db-filename temporal-persistence.db \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true"
```

The Temporal Service manages application state by assigning tasks
related to each Workflow Execution and tracking the completion of 
those tasks. The detailed history it maintains for each execution 
enables the application to recover from a crash by reconstructing 
its pre-crash state and resuming the execution.

### Register Search Attributes
Run the following commands in a new terminal:

```command
temporal operator search-attribute create --name CustomerId --type Keyword
temporal operator search-attribute create --name FulfillmentId --type Keyword
temporal operator search-attribute create --name OrderId --type Keyword
```


### Start the Remaining Services
Run the following command in your terminal:

```command
docker-compose up
```


## Find Your Way Around
This repository is organized as follows:

| Directory | Description |
| --- | --- |
| <code><a href="app/">app/</a></code> | Application code |
| <code><a href="cmd/">cmd/</a></code> | Command-line tools provided by the application |
| <code><a href="docs/">docs/</a></code> | Documentation |
| <code><a href="gen/">gen/</a></code> | Generated Go Code |
| <code><a href="proto/">proto/</a></code> | Protobuf Definitions |
| <code><a href="web/">web/</a></code> | Modified version of [temporalio/reference-app-orders-web](https://github.com/temporalio/reference-app-orders-web)] |


See the [documentation](https://github.com/temporalio/reference-app-orders-go) for more information.
