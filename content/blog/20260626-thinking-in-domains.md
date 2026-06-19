---
date: '2026-06-26T09:58:05-06:00'
draft: true
featured: true
title: 'Thinking in Domains: Clarity Leads to Evolution'
description: >
  A well-organized application divides business capabilities into domains with clear ownership, boundaries and
  dependencies. Learn how domain thinking makes software easier to understand and evolve.
image: /images/posts/20260626/cover.jpg
categories:
  - engineering
  - architecture
---
Building software that evolves over time can be challenging, so much so that people have written many patterns and
practices for organizing applications. A well-organized piece of software can make it much easier to iterate on, add new
features and keep integrations flexible.

One ideal solution is to have an application that is essentially "glue code" connecting different pieces together.

This can be achieved if you design your application in terms of **domains**.

By structuring your application in terms of domains, you let each domain handle its own logic, state management and
business rules, while the application code coordinates these domains to achieve some desired behavior. Thinking this way
reduces unnecessary dependencies and creates explicit boundaries, which promotes a modular, extensible and maintainable
codebase.

In this article, I'll explore how to think in domains, which principles support healthy implementations and how these
ideas connect to Domain-Driven Design (DDD) and related architectures.

## What is a Domain?

I like to define a domain as this:

> "A domain is a single logical unit of functionality that encapsulates a set of related business rules and operations."

In this article, I use "domain" in a practical implementation sense: a cohesive business capability with explicit
ownership and boundaries. In stricter DDD terminology, some of these units might be described as subdomains or bounded
contexts.

Domains are responsible for the logic that operates on the entities they own, and each exposes a well-defined API that
applications or other domains can use to interact with it.

I say "other domains" because domains may depend on each other to define higher-level functionality. In an
e-commerce application, for example, `Order` may depend on `Product` and `Customer` concepts to define the business
rules for creating and processing orders. This establishes dependency directions between domains. We will see more on
this later.

## The four layers of a DDD application

