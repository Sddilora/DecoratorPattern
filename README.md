# Decorator Pattern
The Decorator design pattern is a structural pattern that allows you to dynamically add new behaviors or functionalities to an object without modifying its existing code. It follows the principle of open-closed design, which means that you can extend the functionality of an object without modifying its underlying implementation.

In the context of Go (Golang), the Decorator pattern can be implemented using a combination of interfaces and composition. Let's break down the implementation step by step:

1. Define the base interface: Start by defining an interface that represents the base functionality of the object you want to decorate. This interface should include the common methods that all decorators and the base object will implement.

2. Implement the base object: Create a struct that implements the base interface. This struct represents the core functionality of the object that you want to decorate.

3. Create the decorator interface: Define another interface that extends the base interface and includes additional methods specific to the decorators. This interface should also include a reference to the base interface.

4. Implement the decorators: Create structs that implement the decorator interface. Each decorator struct should have a field of the base interface type to hold a reference to the object being decorated. The decorator struct should also implement the methods defined in the decorator interface, adding new behavior or modifying the existing behavior as needed.

5. Decorate the base object: In your code, create an instance of the base object and then wrap it with one or more decorators. Each decorator wraps the previous decorator or the base object itself. This allows you to stack multiple decorators and dynamically add new functionalities to the object.

Here's a simplified example in Go to illustrate the Decorator pattern:

```go
// Step 1: Define the base interface
type Component interface {
    Operation() string
}

// Step 2: Implement the base object
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
    return "Base operation"
}

// Step 3: Create the decorator interface
type Decorator interface {
    Component
    AdditionalOperation() string
}

// Step 4: Implement the decorators
type ConcreteDecoratorA struct {
    component Component
}

func (d *ConcreteDecoratorA) Operation() string {
    return d.component.Operation() + " + Decorator A"
}

func (d *ConcreteDecoratorA) AdditionalOperation() string {
    return "Additional operation A"
}

type ConcreteDecoratorB struct {
    component Component
}

func (d *ConcreteDecoratorB) Operation() string {
    return d.component.Operation() + " + Decorator B"
}

func (d *ConcreteDecoratorB) AdditionalOperation() string {
    return "Additional operation B"
}

// Step 5: Decorate the base object
func main() {
    component := &ConcreteComponent{}
    decoratedComponent := &ConcreteDecoratorA{component: component}
    decoratedComponent = &ConcreteDecoratorB{component: decoratedComponent}

    result := decoratedComponent.Operation()
    additionalResult := decoratedComponent.AdditionalOperation()

    fmt.Println(result)            // Output: Base operation + Decorator A + Decorator B
    fmt.Println(additionalResult)  // Output: Additional operation B
}
```

In this example, we have a base interface `Component` that defines the `Operation()` method. The `ConcreteComponent` struct implements this interface and represents the base object.

We then define the `Decorator` interface, which extends the `Component` interface and adds the `AdditionalOperation()` method.

Next, we implement two decorators: `ConcreteDecoratorA` and `ConcreteDecoratorB`. Each decorator wraps the previous decorator or the base object and adds new behavior to the `Operation()` method. They also implement the `AdditionalOperation()` method with their own specific functionality.

Finally, in the `main()` function, we create an instance of the base object and wrap it with the decorators. We can then call the methods on the decorated object to see the combined behavior.

By using the Decorator pattern, we can easily add or remove decorators without modifying the base object or other decorators. This allows for flexible and dynamic behavior modification at runtime.The Decorator design pattern is a structural pattern that allows you to dynamically add new behaviors or functionalities to an object without modifying its existing code. It follows the principle of open-closed design, which means that you can extend the functionality of an object without modifying its underlying implementation.