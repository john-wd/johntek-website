---
date: '2026-06-12T13:24:10-06:00'
draft: false
featured: true
title: "Case Study: The Right to be Forgotten in Decentralized Systems"
description: >
  Privacy laws exist in many places and decentralized systems must respect them. 
  Learn more how to implement this feature in your decentralized application.
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

## Nuances of this problem

At first, implementing the right to be forgotten may seem straightforward, just remove whatever user-profile
data and cascade deletion to related data across the database. However, it is much more nuanced than that.

Any data that can be aggregated in such a manner to identify an user can be categorized as sensitive user
data and must be handled with care. This is specially true in the health industry, where a set of health
factors, journeys and records can pin down to a particular user.

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