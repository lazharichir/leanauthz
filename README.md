# **`leanauthz`** – simple authorization library, clone and use!

## Goal

A quick starter to include in a backend project to get relatively flexible authorization working (that is, authorization and not authentication).

## Non-Goals

* This project does not handle authentication. It assumes that you have already authenticated the user and have a user ID available in the request context. 
* It does not handle authorization for things like UI elements (e.g., buttons, links, etc.). It is only meant to be used for API endpoints.

## High-Level Concepts

Such concepts should be familiar to anyone who has worked with AWS IAM or similar systems.

### Principal

A `Principal` is a string representing an identity (e.g., a user, a service, a role, a team) and is usually formatted as `<namespace>/<identifier>` such as `users/1234` or `roles/admin` or even `accounts/123/teams/sales`. It is used to identify the entity that is requesting access to a resource to perform some action.

### Action

An `Action` is a string representing an operation (e.g., read, write, delete, update, search, list) and is usually formatted as a verb such as `Read` or `Write` or `Delete`. 

The action is used to identify the type of access being requested.

### Resource

A `Resource` is a string representing a resource (e.g., a file, an entity, an API endpoint) and is usually formatted as `<namespace>/<identifier>` such as `files/1234.jpg` or `posts/1234` or even `documents:doc_asMAkd29DnOp`. The resource is used to identify the target being accessed.

### Pattern

A `Pattern` is a string representing a pattern that can be used to match resources, actions, and principals. Patterns are used to declare statements.

A pattern can use the symbol "*" to match anything in its place. For example, `Read*` means "match any action that starts with `Read`" and `users/*` means "match any principal that starts with `users/`".

### Statement

A `Statement` is a tuple of patterns for `Principal`, an `Action`, and `Resource` that represent a single persisted permission. 

For example, a statement might be `teams/analysts`, `Read*`, `invoices/*` which would allow any analyst to read any invoice. That way, you do not have to declare each tuple of user, action, and invoice individually.

### Evaluation

An `Evaluation` is a tuple of a `Principal`, an `Action`, and a `Resource` that represents a single request to access a resource. 

For example, an evaluation might be `users/1234`, `ReadInvoices`, `invoices/1234` which would check whether the user with the ID `1234` is allowed to read the invoice with the ID `1234`. This would be triggered by an API call to `GET /invoices/1234` for example, with the user ID being extracted from the request context (cookies, JWT).

### Outcome

An `Outcome` is an simple object that represents whether an `Evaluation` resulted in an ALLOW or DENY effect. 

By default, if no statement(s) were found to ALLOW an Evaluation, the outcome is DENY. If at least one statement was found to ALLOW an Evaluation, the outcome is ALLOW. If there is a tie, the outcome is DENY. An outcome also has an `Explicit` property which indicates whether the outcome was explicitly declared or inferred as a default DENY.

Finally, there is a `Decider` property which indicates which statement was used to make the decision. This is useful for debugging.

### Store

A `Store` is a simple interface that allows you to persist and retrieve statements.

The default implementation is an in-memory store that we recommend you only use for local development and prototyping. But you can easily implement your own store to use a database or a file system or whatever you want.

```go
type Store interface {
	SaveStatements(ctx context.Context, stmts []Statement) error
	GetStatementsByPrincipal(ctx context.Context, p Pattern) ([]Statement, error)
	DeleteStatements(ctx context.Context, stmts []Statement) error
	FindCandidates(ctx context.Context, e Evaluation) ([]Statement, error)
}
```

The `FindCandidates` method is used to find statements that match an evaluation. It is used to find statements that match an evaluation. For example, if you have a statement that says `users/1234`, `Read*`, `invoices/*` and you have an evaluation that says `users/1234`, `ReadInvoices`, `invoices/1234`, then the statement matches the evaluation and should be returned. 

Because the library is storage-agnostic, you need to adapt that candidate retrieval step to your storage mechanism (e.g., PostgreSQL, MySQL, MongoDB). Regardless of the results returned by your `FindCandidates` implementation, the library will filter out irrelevent candidates out when `Evaluate` is called. 

*In some smaller backends, you may want to have a hard-coded list of all statements that you can return without filtering and let the `Evaluate` function perform the filtering. In larger backends, you may want to use a more complex database query to filter out candidates.*

## Usage

The best way to use this project is to clone it and then copy the root folder into your project (e.g., `./pkg/leanauthz`). That way, you can also bring forward app-specific changes to the codebase (e.g., action and principal validation).

## Improvements

No major improvements are planned at this time. If you have any suggestions, please open an issue or a pull request. If I have the time in the future, I may refactor this in order to make it more flexible and easier to use "out of the box" (e.g., use dependency injection for app-wide action validators).

## FAQ

### Why not provide a simple SQL implementation of the store?

First, I wanted to keep this project as simple as possible and not introduce any dependencies. Then, I also wanted to keep the store interface as simple as possible so that it can be easily implemented in any storage mechanism. Finally, each project has its own way of dealing with SQL (e.g., raw SQL, ORM, etc.) and I did not want to force a specific implementation on anyone. Don't get me started on migrations and schema management.

### Why not use a more complex authorization library?

I've used flavours of this library in a few projects and it has worked well. I've also used more complex libraries like [Casbin](https://casbin.org/) and [Oso](https://www.osohq.com/) and they are great. But they are also more complex and require more setup and configuration. I wanted to create a simple repository that can be used as a starting point for any backend project.

### Is this library production-ready?

I would not use it in a production environment without some modifications. For example, I would add a cache layer to the store to avoid hitting the database on every request. I would also add some logging to the `Evaluate` function to help with debugging. Finally, I would add some more tests to cover edge cases.

### Are you available for consulting?

Yes, I am available for consulting.