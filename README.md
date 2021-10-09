
# Mid-term test 1 (variant 1)

The test consists of three tasks located in directories `app`, `physics` and
`ping-pong`. Fork the repository, add/edit files in the directories in order to
complete the test. The completed test must be submitted as a GitHub pull request
to `prog-1/test-1-1` repository.

1. Explain what the program in the directory `app` does. Add your explanations
as comments in the source code. Feel free to add any comments that you think are necessary.

2. Write a program in the directory `physics`. The user provides three numbers:
capacity of three condensers `C1` `C2` and `C3`. The program must calculate and
print the total capacity of the condensers connected in series. The total capacity `C` can be found using the formula `1/C = 1/C1 + 1/C2 + 1/C3`.

   Example:

   ```
   The program finds the capacity of three condensers connected in series.
   Enter capacity of three condensers: 1 2 3
   The total capacity of the three condensers connected in series is 0.5454545454545454.
   ```

3. Write a program in the directory `ping-pong` that prints the numbers from 1
to 100 (each number in a separate line) and for multiples of `2` prints `Ping`
instead of the number and for numbers that are NOT multiples of `5` prints `Pong`.

   Example:

   ```
   Pong
   PingPong
   Pong
   PingPong
   5
   PingPong
   Pong
   PingPong
   Pong
   Ping
   ...
   ```