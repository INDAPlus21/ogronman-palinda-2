## Answers

* What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?
  
  The thing that happens when you switch the order of the statements `wgp.Wait()` and `close(ch)` is that the program will most likely crash. The reason why the `wait()` comes before the `close()` is that the program waits for all the running routines to finish before closing the channel. However if you switch the statements there is a risk that a function will send a value to the channel (`ch`) after it has been closed, since we are still waiting for all functions to finish. When you send a value to a closed channel, then the program chrashehesses.

* What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?

    Since the program is starting and running multiple `Produce` functions at the same time, moving the `close(ch)` to the end of the `Produce` function will after a while crash the program. This is because one of the many `Produce` functions will have finished their for loop and will then close the channel. However the others that have not yet finished will then try to write to a closed channel, which will crash the program.

* What happens if you remove the statement `close(ch)` completely?

    Nothing will happen of you remove the `close(ch)` completely since `golang` has a built in garbage collector that collects and closes all channels that are not in use

* What happens if you increase the number of consumers from 2 to 4?

    The program finishes faster since the consumers can consume collectively faster :)

* Can you be sure that all strings are printed before the program
  stops?

  No we can not be sure of this, since there is no wait group that waits for all consumers to finish, the program only runs as long as the `Producer`s produce something, however the consumers does not consume as soon as something is produced

Finally, modify the code by adding a new WaitGroup that waits for
all consumers to finish.

