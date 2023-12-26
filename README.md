## backend

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This is a primary API gateway for the client.

### Documentation

[OpenAPI 3.0.0](https://spec.openapis.org/oas/v3.0.0) standard.

Build and start local documentation:

```bash
cd docs
npm i
npm run build & npm run start
```
In case everything is successful, you will see the following links:

    ✔ Documentation (ReDoc):      http://localhost:8080
    ✔ Documentation (SwaggerUI):  http://localhost:8080/swagger-ui/
    ✔ Swagger Editor:             http://localhost:8080/swagger-editor/

### Codegen

#### Swagger

We use [openapi-generator](https://github.com/OpenAPITools/openapi-generator).
You can generate model bindings for Go based on docs using `generate.sh` script.

```bash
./generate.sh --help
```

#### EVM

This service interacts with Solidity contract `Posts.sol` from [contracts](https://github.com/apodeixis/contracts).

In `pkg/` directory you can find generated Go bindings for it. To obtain them yourself, at first you need ABI - 
check the reference section in [contracts](https://github.com/apodeixis/contracts).

Then, utilize [abigen](https://geth.ethereum.org/docs/tools/abigen):

```bash
abigen --abi=Posts.abi --pkg=posts --out=pkg/posts/posts.go
```