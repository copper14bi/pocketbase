> List with changes from v0.8.x to v0.15.x.
> For the most recent versions, please refer to [CHANGELOG.md](./CHANGELOG.md)
---

## v0.15.3

- Updated the Admin UI to use the latest JS SDK to resolve the `isNew` record field conflict ([#2385](https://github.com/pocketbase/pocketbase/discussions/2385)).

- Fixed `editor` field fullscreen `z-index` ([#2410](https://github.com/pocketbase/pocketbase/issues/2410)).

- Inserts the default app settings as part of the system init migration so that they are always available when accessed from within a user defined migration ([#2423](https://github.com/pocketbase/pocketbase/discussions/2423)).


## v0.15.2

- Fixed View query `SELECT DISTINCT` identifiers parsing ([#2349-5706019](https://github.com/pocketbase/pocketbase/discussions/2349#discussioncomment-5706019)).

- Fixed View collection schema incorrectly resolving multiple aliased fields originating from the same field source ([#2349-5707675](https://github.com/pocketbase/pocketbase/discussions/2349#discussioncomment-5707675)).

- Added OAuth2 redirect fallback message to notify the user to go back to the app in case the browser window is not auto closed.


## v0.15.1

- Trigger the related `Record` model realtime subscription events on [custom model struct](https://pocketbase.io/docs/custom-models/) save ([#2325](https://github.com/pocketbase/pocketbase/discussions/2325)).

- Fixed `Ctrl + S` in the `editor` field not propagating the quick save shortcut to the parent form.

- Added `⌘ + S` alias for the record quick save shortcut (_I have no Mac device to test it but it should work based on [`e.metaKey` docs](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/metaKey)_).

- Enabled RTL for the TinyMCE editor ([#2327](https://github.com/pocketbase/pocketbase/issues/2327)).

- Reduced the record form vertical layout shifts and slightly improved the rendering speed when loading multiple `relation` fields.

- Enabled Admin UI assets cache.


## v0.15.0

- Simplified the OAuth2 authentication flow in a single "all in one" call ([#55](https://github.com/pocketbase/pocketbase/issues/55)).
  Requires JS SDK v0.14.0+ or Dart SDK v0.9.0+.
  The manual code-token exchange flow is still supported but the SDK method is renamed to `authWithOAuth2Code()` (_to minimize the breaking changes the JS SDK has a function overload that will proxy the existing `authWithOauth2` calls to `authWithOAuth2Code`_).
  For more details and example, you could check https://pocketbase.io/docs/authentication/#oauth2-integration.
  > **Personal note:** This was a breaking change in my local project - remember to update any direct calls to `authWithOAuth2` to use `authWithOAuth2Code` if not relying on the SDK shim.

- Added support for protected files ([#215](https://github.com/pocketbase/pocketbase/issues/215)).
  Requires JS SDK v0.14.0+ or Dart SDK v0.9.0+.
  It works with a short lived (~5min) file token passed as query param with the file url.
  For more details and example, you could check https://pocketbase.io/docs/files-handling/#protected-files.

- ⚠️ Fixed typo in `Record.WithUnkownData()` -> `Record.WithUnknownData()`.

- Added simple loose wildcard 
