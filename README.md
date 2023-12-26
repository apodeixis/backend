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

We use [openapi-generator](https://github.com/OpenAPITools/openapi-generator).
You can generate model bindings for Go based on docs using `generate.sh` script.

```bash
  ./generate.sh --help
```