# Architecture Patterns with Go

## Terms

Abstraction
: A way to encapsulate behavior by identifying a task that needs to be done and then assigning it a well-defined object or function. The object or function is the abstraction.

Domain
: The problem that you are trying to solve.

Domain model
: Mental map of the business.

Encapsulation
: Simplifying behavior and hiding data. Encapsulated behavior is an abstraction.

Model
: A map of a process or phenomenon that captures a useful property.

Ubiquitous language
: Business jargon that is shared and understood by everyone involved in the project, such as product management and engineering.

## Three-layer architecture



## Dependency Inversion Principle

The 'D' in SOLID:
- **S**ingle responsibility
- **O**pen for extension but closed for modification
- **L**iskov substituion
- **I**nterface segregation
- _**D**ependency inversion_

Dependency inversion requires that high-level modules do not depend on--or _know the details_ about--low-level modules. Both high- and low-level modules should depend on abstractions. For example, a driver does not need to know the inner-workings of a car's steering system--the driver only needs to know how to operate the steering wheel.

## Domain modeling

Business goals should drive the approach to software design. Behavior should come first and drive the storage requirements.

The terms _business logic_ and _domain model_ are interchangeable. What they really mean is, "What problem am I trying to solve?". Each 



## Patterns

### Repository
A repository is an abstraction over persistent storage.

### Service layer
The service layer pattern clearly defines where our use cases begin and end.

### Unit of work
The unit of work provides atomic operations, 

### Aggregate
This pattern enforces data integrity.