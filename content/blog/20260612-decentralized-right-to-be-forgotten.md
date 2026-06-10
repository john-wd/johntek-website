---
date: '2026-06-12T13:24:10-06:00'
draft: false
featured: true
title: "Case Study: How to Implement the Right to Be Forgotten in Decentralized Systems"
description: >
  Privacy laws exist in many places and decentralized systems must respect them. 
  Learn how to implement auditable user deletion across decentralized services without breaking domain ownership
image: /images/posts/20260612/cover.jpg
categories:
  - data privacy
  - case study
  - architecture
---
Since the European General Data Protection Regulation (GDPR) began to take effect, many other countries 
have adopted similar privacy laws. One of the policies included in these kind of privacy laws is ["The 
Right to be Forgotten"](https://gdpr-info.eu/art-17-gdpr/), or put simply, the right that users have to 
request the deletion of their data from that service databases.

Even though there are slight variations between the legislations, the core principle remains the same: 
remove user data when requested in a given time frame. However, in some places you may need to retain
the user data for a certain period after the request, even if they no longer use the service. After that,
the data can be deleted or anonymized.

The challenge we present here is how we can effectively implement this right to be forgotten in decentralized 
systems, while maintaining engineering best practices and proper domain-level separation in those
architectures.

This article is a case-study of a similar architecture I proposed for one of my clients. The assumptions
and requirements list represents a real business scenario and the solution here is a practical implementation
for production.

## Nuances of this problem

At first, implementing the right to be forgotten may seem straightforward, just remove whatever user-profile
data and cascade deletion to related data across the database. However, it is much more nuanced than that.

Any data that can be combined to identify a user should be treated as personal data and handled carefully.
This is specially true in the health industry, where a set of health factors, journeys and records can pin 
down to a particular user.

Moreover, we are focusing on decentralized systems, so data is stored and managed by each individual
service, rather than a centralized authority.

Finally, there are requirements in these laws that enforces the systems to be thoroughly auditable and 
explainable in case of judicial inquiries.

## Assumptions

To properly set our problem here, we can assume a few things about our system, some redundant but laid out
explicitly for clarity.

- It is a complex SaaS business deployed in the cloud, with dozens of services.
- Your system is decentralized, many services run independently and communicate with each other.
- Service boundaries separated system domains logically (payments module, identity, cart, orders, etc).
- Services own the business data they are responsible for.
- User-related data may exist outside transactional services, including logs, analytics pipelines, 
  data warehouses and third-party processors, for example.
- The system is eventually consistent, so deletion is not necessarily instantaneous across all domains.


You can relax some of the assumptions here (say it is a monolith solution with proper domain separation) and 
this page can still be relevant.

If by chance your system is monolithic with no domain separation, you can [reach out to me](/contact) and we
can see how we can improve it and make it simpler, more modular and extensible so a solution like this page
can be applied.

## Requirements

With the assumptions in mind, we can define the solution requirements for this system:

- Each domain should handle the deletion of data they own.
- The system should be auditable, record when the request was initiated and when each domain handled it.
- There must be NO god service that knows about all domains.
- The user has a regret grace period to revert the request before data any data is deleted.
- Data must be retained for N days due to regulatory requirements. After this time, data is deleted.
- Deletion can be fulfilled through irreversible anonymization where legally and technically valid.
- The system must propagate deletion requests to relevant third-party processors.
- Monitoring and alerting must be in place to detect and respond to failures and issues on requests.

The last one is particularly important as it is quite often that system's act on behalf of the user, so data may 
exist outside your service. Think salesforce, shopify or whatever dependent services your system interacts with.

## Proposed architecture

We have decentralized microservices, each owning its specific domain data imposes quite some restrictions on system 
designs. We explicitly called out a god service that would have an inventory of all domains in the system,
so we cannot just have a service that connects to all databases to handle deletion requests. Even if your solution
has just one database with segratated tables by domain, it would still be a bad idea to have this. It goes against
the domain ownership model of your system.

Instead, we can have an event-driven architecture and fan-out deletion requests to all microservices that requires it.

{{<figure
  src="/images/posts/20260612/arch.jpg"
  width="100%"
  alt="A diagram of an event-driven architecture with fan-out deletion requests"
  caption="Event-driven architecture with fan-out deletion requests"
>}}

The diagram is involved, so let’s break it into parts.

### Completion tracking without a god service

In this model, the deletion service does not know how to delete data from every domain, but it still needs a way
to track which domains have acknowledged a request.

Domains subscribe to deletion events and emit a `DomainAcknowledged` event when they receive a request. 
From that point on, the deletion service tracks that domain as a participant for that specific request.

Completion is reached only when every acknowledged participant publishes a terminal status, either success or failure.

### Backups and restore safety

Backups naturally will include deleted data, but they must not. So whenever you are restoring from a backup,
replay the deletion events in the metadata store to restore the correct state, without the user data.

Retroactively deleting user data from backups usually is too costly and super difficult, often its management
cost is much higher than just replaying past events.

This implies that the deletion request store is not being rolled back too, so think of syncing its present state
on the restored backup so you can replay the events.

### User account deletion service

This is the core service that handles requests, stores metadata in the database, orchestrates lifecycle and manages
the deletion metadata.

In its database you want to store the request data (which user id is requesting, when the request was made, who made 
the request, etc). Here you also have an audit table that stores entries for created / updated / failed / finished
statuses.

Also add a retained table with basic user-profile data (email, name, date of birth...) encrypted at rest,
restricting access to everyone. A process should be defined to temporarily allow access to decrypting this data 
in case of judicial requests.

Here the following API is defined:
- **Entrypoint:** user-facing endpoint that initiates the deletion request. 
- **Status update subscription:** whenever domains finishes processing requests, messages are published to this
                              queue that are consumed by the deletion service.
- **Scheduler services:** to account for the grace period, this service monitors requests and initiates the process
                      when needed and removing the retained information from the database.

### Deletion protocols

Since this is event-driven, you need to define your protocol for sending and consuming events. This can be done
in a shared library that your services can use directly to remove boilerplate.

You can do something like this:

- Define some `Event` data struct that contains the `user_id`, `date_created`, and `status` fields.
- Define an `Forgeter` interface that accepts a custom function to handle the deletion process on a domain
- Define a custom Subscriber wrapper that subscribes a function to a Subscription and automatically handles 
  retries and error handling.
  - This wrapper publishes `StatusUpdate` messages to a queue to notify the orchestrator that the request
    was acked, successfully processed or failed.
- If retries exceed a threshold, push message to a DLQ (dead letter queue) to be retried later.

This way domains only need to import from this shared library and just implement the deletion logic important
to them, the rest is handled by the shared library.

The entrypoint service emits events using the same structs of this library.

### Audit log system

This should be a read-only database, so records are stored in append-only fashion. Whenever a new audit
log is created, it records the timestamp, userId, requestId, domain and any other relevant information
of that log.

### Domain logic implementation 

Now in this architecture, each domain needs to implement its own logic. Following the proposal above, we have 
a shared module that exposes easy interfaces for us, so we can just deal with the domain logic.

Something like this, supposing we have a Deleter interface that exposes a delete() method receiving the 
user id as parameter.

```go
func (d *DomainDeleter) Delete(ctx context.Context, userId string) error {
    // assume we have a domain storage in d
    ents, err := d.domainStorage.FetchFromUser(userId)
    if err != nil { return err }
 
    for _, e := ents {
         err = d.domainStorage.Anonymize(e)
         if err != nil { return err }
    }

    // call upstream services too
    err = d.upstream.DeleteUser(userId)
    return err
}
```

Each domain will have its own design logic and dependencies, so this is just an example to illustrate 
how you can use it.  The core component reports errors, pushes failed messages to the DLQ, and publishes 
status updates based on the returned error.

### Scheduler lifecycle control

Since user deletion is not something that needs to be processed in real time, as time units are 
measured in days, we can run a scheduler daily to gather all records that are due to be fully purged 
after the decided retention period. This will remove the user’s data from the retention quarantine and 
finally their data will be fully removed from the system.

Scheduler is also used to check for status on deletion execution for each domain. Once they receive an 
ack, they save a status of processing to that domain. After all those statuses are marked success, 
the system may call a notification service to tell the user their data has been removed.

Notice that in this model the orchestrator does not know about all domains, rather the domains send a 
message saying “I exist” before it can tell if requests were fulfilled or not.

## Handing failure

In a complex system like these, there can be a handful of ways the system can fail. You will need to 
keep an eye on these in order to comply with the local laws of your country.

### Failed messages and monitoring

In general you want to have dlqs in all subscriptions queues, so your messages are not lost on failure.

In addition, consider adding alerts to those dlqs as well, so if there is at least one message 
available, your on-call engineers can take a look at the domain handler that is failing and take
action. Sometimes business rules change and the deletion handler drugs silently, other times 
upstream services change. Be aware of these.

### Requests not being fulfilled in time

Consider also having a checker scheduled job to check the state of requests. You'll have to 
fulfill the deletion in a given number of days and be subjected to fines if not fulfilled. 
If it is getting closer to this deadline, consider alerting too so you can keep an eye.

### Retry mechanism

Consider giving a retry mechanism to stuck domain requests. Handlers should be idempotent, 
so you should be able to retry a message for a specific domain if it is stuck for some reason 
(app crash, deployment stopped). These management endpoints should belong to the deletion domain.

### Authorization considerations

Deleting user data is sensitive and not many people should be able to issue requests, not even 
internal personnel. Make strict authorization to these endpoints and only allow a restricted 
number of people to be able to hijack/issue requests on behalf of others.

Ideally,
- only the requesting user should be able to delete their own data, i.e. call the entry point 
  endpoint
- No one should be able to write to deletion queues or topics
- Ops should be able to use the management endpoints to retry existing requests, 
  not fabricate new requests
- No one is allowed to query the retention table, only temporary permissions on approval 
  process for extraordinary judicial cases
- Developers can see logs on failed domain handlers they own, so they can troubleshoot issues, 
  and be able to retry failed requests

## Final considerations

This is a complex topic with a complex architecture that orchestrates user deletion data across 
isolated services. The solution given here is actually active and serving requests with little issues
for over 3 years and it is modular and extensible to new services my client wants to create. 
They just need to provision the required cloud resources and write the deletion logic to that 
domain and the system automatically handles that use case too without “knowing” about that domain 
beforehand.

Monitoring and alerting is invaluable and, I would say, indispensable here, catching little issues 
with API changes early on, never getting close to the legal deadline.

I believe that this use case can help you design something as robust as this for your business reality.

If your user deletion process still depends on manual tickets, direct database scripts, or a central service 
that knows too much about every domain, this is a sign that the architecture may not scale with your 
compliance requirements.

A robust deletion workflow should be auditable, domain-owned, idempotent, observable and safe to retry. 
It should also make failures visible before they become legal or operational problems.

If you are designing this kind of workflow, or if you are unsure whether your current architecture can 
handle deletion requests reliably, [I can help you](/contact) assess the risks and design a practical 
implementation for your system.