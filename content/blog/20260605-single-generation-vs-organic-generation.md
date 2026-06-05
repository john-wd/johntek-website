---
date: '2026-06-19T12:16:44-06:00'
draft: false
title: 'AI Can Generate Code, but It Struggles to Evolve It'
description: 'AI tooling can generate text, image and assets just fine, it looks great, but it struggles to generate code that can evolve.'
image: null
categories: 
  - ai
---

Since ChatGPT was first released, we have seen AI tools evolving from generating plausible text to generating 
various types of content: prose, poems, images, music, videos and now, as it is so hyped, code.

The first iterations of image generations were quite bad by today's standards but it carried a promise of
better and better fidelity as models improved. Similarly, the same was true for all other content types.

{{<figure
  src="/images/posts/20260605/will-spaghetti.webp"
  alt="Will Smith Spaghetti"
  caption="Evolution of AI-generated images, illustrated by the Will Smith Spaghetti meme"
  width="100%"
>}}

Today we have models trained specifically to target the technology sector and its job market, releasing 
tooling that aids Engineers and Software Developers in general to help with their code work, saving them
from typing code manually.

## A brief history on coding AI tooling

First Github's Copilot introduced little auto-complete features that would guess what you would type next,
then integrations for IDEs that let developers make little prompts to write snippets of code. These were often
wrong at the time, but it got to a point where autocomplete is quite good and actually predicts what we want
to type next.

Enter early 2026, "coding agents" started handling larger parts of the coding workflow. These agents
were equipped with system-level tools that let them grep and search for patterns to build a context and write
something that closely resembles what you have in your codebase.

This generated code mimics the codebase quite well, but it often is too verbose and aggressively defensive, 
accruing to a lot of unnecessary Lines of Code (LOCs) that just make the system more complex 
than it should. Moreover, often they choose creative "hack-y" ways to solve problems that can hide bugs or, 
worse, hinder future iterations. This [defines what I call The Hackathon Mentality, previously 
discussed in this blog page](/blog/20260601-the-hackathon-mentality/).

## Types of Generation

Let us first make the distinction of the types of generation we can have in our context. It is sufficient
to define just two: Single Use Generation and Iterative Generation.

### Single Use Generation

When I refer to this kind of generation, I am not referring to one-shot prompts, but rather any kind of content
generation that once done, is not retouched anymore; it is frozen in time in that state.

Use cases that fall in this category include:
- Research
- Images
- Assets
- Music
- Text (content, summaries, etc.)
- Reports
- Videos
- Static landing pages

and so on. The person can prompt many times to refine this and that detail, but once finished, they call it a day
and use that content as they seem fit. 

We can also include static landing pages into this basket because often they can be a one-off thing that won't
be iterated over.

{{<figure
  src="/images/posts/20260605/mascot.jpg"
  width="80%"
  alt="Johntek Mascot"
  caption="A very cringy mascot for this website"
>}}

For this kind of content, AI agents do a very good job and it really has a profound impact. It is faster to research
something with an agent, browse many websites and make a summary of them than to do it yourself. It also has opened
an new avenue for small business to make assets for pages, websites and information material which they wouldn't
before it.

Of course there is the argument that the quality is not all there yet, but it is good enough that many 
businesses are taking right now.

### Iterative Generation


Iterative generation on the other hand means any kind of content generation that is done over time, each generation
is the input to the next one, with the system learning things from the previous state and from next prompts.

Any kind of evolutive work fits in this category, but for our purposes we will focus on **code generation** only,
as it is my specialty. Thus, we have the following characteristics:
- Previous generation is context for the next iteration
- Hacks, shortcuts and mistakes accumulate over time
- Hacks, shortcuts and mistakes become references in the future
- Context becomes harder to preserve
- Patterns drift over time

In contrast with single generation, these points raise red flags that we need to be aware of when writing
serious production code that thousands if not millions of users depend on. Let us look carefully at some of these

{{< figure
  src="/images/posts/20260605/iterative.jpg"
  width="100%"
  alt="Iterative generation"
  caption="An example of iterative generation"
>}}

### Future generation depends on previous generation

[We know that agents are susceptible to being verbose and smart to take shortcuts](/blog/20260601-the-hackathon-mentality/),
so in the context of iterative generation, these clever hacks will serve as reference to future iterations. So
in a way, without diligent review, these patterns tend to repeat to new use cases and start to make system "radioactive"
with each iteration.

These anti-patterns are not healthy and may cause on-call engineers to wake more often at 3 am to troubleshoot hidden
issues that are blowing up production, or may hinder the ability of the product to evolve easier over time. This cost
is a debt on developer experience that will be paid later on.

### Context becomes harder to preserve

If generated code is too verbose, defensive and frequently single-use only, then the next context window will be 
larger to accommodate it all. As we know, [context tends to degrade with size](https://www.producttalk.org/context-rot/),
so middle part of it will silently be forgotten favoring the beginning and end of the context window.

The consequences are basically agents forgetting their task at hand and hallucinating new patterns instead of the ones
defined by you. This further degrades the quality of your code and the issue compounds over time.

### Patterns drift over time

Agents are now creating their own patterns and free floating functions to deliver their prompted tasks. Great, now
this will be reinforced as a system pattern in next iterations, further increasing the volume of these and increasing
the context window next. As we saw above, context degradation is amplified and more new patterns emerge. 

Give a couple dozen of evolutions and your defined patterns will be forgotten and replaced by new or a frankensteinian
monster child of them both.

## When this becomes a business problem

{{<figure
  src="/images/posts/20260605/melting.webp"
  width="100%"
  alt="Melting IDE"
  caption="Codebase decay given multiple generations"
>}}

This becomes risky when AI-generated code starts entering production without architectural review, clear ownership 
or stable engineering standards.

Warning signs include:
- PRs are getting larger but not easier to review
- Similar features are implemented in different ways
- Error handling and observability are inconsistent
- Developers are spending more time correcting generated code than designing the system
- Incidents are harder to debug because the code is harder to reason about

## Conclusion

LLMs are quite capable and they give the illusion of aptitude, when in reality they are just ellaborate mimicry machines.

For single-use generation it can do alright and produce quite convincing assets that can be used in presentations,
landing pages and styling options. They can help with report summaries and even design look-and-feel. Those are great
assets that can be readily used.

For iterative generation, however, this is not necessarily true. Products and production environments are often complex
systems that require a high degree of stability and predictability. Replacing developers with these agents will
inevitably lead to unmaintainable chunky code that is hard to understand and debug, let alone troubleshoot during
incidents. Without **proper human oversight**, the quality of the codebase will degrade over time.

To mitigate this, while using coding agents, we suggest 
- thorough code reviews and simplification as much as possible
- enforcing adherence to coding standards and best practices
- monitoring and alerting to detect and respond to incidents quickly

I believe the strongest use for AI in developer workflows involves these steps:
- Use AI for research, prompt it to give alternative solutions to your pre-conceived ideas
- Use codeagents for POCs, draft something together, see how it behaves and if it meets your acceptance criteria
- Generate documentation once you are happy with your code, to have a stable API you are happy with
- Ask it to apply some small pattern in your components (like monitoring or replacing error types)
- Put it in PR reviews to make a summary of changes, so reviewers can quickly understand what is intended before
  reading the first line of code

In essence, don't delegate your trust to AI yet.

If AI-generated code is already entering your production codebase, Johntek can help you review the workflow, 
identify risk areas, and define guardrails before technical debt compounds. [Get in touch](/contact).