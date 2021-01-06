# Core Engineer

Complete *one* of the following exercises.

## Exercise 1: System.Reactive test

Included is a console app which through a simple cold observable publishes a stream of "strings" through to the console.

1. Add another concurrent data source *Number* that periodically publishes a random Integer between 0 and 9.
2. If the latest integer you received is *n* and the latest letter you received is *x*, write out the letter at *x-n*.  E.g. If you have received the sequence *A B C D E F* and the number *2*, you write out *D*.
3. Suggest ways and/or refactor the example to handle the dispose of data sources in a better way.

### Stretch goals: (These are not compulsory)
1. Seperate the publisher and subscriber in to seperate processes.
2. Add in an adequate means of service discovery for the subscriber to find the provider.

## Exercise 2: Azure functions test

1. Replicate the functionality of the reactive test as two or more Azure functions conncted via ServiceBus.
2. Provide an ARM deploy template for your functions project.

## Exercise 3: Encoding test

Given the following BER-TLV encoded string:
```
0812561ae6bd42cfb7bee1311a29893508e71c340912e4614f3d9e6c61dccdd5af228051eb4373170A058394b17ac301086ef72ac5367180eb051040d4241e426bbde12bca805f665227e8060234ec07042859ffbc0202b4ed0304bdadb4e00402de60
```
1. Extract and order the tag values for the following tag ids by ascending numeric tag id.
```
0x01, 0x03, 0x04, 0x05, 0x06, 0x0a
```
2. Create an uppercase string with a full stop between each tag value.
3. Convert the string to a byte array.
4. Hash the byte array using the Blake2b algorithm.
5. Base64url encode the resulting hash.
Starting from the above example you should get the following output:
```
EQ9q1Hjyr_pfSBj1_M2vqAMFH-MYhP2zwhgYfC8u8-g
```

## Exercise 4: Service / Docker test

For this exercise you can used either C# / Golang or a language of choice.

1. Implement a service which provides an HTTP endpoint to Base64 Decode supplied Base64 encoded strings.
2. Write a Dockerfile which will generate a ready-to-run image for your service.
3. Provide instructions to build your image.

### Stretch goals (not compulsory)
1. Document the API to your service.
2. Only docker / docker-compose are required as pre-requisites on the build machine.

## Submissions

Please include the following in your submission:

* Source code for your solution.
* Your reasoning behind your choice of solution.
* An indication of how long you spent on your particular solution.
