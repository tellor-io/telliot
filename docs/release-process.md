# Release Process

This page describes the release cadence and process for the project.

We use [Semantic Versioning](http://semver.org/).

NOTE: As [Semantic Versioning](http://semver.org/spec/v2.0.0.html) states all 0.y.z releases can contain breaking changes in API (flags, grpc API, any backward compatibility)

## Cadence

* We aim to keep the master branch in a working state at all times. In principle, it should be possible to cut a release from master at any time. In practice, things might not work out as nicely. A few days before the pre-release is scheduled, the shepherd should check the state of master. Following their best judgement, the shepherd should try to expedite bug fixes that are still in progress but should make it into the release. On the other hand, the shepherd may hold back merging last-minute invasive and risky changes that are better suited for the next minor release.
* The release shepherd cuts the release and creates a new branch called  `release-<major>.<minor>` starting at the commit tagged for the release.
* No feature should block release.

See the next section for details on cutting an individual release.

### Branch management and versioning strategy

We use [Semantic Versioning](https://semver.org/).

We maintain a separate branch for each minor release, named `release-<major>.<minor>`, e.g. `release-1.1`, `release-2.0`.

Note that branch protection kicks in automatically for any branches whose name starts with `release-`. Never use names starting with `release-` for branches that are not release branches.

The usual flow is to merge new features, changes and bugfixes into the master branch and to backport bug fixes into the latest release branch at best effort possible.


### 0. Updating dependencies

A few days before a major or minor release, consider updating the dependencies:

```
make update-go-deps
git add go.mod go.sum vendor
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

At the start of a new major or minor release cycle create the corresponding release branch based on the master branch. For example if we're releasing `0.2.0` and the previous stable release is `0.1.0` we need to create a `release-0.2` branch.

```bash
$ git checkout -b release-0.2
$ git push origin release-0.2
```

Note that all releases are handled in protected release branches. Patch releases for any given major or minor release happen in the same `release-<major>.<minor>` branch. Do not create a new branch for patch releases.

Update `CHANGELOG.md`. Do this in a proper PR pointing to the release branch as this gives others the opportunity to chime in on the release in general and on the addition to the changelog in particular.

Note that `CHANGELOG.md` should only document changes relevant to users of the project, including external API changes, performance improvements, and new features. Do not document changes of internal interfaces, code refactorings and clean-ups, changes to the build process, etc. People interested in these are asked to refer to the git history.

### 2. Draft the new release

Tag the new release via the following commands:

```bash
$ export TAG="v0.1.0"
$ git tag $TAG
$ git push origin $TAG
```

Optionally, you can use this handy `.gitconfig` alias.

```ini
[alias]
  tag-release = "!f() { git tag $TAG && git push origin $TAG; }; f"
```

Then release with `git tag-release`.

Once a tag is created, the release process through the CI will be triggered for this tag and the CI will draft the GitHub release using a `bot` account.

Finally, wait for the build step for the tag to finish. The point here is to wait for tarballs to be uploaded to the Github release and the container images to be pushed to the Docker Hub and Quay.io. Once that has happened, click _Publish release_, which will make the release publicly visible and create a GitHub notification.

### 3. Wrapping up

Once the binaries have been uploaded, announce the release on the communication channels. Check out previous announcement mails for inspiration.
