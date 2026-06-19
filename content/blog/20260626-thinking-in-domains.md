---
date: '2026-06-26T09:58:05-06:00'
draft: true
featured: true
title: 'Thinking about Domains: Clarity Leads to Evolution'
description: >
  A well organized and structured application divides its logical features into domains. These domains are responsible
  for their own logic and state, while your system depends on these domains to provide business value.
image: /images/posts/20260626/cover.jpg
categories: 
  - engineering
  - architecture
---
Building software that evolves over time can be challenging, so much so that people have written many different patterns
on the best practices for organizing your application. It is a fact that a well organized piece of software can make it
much easier to iterate upon, adding new features and flexibility to integrations.

One of the ideal solutions you can think of is having a solution that is essentially "glue-code" that connects different
bits together.

This can be achieved if you design you application in terms of **domains**.

By structuring your application in terms of domains, you will leave all the logic, state management and business rules be
handled by the domain itself, while the "application" code interacts with these domains to achieve some desired behavior.
You can see how thinking like this reduces the amount of dependencies and sets hard boundaries, which in turn promotes
a very modular, extensible and maintainable codebase.

In this article I'll explore how to think in domains, the rules that you must follow for healthy implementations and 
understand the Domain Driven Design (DDD) principles, briefly flying over some of the derivative architectures.

## What is a Domain?

I like to define a domain as this:

> "A domain is a single logical unit of functionality that encapsulates a set of related business rules and operations."

Domains are responsible for all the logic of operating on entities they own, and they all export
a well-defined API that applications or other domains can use to interact with them.

I say "other domains" above because in an application, domains may depend on other domains to define higher-level
functionality. In the e-commerce example above, you can imagine that `Order` can depend on `Product` and `Customer` to
define the business rules for creating and processing orders. This establishes "dependency directions" of these domains.
We will see more on this later.

## The four layers of a Domain

