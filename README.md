# todolint

Go linter for checking that `TODO` comments in source code match a format:
`TODO(<context>): <summary>`.

## Rationale

The [godox] linter prevents `TODO` etc comments with the explanation _"You
should create tasks if some TODOs cannot be fixed in the current merge
request"_. While the goal here is good, this leaves no trace in the code for
this task. Instead, the `todolint` linter can permit these comments, if the
provide additional _context_. This context can be the link to the task.

Sometimes a separate task may be overkill. For this, it can be OK to leave the
`TODO`, but with the author as additional context so that the person can be
asked instead when somebody is attempting to find the origin.

## Context

The context provides additional context for what needs to change. This can be a
ticket number in an issue tracker or maybe a full URL for an external
dependency. Sometimes it's just something small, in which case the author name
`// TODO(akupila): This is an example` is sufficient. Some context must be
provided, so that it's easier in the future to trace the origins.

By default, the context is just set to not be empty. You can require a stricter
context by passing a regular expression for the `context`, such as `^#\d+$` for
a GitHub issue or `^https://` for a URL.

## Keywords

The linter looks for `TODO`, `FIXME` and `BUG` comments by default and ignores
any comments that don't include these keywords. Only single line (`//`)
comments are included. You can add other keywords by setting `keywords` to
another value, such as `TODO,OPTIMIZE`.

### Why not `git blame`?

Git blame works in the short term, but after a few cycles of refactoring,
reformatting or other bigger changes, it's not immediately obvious who wrote
the comment. By enforcing the comment itself to include the author name, it's
easier to ask that person directly. All problems don't need to be solved with
technology.

## Format

The format is set to match the `// BUG(who): summary` format described in
[documenting Go code], except instead of `who` we use a more generic `context`.
The `context` allows linking to a ticket, which is often useful in commercial
projects.

The linter can automatically fix `// TODO(foo)bar` -> `// TODO(foo): bar`.

## Credits

Thanks to [@rcambrj] for feedback on the idea.



[godox]: https://github.com/matoous/godox
[documenting Go code]: https://go.dev/blog/godoc
[@rcambrj]: https://github.com/rcambrj
