## Installation

1. Clone the repository
   ```bash
   git clone git@github.com:karuhun-developer/goravel-template.git
   ```
2. Change directory
   ```bash
   cd goravel-template
   ```
3. Install dependencies
   ```bash
   go mod tidy
   ```
4. Copy `.env.example` to `.env` and modify the configuration as needed
   ```bash
   cp .env.example .env
   ```
5. Chmod the ./air.sh
   ```bash
   chmod +x ./air.sh
   ```
6. Generate the application key
   ```bash
   go run . artisan key:generate
   ```
7. Generate the jwt secret key
   ```bash
   go run . artisan jwt:secret
   ```
8. Run database migrations
   ```bash
   go run . artisan migrate
   ```
9. Run database seeders (optional)
   ```bash
   go run . artisan db:seed
   ```
10. Run the application
    ```bash
    ./air.sh
    ```

## Documentation

Online documentation [https://www.goravel.dev](https://www.goravel.dev)

Example [https://github.com/goravel/example](https://github.com/goravel/example)

> To optimize the documentation, please submit a PR to the documentation
> repository [https://github.com/goravel/docs](https://github.com/goravel/docs)

## Main Function

|                                                                                        |                                                                 |                                                                          |                                                                       |                                                                                |
| -------------------------------------------------------------------------------------- | --------------------------------------------------------------- | ------------------------------------------------------------------------ | --------------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| [Config](https://www.goravel.dev/getting-started/configuration.html)                   | [Http](https://www.goravel.dev/the-basics/routing.html)         | [Authentication](https://www.goravel.dev/security/authentication.html)   | [Authorization](https://www.goravel.dev/security/authorization.html)  | [Orm](https://www.goravel.dev/orm/getting-started.html)                        |
| [Migrate](https://www.goravel.dev/database/migrations.html)                            | [Logger](https://www.goravel.dev/the-basics/logging.html)       | [Cache](https://www.goravel.dev/digging-deeper/cache.html)               | [Grpc](https://www.goravel.dev/the-basics/grpc.html)                  | [Artisan Console](https://www.goravel.dev/digging-deeper/artisan-console.html) |
| [Task Scheduling](https://www.goravel.dev/digging-deeper/task-scheduling.html)         | [Queue](https://www.goravel.dev/digging-deeper/queues.html)     | [Event](https://www.goravel.dev/digging-deeper/event.html)               | [FileStorage](https://www.goravel.dev/digging-deeper/filesystem.html) | [Mail](https://www.goravel.dev/digging-deeper/mail.html)                       |
| [Validation](https://www.goravel.dev/the-basics/validation.html)                       | [Mock](https://www.goravel.dev/testing/mock.html)               | [Hash](https://www.goravel.dev/security/hashing.html)                    | [Crypt](https://www.goravel.dev/security/encryption.html)             | [Carbon](https://www.goravel.dev/digging-deeper/helpers.html)                  |
| [Package Development](https://www.goravel.dev/digging-deeper/package-development.html) | [Testing](https://www.goravel.dev/testing/getting-started.html) | [Localization](https://www.goravel.dev/digging-deeper/localization.html) | [Session](https://www.goravel.dev/the-basics/session.html)            |                                                                                |