Eric Evans in his book ["Domain-Driven Design" (2003)](https://www.goodreads.com/en/book/show/179133.Domain_Driven_Design)
described the principles of DDD to structure applications. One of the key concepts introduced is the idea of a four layer
architecture for applications.

{{<figure
  src="/images/posts/20260626/ddd.jpg"
  alt="DDD Four Layers"
  caption="Four layers of an application in DDD"
  width="100%"
>}}

The four layers are, in descending order:

- **User Interface**: handles user input and formats it as payload to services in the application layer and encodes 
                      display logic for the user.
- **Application**: defines data transfer objects (DTOs) and uses it to communicate with components in the domain layer.
- **Domain**: implements the core business logic and interacts with infrastructure (e.g. DB) entities. This is the _core_
              layer for the business.
- **Infrastructure**: provides the technical implementation details, such as database access and external service integrations.

As you can see, this separation lifts a lot of cognitive load from understanding your application, so it is obvious
which pieces do what. This separation also promotes code hygiene and separate concerns between layers. If you are not
mixing UI with DB models, it is much easier to iterate and maintain your code.

### Example of a domain implementation

Suppose an e-commerce application

- `GET /orders` returns a list of orders. The view logic implements the **user interface** layer. This handler calls
- `OrderService` to retrieve the list of orders in data transfer objects (DTOs) on the **application** layer. In turn, it calls
- `OrderRepository` to retrieve the list of orders from the **domain** layer. This repository encodes which orders can
  be retrieved by the user, has logic on filters defined by product requirements etc. The repository finally calls the
- `OrderModel` to actually retrieve the orders from the configured database. This is the **infrastructure** layer.

## Tenets of Domain-Driven Design

Let us see some of the key principles of Domain-Driven Design that you need to understand to effectively apply it to your 
own projects.

### Model domains according to logical boundaries

Think about your application high level components and how they interact with each other, identify where logical boundaries
exist in your mental model and start scratching how them play together.

For example, if you have an e-commerce application, several domains come to mind:

- **Product**: Manages the lifecycle of products, including inventory, pricing, and descriptions.
- **Order**: Handles the creation, processing, and fulfillment of orders.
- **Cart**: Manages the shopping cart, including adding, removing, and updating items.
- **Customer**: Manages customer information, including profiles and preferences.
- **Payment**: Processes payments and manages transaction records.

Remember, each of these domains is responsible for all the logic of operating on entities they own, which brings us to the next step.

### Domains interact with each other via application-level or UI-level APIs

Domains are an abstraction of a business purpose, so imports should **never** use internal entities and functionality loosely.
Instead, domains are expected to expose a well-defined API that the outside world can use in a standard manner.

Suppose your `Cart` domain needs to list the `Product`s in that cart instance. Instead of letting the `Cart` query the `Product`
database directly, you'd use a `ProductService` application-level API to fetch the products for you, without exposing internals
to `Cart`. Similarly, if you are playing with a microsservices mesh, you'd interact with a gRPC or HTTP interface of `Product`
to get a list of objects.

Always remember, **domains abstract away details**, use the defined APIs to interact with other domains.

### Dependency arrow should point from generic to specific

{{<figure 
  src="/images/posts/20260626/imports.svg" 
  caption="e-commerce dependency diagram"
  alt="e-commerce dependency diagram"
  width="70%"
>}}

When drawing a dependency diagram between domains, think of what is more generic and isolated and what is more specific.
Think in your mind about how the domains interact with each other, does it make sense to for product to depend on the cart?
Is the payment domain dependent on the customer, or is the customer a reference in payment?

After some rounds of thinking, you should have a good understanding of the system dependencies between domains.

The image above shows a possible solution to the e-commerce problem. Here:
- `Payment` can live on its own, so it is a "level-1" type domain. 
- `Order` depends on `Payment`, so it is a "level-2" domain. 
- `Cart` depends on `Product` and needs to know about an order to list the products used in the cart for a purchase. 
- And so on.

You can see that the dependency arrow flows in one direction: from generic to specific. If we were to reverse the direction for
some imports, we'd introduce circular dependencies that makes the system unstable and harder to maintain.

`Core components` in the diagram are just generic helpers and abstractions your business logic have to interact with, say,
infrastructure, external services, etc. Think of it as a "level-0" domain, that is pure in its essence and **does not depend
on any other domain**.

**Note:** Obviously this is not the most optimal design, that depends on your business and product requirements, but the example
is sufficient to illustrate the point.

## Thinking in domains

To summarize, whenever you are modeling domains, keep this close to your heart:

- Separate logical capabilities of your application
- Think how domains interact with each other
- Define a dependency graph on these domains, following the natural flow of dependencies

In addition to these, borrow Evans' advice too and **Model Out Loud**

> "One of the best ways to refine a model is to explore with speech, trying out loud various constructs
   from possible model variations. Rough edges are easy to hear.
   
> - “If we give the Routing Service an origin, destination, and arrival time, it can look up the stops
the cargo will have to make and, well . . . stick them in the database.” **(vague and technical)**

> - “The origin, destination, and so on . . . it all feeds into the Routing Service, and we get back an
Itinerary that has everything we need in it.” **(more complete, but verbose)**

> - “A Routing Service finds an Itinerary that satisfies a Route Specification.” **(concise)**

> (Domain-Driven Design 2003, page 19)

He argues that our language is powerful and phrasing it in different formats helps you identify your
domain boundaries, dependencies and responsibilities faster.

## Conclusion

Should you adopt a Domain-Driven approach to model your business application? I would argue that in general, yes.
Thinking in domains unlocks your ability to understand the problem and design modular, flexible and maintainable logical
units that operates independently, but exactly as logic would dictate.

The layer separation within a domain helps separating concerns on the different aspects a domain have and your craft flows
naturally from these premises. This makes your entire codebase predictable and super easy to evolve.

If domains are independent, adding new features and managing them mean:

- Adding a new independent domain and introducing domain dependencies using a simple API
- Testing is also easy, as modules can be tested independently and mocked on boundaries
- Deploying can also be easy, as domains can be deployed independently and scaled separately
- Changing business logic inside a domain means changing code in a single place, inside the domain layer

If anything, remember that code understanding is key to maintaining and evolving your codebase.

However, if modeling an application in domains is adding a lot of friction with frameworks, for example, consider simplifying
the assumptions and strip some of the domain layers. Perhaps a framework deals with infrastructure in a standard way so that
trying to separate layers is just too costly. **Prefer the idiomatic way** of your language and framework, but **always try to
keep some domain mental model of your application in mind**.
