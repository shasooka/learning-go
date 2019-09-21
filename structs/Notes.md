### What I learned

* Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
* Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
* Adding methods so you can add functionality to your data types and so you can implement interfaces
* Table based tests to make your assertions clearer and your suites easier to extend & maintain

Interfaces are a great tool for hiding complexity away from other parts of the system. In our case our test helper code did not need to know the exact shape it was asserting on, only how to "ask" for it's area.
