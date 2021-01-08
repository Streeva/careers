# Core Engineer

Complete *one* of the following exercises.

## Exercise 1: Encoding test

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

## Exercise 2: Service / Docker test

For this exercise you can use a language of your choice, although C# or Golang is preferred.

1. Implement a service which provides an HTTP endpoint to Base64 Decode supplied Base64 encoded strings.
2. Write a Dockerfile which will generate a ready-to-run image for your service.
3. Provide instructions on how to build your image.

### Stretch goals (not compulsory)
1. Document the API to your service.
2. To build your project only Docker / Docker-compose is required.

## Submissions

Please include the following in your submission:

* Source code for your solution.
* Your reasoning behind your choice of solution.
* An indication of how long you spent on your particular solution.
