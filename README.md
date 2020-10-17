# Golang Design Patterns
**Status: WIP**

List of Golang Design Patterns based on Gang of Four (GoF).

## Principles & Building Blocks

Author understands that the implementation of design patterns in Golang can vary from place to place. In these examples,
author is using these principles & building blocks:
* Following SOLID principles
* Maximizing usage of interface for abstractions instead of standard functions for easier mocks & testability
* Ensure a concrete folder structure to make it resemble real world usage

## Contexts

Author is using a real-world example in order to ease the reader in understanding the real-world usage of the design
patterns. The real-world context used here is e-commerce industry.

## Design Patterns List

### Structural Patterns

| Pattern | Description | Example's Context |
|:-------:|:----------- |:----------------- |
| [Decorator](/decorator/README.md) | Wrapping new behaviors on another behavior | Database and caching |

### Behavioral Patterns

| Pattern | Description | Example's Context |
|:-------:|:-----------|:------------------ |
| [Chain of Responsibility](/chainofresponsiblity/README.md) | Passing request along a chain of handlers | Checkout process |
| [Strategy](/strategy/README.md) | Enables a family of algorithms to be selected at runtime and interchangeable | Get shipping rates from multiple shipping partners | 
| [Template Method](/templatemethod/README.md) | Defines a skeleton class which defers some methods to subclasses | Send sales report for multiple recipients (eg. Seller, Investor) |

## References
https://refactoring.guru/design-patterns
https://learning.oreilly.com/library/view/go-design-patterns/9781788390552/
https://golangbyexample.com/all-design-patterns-golang/
https://github.com/tmrts/go-patterns