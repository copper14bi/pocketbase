## v0.38.0

- Fixed UI logs pagination when no custom range is specified.

- Fixed default CSP not allowing audio/video previews ([#7677](https://github.com/pocketbase/pocketbase/issues/7677)).

- Serve fixed `Content-Type` for `.xlsx`, `.docx` and `.pptx` files to allow previews on iOS ([#7467](https://github.com/pocketbase/pocketbase/discussions/7467)).

- Changed settings app URL input to `type="text"` for compatibility with earlier versions ([#7681](https://github.com/pocketbase/pocketbase/issues/7681)).

- Added an internal watcher to sync various runtime states between multiple PocketBase processes (e.g. memory store) using the same `pb_data`.
    _This is helpful in case for example a separate PocketBase console command change the collections or application settings while the server is still running._
    _The watcher is debounced and implemented by watching the special `pb_data/.notify` dir as a workaround to avoid depending on OS and SQLite driver specific APIs._

- Added new [Superuser IPs/CIDR subnets whitelist setting](https://pocketbase.io/docs/going-to-production/#limit-superusers-to-specific-ipssubnets).
    The optional setting can be changed from the UI under _Dasboard > Settings > Application > Superuser IPs_.
    To avoid lockout in case your superuser IP change, the ips whitelist can be updated also via the `superuser ips` console command:
    ```sh
    # note: --dir is optional and defaults to pb_data next to the executable

    # clear whitelisted IPs
    ./pocketbase superuser ips --dir=/custom/path/to/pb_data

    # OR change the whitelisted IPs to 127.0.0.1 and 10.0.0.0 (replace with your real IP(s))
    ./pocketbase superuser ips 127.0.0.1 10.0.0.0 --dir=/custom/path/to/pb_data
    ```

- Added rate limit option to exclude IPs/CIDR subnets ([#6410](https://github.com/pocketbase/pocketbase/issues/6410)).

- Bumped min Go GitHub action version to 1.26.3 because it comes with some [minor bug and security fixes](https://github.com/golang/go/issues?q=milestone%3AGo1.26.3).

> **Personal note:** Upgrading to this version — the superuser IP whitelist feature is exactly what I needed for my home server setup. Remember to set this after deploying so only my home IP can access the admin dashboard.
> Also worth noting: the internal watcher for `pb_data/.notify` is a neat approach — avoids inotify limits that can be a pain on low-resource VPS setups.

## v0.37.5

- Fixed password fields not being detected as changed ([#7670](https://github.com/pocketbase/pocketbase/issues/7670)).

- Added the local time zone name next to the `date` field label.

- Reload trusted proxy info UI after settings save.

- Other minor improvements (skips the duplicated record ids from the `IN` expand list, reordered confirm-email-change error checks to minimize enumeration attacks, etc.).


## v0.37.4

- Added backups list scroll container ([#7655](https://github.com/pocketbase/pocketbase/issues/7655)).

- Optimized record upsert and preview modals data loading to minimize layout jumps.

- Fixed SMTP IPv6 network address format ([#7659](https://github.com/pocketbase/pocketbase/issues/7659)).

- Fixed autoco