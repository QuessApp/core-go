# Core API

#### Where the magic happens 🪄

Core API is the entry point to every action. If we want to send a new question, Core API will send it and so on.

## Before you start

1 - This project follows [this project layout](https://github.com/golang-standards/project-layout).<br/>
2 - I guess that would be cool if you know that Quess don't have only this project.
We have others projects, you can visit them at:

- [Toolkit - Helpers modules to services](https://github.com/QuessApp/toolkit)
- [Email Service - Service to send email using SES](https://github.com/QuessApp/email-service)
- [Geo Service - Service to fetch user location by IP](https://github.com/QuessApp/trusted-geo-service)
- [Ban Service - Service to ban users](https://github.com/QuessApp/ban-service)
- [Web App - Our beautiful Web UI!](https://github.com/QuessApp/web-app)

[You can visit all projects here](https://github.com/orgs/QuessApp/repositories)

3 - This project is written in GO, but in the past we developed it using JavaScript. [Check it out!](https://github.com/QuessApp/core)

## Tech stack

- GO
- Gofiber
- RabbitMQ
- AWS SES & AWS S3
- MongoDB
- Swagger
- Testify
- JWT <br/>
  ...

## Run Locally

Clone the project

```bash
$  git clone https://github.com/QuessApp/core-go
```

Go to the project directory

```bash
$  cd core-go
```

Run commands to start the project:

```bash
$ make start
```

or

```bash
$ ./scripts/start.sh
```

When you run the above command, it will perform some actions, such as:

- Will check if `.env` file exists
- Copy `.env.example` file to `.env` if doesn't exist yet
- Run containers from `docker-compose.yml` file (you need to run Docker on pc)
- Finally, start the project based in `ENV` property from `.env` file

If you want to destroy everything, you can run the following command:

```bash
$ make destroy
```

or

```bash
$ ./scripts/destroy.sh
```

When you run the above command, it will perform some actions, such as:

- Delete all containers created previously
- Delete `tmp` folder

## Roadmap

- Write more tests

- New features

- Improve Swagger docs

## Contributing

Contributions are always welcome!

See [contributing.md](https://github.com/QuessApp/core-go/blob/master/.github/CONTRIBUTING.md) for ways to get started.

Please adhere to this project's code of conduct.

## Authors

- [Caio Augusto (owner & maintainer)](https://www.github.com/caioaugustoo)

## Support

For support, email caioamfr@gmail.com

## License

[MIT](https://choosealicense.com/licenses/mit/)
