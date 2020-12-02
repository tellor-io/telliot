---
description: Release cadence and process for the project.
---

# Release Process

We use [Semantic Versioning](http://semver.org/).

NOTE: As [Semantic Versioning](http://semver.org/spec/v2.0.0.html) states all 0.y.z releases can contain breaking changes in API (flags, grpc API, any backward compatibility)

## Cadence

* We aim to keep the master branch in a working state at all times. In principle, it should be possible to cut a release from master at any time. In practice, things might not work out as nicely. A few days before the pre-release is scheduled, the shepherd should check the state of master. Following their best judgement, the shepherd should try to expedite bug fixes that are still in progress but should make it into the release. On the other hand, the shepherd may hold back merging last-minute invasive and risky changes that are better suited for the next minor release.
* No feature should block release.

### Branch management and versioning strategy
We use [Semantic Versioning](https://semver.org/).

We are a small team and don't have the capacity to maintain patch release branches so will create branches only for major releases: `release-<major>`, e.g. `release-1`, `release-2`.

Note that branch protection kicks in automatically for any branches whose name starts with `release-` so never use names starting with `release-` for branches that are not release branches.

### Updating dependencies
A few days before a release, consider updating the dependencies:

```
make update-go-deps
git add go.mod go.sum
git commit -m "Update dependencies"
```

Then create a pull request against the master branch.

Note that after a dependency update, you should look out for any weirdness that
might have happened. Such weirdnesses include but are not limited to: flaky
tests, differences in resource usage, panic.

In case of doubt or issues that can't be solved in a reasonable amount of time,
you can skip the dependency update or only update select dependencies. In such a
case, you have to create an issue or pull request in the GitHub project for
later follow-up.

# Cutting a release

## Update `CHANGELOG.md`
Do this in a PR against `master` as this gives others the opportunity to chime in on the release and the changelog changes.

Note:
> that `CHANGELOG.md` should only document changes relevant to users of the project, including external API changes, performance improvements, and new features. Do not document changes of internal interfaces, code refactorings and clean-ups, changes to the build process, etc. People interested in these are asked to refer to the git history.

## Draft the new release
On `master` tag the current commit as a new release via the following commands:

```bash
$ export TAG="v0.1.0"
$ git tag $TAG
$ git push origin $TAG
```

Once a tag is pushed, the release process is triggered through the CI and it will draft the GitHub release and upload all artifacts.
If all looks good click _Publish release_. This will make the release publicly visible and create a GitHub notification.

## Wrapping up

Announce the release on the communication channels.<br/>
Check out previous announcements for inspiration.
