# Contributing

## Guidelines for pull requests

- Write tests for any changes.
- Separate unrelated changes into multiple pull requests.
- For bigger changes, make sure you start a discussion first by creating an issue and explaining the intended change.
- Use [conventional changelog conventions](https://github.com/bcoe/conventional-changelog-standard/blob/master/convention.md) in your commit messages.

## Development dependencies

- go 1.11 or higher

## Setting up a development machine

Install all dependencies, lint code, run all tests
```
make
```

## During development

Commits to this codebase should follow the [conventional changelog conventions](https://github.com/bcoe/conventional-changelog-standard/blob/master/convention.md).

- `make verify` - Runs all the tests and lints commit messages. Execute it before pushing any changes.
- `make watch` - Runs tests on any changes to the code base