In his book
["Domain-Driven Design" (2003)](https://www.goodreads.com/en/book/show/179133.Domain_Driven_Design), Eric Evans
described principles for structuring applications around a domain model. One of the concepts he presents is a
four-layer architecture for applications.

{{<figure
  src="/images/posts/20260626/ddd.jpg"
  alt="DDD Four Layers"
  caption="Four layers of an application in DDD"
  width="100%"
>}}

The four layers are, in descending order:

- **User Interface**: receives user input, calls application use cases and presents results.
- **Application**: coordinates use cases, transactions and interactions between domain objects or services. It may also
  map data transfer objects (DTOs), but it does not implement the core business rules.
- **Domain**: contains the core business concepts and rules, including entities, value objects, aggregates and domain
  services.
- **Infrastructure**: implements technical concerns such as persistence, messaging and external service integrations.

This separation removes a lot of cognitive load from understanding an application because it makes each responsibility
clear. It also promotes code hygiene and separation of concerns. If you are not mixing UI code with database models, it
is much easier to iterate on and maintain your code.

### Example of a domain implementation

Suppose an e-commerce application exposes `GET /orders`:

1. The HTTP handler belongs to the **user interface** layer. It validates the request and calls an application use case.
2. `ListOrdersService` belongs to the **application** layer. It coordinates authorization, pagination and the retrieval
   operation.
3. The service depends on an `OrderRepository` interface or query contract defined near the application or domain layer.
4. An infrastructure implementation uses an ORM model or database client to retrieve the data.
5. Domain policies and objects enforce business rules when those rules are relevant to the operation.

For a simple read operation, a query service may return DTOs from a dedicated read model. Not every query needs to
reconstruct a complete domain aggregate.

## Tenets of Domain-Driven Design

Let us look at some of the key principles you need to understand to apply domain thinking effectively to your own
projects.

### Model domains according to logical boundaries

Think about your application's high-level components and how they interact with each other. Identify where logical
boundaries exist in your mental model and start sketching how they fit together.

For example, if you have an e-commerce application, several domains come to mind:

- **Product**: manages the lifecycle of products, including inventory, pricing and descriptions.
- **Order**: handles the creation, processing and fulfillment of orders.
- **Cart**: manages the shopping cart, including adding, removing and updating items.
- **Customer**: manages customer information, including profiles and preferences.
- **Payment**: processes payments and manages transaction records.

Remember that each of these domains is responsible for the logic that operates on the entities it owns, which brings us
to the next point.

### Domains interact through explicit contracts

Domains abstract a business purpose, so other modules should not import their internal entities or call internal
functionality directly. Instead, each domain should expose a well-defined contract that the outside world can use in a
controlled manner.

Suppose your `Cart` domain needs information about the products in a cart. Instead of letting `Cart` query the
`Product` database directly, you can use a `ProductService` application-level API or a public interface to fetch the
required information without exposing `Product` internals. In a microservices architecture, the equivalent contract
might be a gRPC or HTTP API.

Other valid contracts include events, ports and message-based integrations. The implementation depends on your
architecture, but the principle is the same: **domains abstract away details, so use their defined contracts**.

### Dependencies should point toward stable, well-defined boundaries

{{<figure
  src="/images/posts/20260626/imports.svg"
  caption="e-commerce dependency diagram"
  alt="e-commerce dependency diagram"
  width="70%"
>}}

When drawing a dependency diagram between domains, think about which capabilities are more stable and isolated and which
ones coordinate or build on top of them. Ask how the domains interact with each other. Does it make sense for `Product`
to depend on `Cart`? Does `Payment` depend on an `Order`, or should it receive a payment request through an explicit
contract?

After a few rounds of modeling, you should have a clearer understanding of the system's dependencies and the boundaries
between responsibilities.

A possible relationship between the e-commerce domains is:

- `Cart` depends on product information.
- `Order` depends on product and customer information.
- `Payment` depends on an order or payment request.
- An application-level checkout service coordinates `Cart`, `Order` and `Payment` without forcing every domain to import
  the others.

The important point is to make dependency direction explicit and avoid circular dependencies. Circular imports usually 
indicate that responsibilities are unclear or that orchestration logic belongs in a higher-level application service.

**Note:** This is not necessarily the optimal design for every system. The correct boundaries depend on your business
and product requirements, but the example is sufficient to illustrate the modeling process.

## Thinking in domains

To summarize, whenever you are modeling domains, keep these ideas close :

- Separate the logical capabilities of your application.
- Think about how domains interact with each other.
- Define a dependency graph that follows business ownership and explicit contracts.

In addition to these, borrow Evans' advice and **Model Out Loud**:

> "One of the best ways to refine a model is to explore with speech, trying out loud various constructs from possible
> model variations. Rough edges are easy to hear.
>
> - “If we give the Routing Service an origin, destination, and arrival time, it can look up the stops the cargo will
>   have to make and, well . . . stick them in the database.” **(vague and technical)**
> - “The origin, destination, and so on . . . it all feeds into the Routing Service, and we get back an Itinerary that
>   has everything we need in it.” **(more complete, but verbose)**
> - “A Routing Service finds an Itinerary that satisfies a Route Specification.” **(concise)**
>
> Domain-Driven Design, 2003, page 19

He argues that language is powerful, and phrasing a model in different ways helps you identify domain boundaries,
dependencies and responsibilities faster.

## Conclusion

Should you adopt a domain-driven approach to model your business application? I would argue that, in general, yes.
Thinking in domains improves your ability to understand a problem and design modular, flexible and maintainable logical
units with clear responsibilities.

Layer separation also helps keep concerns distinct. This makes the codebase more predictable and easier to evolve
without forcing every application into the same structure.

If domains are independent, adding new features and managing them can mean:

- Adding a new domain and introducing dependencies through a simple, explicit API.
- Testing modules independently and mocking them at their boundaries.
- Changing business logic inside one domain without spreading the same rule across the codebase.
- Extracting, deploying or scaling a domain independently later, when the operational benefits justify doing so.

If anything, remember that understanding the code is key to maintaining and evolving it.

However, if modeling an application in domains creates too much friction with a framework, consider simplifying your
assumptions and stripping away some layers. A framework may already handle infrastructure in a standard way, making a
strict separation unnecessarily costly. **Prefer the idiomatic way** of your language and framework, but **keep a
domain-based mental model of your application in mind**.

When a codebase starts growing faster than its architecture, unclear ownership and tangled dependencies can turn
ordinary changes into risky projects. The earlier you identify these problems, the more options you have for correcting
them incrementally.

For teams already struggling with these problems, Johntek Consulting provides architecture assessments that turn the
findings into a practical evolution plan.
