# Nameco

Nameco is a tiny name generator. It could also be a word generator, if it wanted
to, but at the moment it has no such wish.

Nameco is also a pun on video game company, Namco. Nameco is not Namco; it is
merely a tiny name generator.

Nameco is based on Docker's
[namesgenerator](https://pkg.go.dev/github.com/moby/moby/pkg/namesgenerator).
Sorry, I should say that it's Moby's package. It's one and the same.

# Running

There is nothing special about Nameco's build procedures. Simply run `go build`
with any options you wish to build it; or you can run it directly with `go run`.
You should try to have fun with whatever choice you make.

When you execute Nameco, it will print out a name -- or possibly a word --
followed by a newline. You don't have any strict guarantee that the name will
have any particular form, other than it will be ASCII-compatible.

If you do not like the name that Nameco prints out, then you should run it
again, until you find a name that you like, or become bored.

# Testing

There are currently no tests for Nameco. Hopefully, it works as intended. Should
there be further code committed to this repository, tests will be added, to the
extent that they are desired.