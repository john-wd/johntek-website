---
date: '2026-07-03T10:28:38-06:00'
draft: false
featured: true
title: "How do I use AI to boost my productivity? A Senior Engineer's Perspective"
description: >
  AI tools are here claiming to boost productivity by 10x, but how do Johntek Consulting leverages it?
  Check what strategies I use to enhane my work productivity.
image: /images/posts/20260703/cover.jpg
categories: 
  - ai
  - johntek
---
It is no secret that I have my skepticism with AI **coding** agents, as you can see from my previous posts in
this blog. However, like pretty much everyone nowadays, I do use AI tools extensively, both in personal life
as well as in my professional work.

Though I do not use or leverage AI agents quite as much because I believe there is still a long way for them to
become good enough to generate code capable of evolving healthly and safely. So how do I use these tools
professionally?

In this blog I want to share my experience and strategies for effectively using AI **as a tool** (not as my
replacement) to one-up my decision process.

## Initial thoughts

Before diving into the strategies, let's first establish some initial thoughts about my view on AI tooling.

### I do not transfer trust to AI

The final word on some new pattern or solution should be done by an experienced human who knows the surrounding
business context and constraints.

There are effectively endless ways something can be implemented, but picking the correct approach requires deep
understanding of the problem domain and available resources. 

### Respect for code styling and patterns

Patterns and code styling defines the way your product is logically structured, to alleviate the thinking burden
when evolving it. A well-designed codebase is built on solid principles that separates domains and concerns,
is modular and ready for future changes.

AI agents does not understand the nuances and best practices of your solution, and often they hallucinate new 
patterns that does not fit your codebase. This introduces drift that 
[can be difficult to reason about and serve as reference in future iterations](/blog/20260619-single-generation-vs-organic-generation).

Even if you sat and wrote many skills for each pattern, business context and detailed your task with enough 
granularity, the problem of context rot may still kick up and new unwanted code may be introduced. 

## Strategies

### High-Level Brainstorming and Risk Identification

One of the best strategies I like is using AI to brainstorm ideas. Here I usually have a good understanding of the problem space
and requirements, and I have a very high level working architecture in mind. So I outline the components, system patterns and 
implementation details I'm considering, then I ask the AI for direct feedback.

This initial step let me see suggestions for enhancements, like: better reliability, better scalability, enhanced modularity
and, perhaps most important, better security. AI is great at formulating suggestions based on real ideas instead of a generic
outcome-oriented prompt.

In follow-up conversations, I push to stress-test the system for production reliability, what can go wrong and how big it may
blow up.

Here I am directing the design of the system and the AI gives me a broader perspective, bridging edge cases to my attention
before they become baked into the infrastructure.

### Technical Competitive Analysis

After the initial design feedback, I then like to run a competitive analysis against my own choice, not for robustness, but rather
to see if there are better architectures out there that I could have chosen given my constraints. After all, our job as engineers
is picking up the "better" solution considering the requirements, since rarely there is a single "right" solution. It always
is an art of finding the right balance.

I give the AI my proposed design and ask for three to five alternative architectural suggestions. Then I refine it better by asking
to optimize for some specific characteristic, such as high reliability or low latency. Finally I ask it to format into a comparison
matrix table so I can easily see the pros and cons of each alternative (and use it in my decision documentation).

Comparing these AI-generated alternatives against my original plan ensures that my final technical choices are intentional, 
justified and thoroughly vetted.

### Rapid Proof of Concepts (POCs)

Sometimes you need to test if an idea is viable before pitching it to leadership or prioritizing it on a product roadmap. 
This is where code generation shines, provided you treat the output correctly. I will feed a segment of the codebase into 
the AI and ask it to draft a rough implementation of a pattern or feature just to see the code working.

The crucial rule here is to treat this generated code purely as temporary, exploratory scaffolding. You should never copy-paste 
an AI POC directly into production. Relying 100% on AI-generated production code introduces serious long-term maintenance issues.
Use the POC to prove viability, then throw it away and use your vetted architecture patterns to build the real feature sustainably.

You can quickly iterate your POC to see if it matches your desirable outcomes, then use the POC as a demo to convince stakeholders
to prioritize this feature.

### Bootstrapping System Documentation

A well-architected solution will always evolve slightly during implementation, so the code in production is rarely a 100% replica 
of the original technical design paper. Once a feature is deployed and working smoothly, I like to use AI to kickstart the 
documentation process.

There is some preliminary documentation on the decision making process, things like problem and solution definitions, architecture
decisions and other related context, but once the feature is finished, it is always good to have proper documentation, things like:

- Architecture deep dive, including component interactions and dependency graphs.
- Developer guide, entrypoints and domain boundary APIs
- Deep internal documentation on domain components, models, dataflows etc
- High level architecture overview, how the feature fits into the overall system.

So what I do is pointing the AI towards bits of the codebase at a time, so it can help me map out all these. 
This output serves as a baseline that I refine into thorough, detailing more in technical terms the text and pruning where
AI is too verbose or repetitive.

Remember that a well documented system makes enhances developer experience by providing clear, up-to-date documentation, but it
also makes AI tooling better too, by providing useful, rich and accurate context about the codebase.

### Targeted Unit Testing and the 2-Example Rule

Writing unit tests is easily the most impactful daily use case for AI, but letting an agent loose on an empty test file is
a mistake. Letting it generate tests alone tends to write bloated tests that check meaningless things, like verifying a basic 
constructor returns an instance or testing if a constant is defined correctly, delaying precious CI time.

To keep tests lean and meaningful, I first write one or two cases manually, testing two branches of a function, then delegate
the rest to the AI. Because it has an explicit template to  follow, the AI is remarkably accurate at filling in the blanks. 
Finally, review the generated suite, strip away any extra fluff, and commit a tight, meaningful, high-performing test suite.

Giving examples unlocks the AI's ability to generate tests like you would normally, which optimizes your time very much.

## The Bottom Line

AI is an excellent daily driver for brainstorming, exploring, documenting and generating tests. But if you treat it as an 
autonomous developer, it will inevitably leave you with fragmented, unmaintainable systems. Codebase decay compounds quickly, 
shifting the burden onto on-call engineers and stalling your product roadmap.

This article shows good use cases that enhances senior engineers' ability to write high-quality, maintainable tests, while
also expanding their architecture vision, shading light into unconsidered corners of the their designs. If used effectively,
seniors can really get more done faster without delegating all the trust to the AI.

Now, if AI-generated code snippets are silently creeping into your production environment without architectural guardrails, 
it's time to clean up the technical debt before it becomes a business crisis.  [Reach out to Johntek Consulting](/about) 
today to set up clear engineering standards, establish workflow guardrails, and keep your systems scalable for the long haul.
