# orchestra 0.3.2 [2020-01-21]

Changes:

* Go 1.13 support including migrating to Go modules
* Fetch tor targets via authenticated orchestra (#83)

# orchestra 0.3.1 [2019-12-17]

Changes:

* add Dockerfile for the daemons
* add Dockerfile for frontend
* fix bug in list-target query (dfa1d7c527)
* bump frontend dependencies
* distribute psiphon config (#80)

# orchestra 0.3.0 [2019-06-05]

Changes:

* Returns by default all URLs in the test-lists/urls endpoint

# orchestra 0.2.5 [2019-01-22]

Changes:

* Support registering clients with an empty push notification token: https://github.com/ooni/orchestra/issues/41


# orchestra 0.2.4 [2018-12-13]

Changes:

* Validate test URLs when returning them and ignore invalid URLs

Fixes:

* Bug with test URLs endpoint: https://github.com/ooni/orchestra/issues/59
* Bug with test-helpers endpoint: https://github.com/ooni/orchestra/issues/54

# orchestra 0.2.2 [2018-10-22]

Adds:

* test-lists/urls endpoint
* Support for returning full list of URLs with limit=-1

# orchestra 0.2.1 [2018-07-13]

Changes:

* Rename proteus to orchestra
* Switch to dep for dependency management

Adds:

* Pagination support for result listing

# proteus 0.2.0 [2018-01-25]

Fixes:

* Update next.js in proteus-frontend to 3.2.3 to remediate: https://github.com/zeit/next.js/releases/tag/4.2.3 (the fix has been backported to 3.2.3 too)

# proteus 0.2.0-rc.2 [2018-01-16]

Fixes:

* Make unauthenticated endpoints work without auth header

# proteus 0.2.0-rc.1 [2017-12-14]

Changes:

* Rename proteus-events to proteus-orchestrate
* Remove the deprecated proteus-notify

Adds:

* New orchestrate endpoints for /test-helpers, /collectors and /urls

# proteus 0.1.0-rc.1 [2017-09-29]

Changes:

* Improvements to testing and code style

* Add unittests that speak to a database

# proteus 0.1.0-beta.9 [2017-08-06]

* fix(registry): Language may be null
* fix(database): inconsistent migrations

# proteus 0.1.0-beta.8 [2017-08-02]

* refactor(Makefile): simplify the proteus targets dependencies
* feature(README.md): build and release instruction
* feature(events): fwd `click_action` for Android
* fix(registry): adjust syntax of `add_language_column`
* fix(Makefile): don't assume a tool has bindata
* regen(bindata): run `make bindata`
* fix(Makefile): always update embedded binary data

# proteus 0.1.0-beta.7 [2017-07-30]

* fix(proteus-events): do not propagate topic on Android

# proteus 0.1.0-beta.6 [2017-07-24]

* Improvements to the scheduling UI

* Add minimal progress bar

# proteus 0.1.0-beta.5 [2017-07-19]

* Fix bug in migration script

* Re-enable measurement job scheduling from UI

# proteus 0.1.0-beta.4 [2017-07-17]

* Fix bug in deletion of jobs

# proteus 0.1.0-beta.3 [2017-06-23]

* Add support for alert notifications

* Temporarily disable measurement job scheduling from UI

* Use gorush instead of proteus-notify

# proteus 0.1.0-beta.2 [2017-05-28]

* Fix the schema of proteus-events

# proteus 0.1.0-beta.1 [2017-05-26]

* Add support for Admins to manage jobs via the web UI

# proteus 0.1.0-dev [2017-05-13]

Development release of proteus.

Includes:

* Ability for probes to register with the orchestration registry
* Ability for probes to update the metadata about their probe
* Support for sending push notifications via Apple Push Notifcation service and Firebase Cloud Messagging
* Support for administrators to schedule measurements via a web interface
* Ability for probes to receive the tasks they have been notified about and mark them as accepted, rejected or done
* Multiple architecture build system

