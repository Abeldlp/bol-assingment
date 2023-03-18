<p align="center">
  <img
    src="assets/golang.png"
    height="150px"
  />
</p>

## International Space Station app

This project gets the current location from the international space station api, and stores it in a database for later use.

The whole purpose is to store the current location, together with the timezone, to be able to check on a specific date at a specific hour, how many minutes were spent by the ISS on a specific Timezone.

- [Application Structure](#application-structure)
- [Application Dataflow](#application-dataflow)
- [Running app locally](#running-app-locally)
- [Application testing](#application-testing)
- [Application usage](#application-usage)
- [Author message](#author-message)

### Application Structure

This application is structured by:

- **Cron Service (Go)** _A service which purpose is to run cronjobs_
- **Api Service (Go)** _A web server to be consumed by client application_
- **Mail Service (Go)** _A service that takes care of sending emails_
- **Api Gateway (Go)** _An entry point to the services_
- **Client (Vue)** _An application the end user can interact with_
- **RabbitMQ** _Message broker_
- **PostgreSQL Database** _Database to store ISS location_

For local development all the applications run in docker containers and it is orchestrated via docker-compose

### Application Dataflow

<p align="center">
  <img
    src="assets/app_design.png"
  />
</p>

1. The Cron Service pulls the ISS's location every minute from the iss api and it is converted to a schema needed for our research.
2. Once the data is converted, the Cron Service stores the data into the database, this proccess only goes one way and cron service is only responsible for writing data to the DB.
3. After the data is saved, the Cron Service queues the timezone into a rabbitMQ queue.
4. The Mail Service consumes the Notification Queue and on a receive event, it triggers the notification functionality.
5. The Mail Service queries the database to find subscribed users based on the timezone recieved.
6. After the users are found, it sends a notification to all of them.
7. The Api Service can get and query the ISS location data.
8. The Api Serice can also write subscribed users to the database.
9. The Api Service recieves all the calls via the api gateway.
10. The Client Vue application is responsible for displaying datatables with the data and create subscribers.
11. The end user is capable to get satelite data and post it's email to our database.

### Running app locally

> To be able to run the application you need docker and docker-compose installed. Please refer to the [docs](https://docs.docker.com/compose/install/) to install them before taking any further steps

For the Mail service to work add your email credentials to `.env.docker`, the emails will be sent from the specified account.

```
MAIL_HOST="smtp.gmail.com"
MAIL_DOMAIN_ADDRESS=example@mail.com
MAIL_PASSWORD=mypassword
```

Clone this repository

```bash
git clone https://github.com/Abeldlp/iss.git
```

Go to the directory

```bash
cd iss
```

Build docker images (cup of coffee)

```bash
docker-compose build
```

Run all containers in detached mode

```bash
docker-compose up -d
```

At this point you will have five containers running and you can access the client on [http://localhost:3000](http://localhost:3000)

<p align="center">
  <img
    src="assets/preview.png"
  />
</p>
If you want to access the database container use a client of your liking, and connect to postgres with the following.

- USERNAME `user`
- PASSWORD `password`
- HOST `localhost`
- PORT `5432`
- DATABASE `app-db`

If you want to play with the api, the api gateway is listening on [http://localhost:4444](http://localhost:4444)

To stop the containers you can run the following

```bash
docker-compose down
```

> Optionaly if you have `make` and `build-essentials` installed locally you can make use of the `Makefile`

### Application testing

To run all tests at once you can make use of the Makefile. You will need `make` and `build-essentials` installed locally. Refer to your OS/Distro installation docs.
Before running any tests don't forget to stop all the containers, since client e2e testing is using the same port mapping to serve with mock and run tests.

```bash
make test
```

The above command will run api-service tests, cron-service tests, client unit tests and client e2e tests.

To run them individually you can run the following:

**Api Service**

```bash
cd api-service && go test ./test
```

**Cron Service**

```bash
cd schedule-service && go test ./test
```

**Client (unit)**

```bash
cd client && npm run test:unit:ci
```

**Client (e2e)**

```bash
cd client && npm run test:e2e:ci
```

For the client's e2e testing you can also run them in preview mode, so you can visualize what the tests are doing on a Chromium, Electron or Firefox browser.
You can run them with the following command:

```bash
cd client && npm run test:e2e
```

In case of client project showcase, it is also possible to run the project in mocking mode, which will mock api calls and use static data. The project will be served on [http://localhost:3000](http://localhost:3000)
You can serve the project with mock locally with the following command:

```bash
cd client && npm run dev-mock
```

### Application usage

The client application consists on:

- Email subscription
- Autofetch button.
- Aggregated Datatable
- Sample Datatable
- Informational tree

**Email subscription**

By subscribing to the email service, you save the email and timezone in the database.
Everytime a new ISS location is added to the database, the message broker will queue the timezone for the mail service to recieve. If the ISS is on you timezone it will send you an email to notify you.

**Autofetch button**

This button allows the user to fetch data automatically every minute so page reload is not required to fetch data.

**Aggregated Datatable**

Datatable containing the data aggregated. You are able to paginate, sort and filter using the selectors.

**Sample Datatable**

Datatable containing all the data stored in the database. You are able to paginate and sort.

**Informational Tree**

You can always check the explanation of the page parts in case you get lost.

### Author message

This application is simple to use and straight to the point. If you wish to develop locally remember that you will need Golang and Node installed.

If any questions come your way don't hesitate to drop an email to <abel45991690@gmail.com>

<p align="center">
  <img
    src="https://raw.githubusercontent.com/catppuccin/catppuccin/dev/assets/footers/gray0_ctp_on_line.svg?sanitize=true"
  />
</p>
