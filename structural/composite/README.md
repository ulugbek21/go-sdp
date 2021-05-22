# Composite Pattern

The Composite design pattern favors composition (commonly defined as a has a relationship) over inheritance (an is a relationship). The composition over inheritance approach has been a source of discussions among engineers since the nineties. We will learn how to create object structures by using a has a approach. All in all, Go doesn't have inheritance because it doesn't need it!

In the Composite design pattern, you will create hierarchies and trees of objects. Objects have different objects with their own fields and methods inside them. This approach is very powerful and solves many problems of inheritance and multiple inheritances. (Contreras, 2017)

## Implementation

For example, a typical inheritance problem is when you have an entity that inherits from two completely different classes, which have absolutely no relationship between them. Imagine an athlete who trains, and who is a swimmer who swims:

- The _Athlete_ class has a _Train()_ method
- The _Swimmer_ class has a _Swim()_ method

The _Swimmer_ class inherits from the _Athlete_ class, so it inherits its _Train_ method and declares its own _Swim_ method. You could also have a cyclist who is also an athlete, and declares a _Ride_ method.

But now imagine an animal that eats, like a dog that also barks:

- The _Cyclist_ class has a _Ride()_ method
- The _Animal_ class has _Eat()_, _Dog()_, and _Bark()_ methods

Nothing fancy. You could also have a fish that is an animal, and yes, swims! So, how do you solve it? A fish cannot be a swimmer that also trains. Fish don't train (as far as I know!). You could make a Swimmer interface with a Swim method, and make the swimmer athlete and fish implement it. This would be the best approach, but you still would have to implement swim method twice, so code reusability would be affected. What about a triathlete? They are athletes who swim, run, and ride. With multiple inheritances, you could have a sort of solution, but that will become complex and not maintainable very soon. (Contreras, 2017)

## Summary

Use the Composite Design Pattern to avoid typical problems with inheritance.
