# Core Engineer

Included is a console app which through a simple cold observable publishes a stream of "strings" through to the console.

1. Add another concurrent data source *Number* that periodically publishes a random Integer between 0 and 9.
2. Using the integers you receive at the subscriber, use that to fetch that numbered element from the feed of letters, I.e. If you receive a number 4, you print the 4th letter you received from Alphabet.
3. Suggest ways and/or refactor the example to handle the dispose of data sources in a better way.

There are several ways you can approach and make these changes, please include the following in your submission:
* Source code for your solution.
* Your reasoning behind your choice of solution.
* An indication of how long you spent on your particular solution.

### Stretch goals: (These are not compulsory)
1. Seperate the publisher and subscriber in to seperate processes.
2. Add in an adequate means of service discovery for the subscriber to find the provider.