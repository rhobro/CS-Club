# Computer Science Club

I run a computer science club at school for younger students.

I have found that CS syllabi usually restrict themselves to problems which can be solved by a single machine.
That is why I am striving to teach members of the club about HTTP and other protocols which allow computers to
interact with eachother over the Internet, even if they are not in geographical proximity.

Having taught the fundementals of requests, I think my students will need more practice using them in order to
gain confidence and fluency with HTTP practices. At the same time, I would like to develop their problem solving
skills by teaching them about algorithms and their uses.

It is with these two goals in mind that I have created a web server in Go which acts as a system by which algorithmic
challenges can be hosted at certain endpoints, delivering problems and receiving solutions using simple requests. It
will be possible for me to add new endpoints as I develop new problems for particular topics.

It is my hope that my successor will be able to create challenges in this way.

### Structure
- `main.go` contains the basic code to start the server on localhost port 8080.
- `db/` implements an in-memory database to store active problems in a map system. In order to prevent data race conflicts where possible,
mutexes have been used to ensure concurrent goroutines do not cause errors.
- `handler/` contains the paths and implements HTTP handlers for the challenges. It utilises a subrouter system to allow the challenges to
be split into topics. Generally, problem generating endpoints will be of the format "give..." and receiving endpoints will be named "receive..." . The general convention for this project is to use the same HTTP endpoint for the same problem but to use GET and POST requests to
differentiate between problem generation and submission.
- `challenge/` contains some general code for challenges. `notify.go` is code that will be run on every submission - just for visual feedback.