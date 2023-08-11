⚠️WIP⚠️
# GSLoC Go SDK

GSLoC is a open source Global Server Load Balancing (GSLB) solution.

GSLoC go sdk is an api generated of [lbaas-api](https://github.com/orange-cloudfoundry/gsloc-api). 

It has a client implementation for interacting with GSLoC with authentication and skeleton to make server (used in [gsloc](https://github.com/orange-cloudfoundry/gsloc))

## Other GSLoC repositories 

- [gsloc](https://github.com/orange-cloudfoundry/gsloc): GSLoC server implementation.
- [gsloc-cli](https://github.com/orange-cloudfoundry/gsloc-cli): GSLoC cli for interacting with server.
- [gsloc-api](https://github.com/orange-cloudfoundry/gsloc-api): Api definition in protobuf and grpc with code generation tools and architectural docs and doc api.
- [gsloc-go-sdk](https://github.com/orange-cloudfoundry/gsloc-go-sdk): Go sdk for gsloc api used in gsloc implementation.

## Pre-Commit Tool

We use the `pre-commit` tool to automatically enforce coding standards and run various checks on code changes. To get
started, follow these steps:

1. Install `pre-commit` globally on your local machine. You can do this by running the following command in your
   terminal (prefer using pip3):

```bash
pip install pre-commit
```

2. After installing `pre-commit`, navigate to the root directory of the project and run the following command to install
   the pre-commit hooks:

```bash
pre-commit install
```

3. Now, every time you make a commit, `pre-commit` will automatically run the configured checks on your changes and
   prevent the commit if any issues are found.
