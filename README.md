# Design Patterns in Go

## Creational Patterns
- **Singleton**: Ensures a class has only one instance and provides a global point of access to it.
- **Builder**: Separates the construction of a complex object from its representation.
- **Factory Method**: Defines an interface for creating an object, but lets subclasses alter the type of objects that will be created.
- **Abstract Factory**: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
- **Prototype**: Creates new objects by copying an existing object, known as the prototype.

## Structural Patterns
- **Adapter**: Allows incompatible interfaces to work together.
- **Bridge**: Separates an object’s abstraction from its implementation so that the two can vary independently.
- **Composite**: Composes objects into tree structures to represent part-whole hierarchies.
- **Decorator**: Adds new functionality to an object without altering its structure.
- **Facade**: Provides a simplified interface to a complex subsystem.
- **Flyweight**: Reduces the cost of creating and manipulating a large number of similar objects.
- **Proxy**: Provides a placeholder for another object to control access to it.

## Behavioral Patterns
- **Chain of Responsibility**: Passes a request along a chain of handlers.
- **Command**: Encapsulates a request as an object, thereby allowing for parameterization of clients with queues, requests, and operations.
- **Interpreter**: Implements a specialized language.
- **Iterator**: Provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
- **Mediator**: Defines an object that encapsulates how a set of objects interact.
- **Memento**: Captures and externalizes an object's internal state so that the object can be restored to this state later.
- **Observer**: Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
- **State**: Allows an object to alter its behavior when its internal state changes.
- **Strategy**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable.
- **Template Method**: Defines the skeleton of an algorithm in an operation, deferring some steps to subclasses.
- **Visitor**: Represents an operation to be performed on the elements of an object structure.