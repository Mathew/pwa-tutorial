Implementation of the PWA tutorial from [Codelabs](https://codelabs.developers.google.com/codelabs/your-first-pwapp/#0).

Uses Go to serve the files.


Running
-------

* `go run main.go`
* Visit http://localhost:8080 in a browser.


Notes
-----

* Ensure all files are named correctly throughout your offline.html, manifest and service-worker references.
* Use dev-tools in chrome, particularly the Application tools allowing you to simulate the worker offline and clear it.
* Use the Lighthouse plugin for Chrome to check your site's performance.