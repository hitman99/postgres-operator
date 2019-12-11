# External PostgreSQL server operator for Kubernetes

## Features
* Creates a database from a CR
* Creates a role with random username and password from a CR
* If the database exist, it will only create a role
* Multiple user roles can own one database
* Creates Kubernetes secret with postgres_uri in the same namespace as CR

## CRs

### Postgres
```yaml
apiVersion: db.movetokube.com/v1alpha1
kind: Postgres
metadata:
  name: my-db
  namespace: app
spec:
  database: test-db # Name of database created in PostgreSQL
  dropOnDelete: false # Set to true if you want the operator to drop the database and role when this CR is deleted
  masterRole: test-db-group
  schemas: # List of schemas the operator should create in database
  - stores
  - customers
```

This creates a database called `test-db` and a role `test-db-group` that is set as the owner of the database.
Reader and writer roles are also created. These roles have read and write permissions to all tables in the schemas created by the operator, if any.

### PostgresUser
```yaml
apiVersion: db.movetokube.com/v1alpha1
kind: PostgresUser
metadata:
  name: my-db-user
  namespace: app
spec:
  role: username
  database: my-db # This references the Postgres CR
  secretName: my-secret
  privileges: OWNER # Can be OWNER/READ/WRITE
```

This creates a user role `username-<hash>` and grants role `test-db-group`, `test-db-writer` or `test-db-reader` depending on `privileges` property. Its credentials are put in secret `my-secret-my-db-user`.

`PostgresUser` needs to reference a `Postgres` in the same namespace.

Two `Postgres` referencing the same database can exist in more than one namespace. The last CR referencing a database will drop the group role and transfer database ownership to the role used by the operator.

## Installation

1. Configure Postgres credentials for the operator in `deploy/operator.yaml` 
2. `kubectl apply -f deploy/crds/db.movetokube.com_postgres_crd.yaml`
3. `kubectl apply -f deploy/crds/db.movetokube.com_postgresusers_crd.yaml`
4. `kubectl apply -f deploy/namespace.yaml`
5. `kubectl apply -f role.yaml`
6. `kubectl apply -f role_binding.yaml`
7. `kubectl apply -f service_account.yaml`
8. `kubectl apply -f operator.yaml`

