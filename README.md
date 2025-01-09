# temporal-terraform-auth0

Template temporal workflow to manage auth0 resources using terraform.

To run this:

1. Run Temporal using docker
```console
git clone https://github.com/temporalio/docker-compose.git
cd docker-compose
docker compose up
```
2. Clone this repository
```console
$ git clone https://github.com/mlgupta/temporal-terraform-auth0
```
3. Make appropriate changes to terraform credentials for auth0.
4. Terraform config uses HCP to store statefiles, make sure to configure that first.
5. Apply workflow
```console
$ go run worker/main.go
```
6. Trigger workflow
```console
$ go run start/main.go
```

