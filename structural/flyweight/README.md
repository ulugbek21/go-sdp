# Flyweight Pattern

Flyweight is a pattern which allows sharing the state of a heavy object between many
instances of some type. Imagine that you have to create and store too many objects of some heavy type that are fundamentally equal. You'll run out of memory pretty quickly. This problem can be easily solved with the Flyweight pattern, with additional help of the Factory pattern. The factory is usually in charge of encapsulating object creation, as we saw previously (Contreras, 2017)

## Implementation

Imagine Championship Final League game that is watched by millions of people. You have a website of live game's textual translation and, additionally, your website has a history of football clubs playing in the final. This is plenty of information, which is usually stored in some distributed database, and each team has, literally, megabytes of information about their players, matches, championships, and so on.

If a million users access information about a team and a new instance of the information is created for each user querying for historical data, we will run out of memory in the blink of an eye. We will store each team's information just once, and we will deliver references to them to the users. So, if we face a million users trying to access information about a match, we will actually just have two teams in memory with a million pointers to the same memory direction.

Here we have a map of teams, that stores pointer to specific team:

```go
type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}
```

This way we will not allocate the information in the memory about the team each time, when users request historical information about the football club

## Final thoughts

Flyweight is very commonly used in computer graphics and the video game industry, but not so much in enterprise applications. (Contreras, 2017)
