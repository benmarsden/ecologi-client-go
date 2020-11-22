# How to contribute to ecologi-client-go

Thanks for considering a contribution! Whether it's a bug fix, feature idea, pull request review or other, they are welcome ✨

## Pull Requests

[Fork](https://help.github.com/articles/fork-a-repo/) this repository to your own GitHub account, make any desired changes and then submit a pull request to the benmarsden/ecologi-client-go repository.

## Bugs and Feature Requests

Spotted an issue, know of an endpoint that is not yet supported, or want some other feature? Raise it as issue in this repository. Please include as much detail as possible when doing so.

## Testing and QA

### Makefile

For testing convenience, `Makefile` provides some nice testing targets you can use. For example:

- `make test` runs all repository tests, and is a quick and easy way to verify you have not broken anything. 
- `make display-coverage` runs all repo tests and then loads an interactive coverage viewer.  
### pre-commit

[pre-commit](https://pre-commit.com/) has been set up in the repository. Use this to keep your commits clean.

### CI 

CI runs (that test and build the client library) take place on every commit and pull request to the repository. This must pass before any PR is considered valid for merge.

To test the outcome of CI locally, you can use [act](https://github.com/nektos/act). The Makefile target `make act` has been set up to support this.

