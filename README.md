# Let's Go!

This project is an initial attempt to work with the Go language. I'll track some process and include
documentation here.

## Setup

To set up, clone this repository and open in VS Code. This project uses a VS Code Dev Container to
install all dependencies needed to develop and run the project.

Next, copy `.env.example` to `.env` in the root folder.

Then, run `Cmd` (Mac) / `Ctrl` (PC) + `Shift` + `P` to open the command palette, then run "Reopen in
Container"

VS Code should pull the necessary Docker images, configure the workspace and reopen the project in a
dev container.

The dev container should contain:

- Postgres (available on port `5432` using default `postgres`/`postgres` login credentials)
- Go Server for backend.
- Nodemon Server for frontend.
- All necessary extensions and scripts for Go and associated linting, testing, etc. To see all
  available scripts, run `npm run` without any arguments.

### Initialize the Database

This setup uses Postgres for the database, along with Prisma for the schema configuration.

1. To initialize the schema, run `npm run prisma:migrate` from the root folder. Allow the `npx`
   installation of `prisma` to proceed.

At the end of a successful setup, you should see a message
`Your database is now in sync with your schema.`.

2. Then run `npm run prisma:generate` to generate the interfaces for types based on the schema.

3. Seed the database by running `npm run prisma:seed`. You should see post and postmeta data in the
   tables:

   <img width="1601" alt="image" src="https://github.com/user-attachments/assets/1993864d-0ecd-4475-8bd5-239b61654729" />


### Starting Backend

To start the backend Go server, run `npm run backend:dev`. The server should start at
`localhost:9000`.

### Starting Frontend

To start the frontend dev server, run `npm run frontend:dev`. The frontend should be available at
`localhost:8080`.

---

## History

### Go Docs

I began this effort by visiting the Go documentation here: https://go.dev/doc/articles/wiki/

While the doc is informative, the project wasn't exactly something I wanted to tackle. I didn't
really want to create/edit/delete and serve pages of content. I'm looking more for an API tutorial.

### VS Code Dev Container

Then, I thought I'd try to experiment with a VS Code Dev Container:
https://github.com/microsoft/vscode-remote-try-go

I went this route because I know from experience that dev containers are an easy way to spin up new
projects with dependencies in Docker, without needing to install them directly on my machine.

### Setting up Let's Go

From the VS Code Dev Container, I decided to pick a few key files and copy them into a new
repository. This would serve as my boilerplate.

### Getting Automated Help

Since I'm new to Go, I figure it makes sense to find a linter, and set up a quick CI process to help
me identify issues faster.

I could also add git commit hooks, but I don't want to spend time on that right now. Next, I'll move
into setting up some basic API endpoints.
