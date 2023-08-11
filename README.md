# cronTool

A Golang app that validates, parses and runs Cron schedules (each run is a simple text log) without the use of third-party Cron-specific libraries .

> Usage: cronTool -s '<cron_schedule>'

The `<cron_schedule>` parameter must be a string containing a Cron expression.

A cron expression is a string comprised of 7 fields separated by white space (leading and trailing whitespace in the `<cron_schedule>` string will be disregarded and any extra whitespaces will be sanitized during parsing).

Fields can contain any of the allowed values, along with various combinations of the allowed special characters for that field.

The fields are as follows:

|Field Name|Allowed Values|Allowed Special Characters|
|--|--|--|
|Seconds|0-59|`, - * /`|
|Minutes|0-59|`, - * /`|
|Hours|0-24|`, - * /`|
|Day of Month|1-31|`, - * /`|
|Month|1-12|`, - * /`|
|Day of Week|0-6|`, - * /`|
|Year|1970-2099|`, - * /`|

## Meaning of special characters:

**- Comma (`,`):** used to separate items of a list. For example, using `1,3,5` in the 6th field (day of week) means Mondays, Wednesdays and Fridays;
**- Dash (`-`):** defines ranges. For example, `2000–2010` indicates every year between 2000 and 2010, inclusive;
**- Asterisk (`*`):** also known as wildcard, represents "all". For example, using `* * * * * * *` will run every minute. Using `* * * * * 1 *` will run every minute only on Monday;
**- Slash (`\`):** can be combined with ranges to specify step values. For example, `*/5` in the minutes field indicates every 5 minutes. It is shorthand for the more verbose form `5,10,15,20,25,30,35,40,45,50,55,0`.

## Unsupported features:
- Cron expressions with length 5 and 6;
- Digit `7` for Sundays;
- Special characters `?`, `#` and `%`;
- Characters `L` and `W`;
- Abbreviations `JAN–DEC` and `SUN–SAT`;
-  `@yearly`  `@annually`  `@monthly`  `@weekly`  `@daily`  `@midnight`  `@hourly`  `@reboot` scheduling definitions.

## :white_check_mark: Examples of valid expressions
-  **`1 0 * * * * *`:** will run at one minute past midnight (00:01) every day;
-  **`45 23 * * 6 * *`:** will run at 23:45 (11:45 PM) every Saturday;
-  **`*/5 * * * * * *`:** will run at every 5s;
-  **` 0,10,20,30,40,50 * * * * * `:** will run at 0, 10, 20, 30, 40, and 50 seconds past the minute (valid one due to disregarded leading and trailing whitespace and extra whitespaces sanitized);
-  **`3-15 * * * * * *`:** will run at every minute from 3 through 15.

## :no_entry: Examples of invalid expressions
-  **`60 * * * * * *`:** invalid value for seconds field;
-  **`* * * * *`:** Cron expression with length 5;
-  **`* * * * * * 1969`:** invalid value for year field;
-  **`2-3/10 * * * * * * * "`:** invalid combination of allowed special characters in seconds field;
-  **`"? invalid * m * * 2023"`:** characters not allowed.

## Running tests:
- To run all unit tests, use `go test ./... -cover -coverprofile=coverage.out` in your terminal;
- To view the test report, use `go tool cover -html=coverage.out` in your terminal.
- To run all benchmarks, use `go test -bench . ./...` in your terminal;
- The main method has only one benchmark and does not have unit tests due to the complexity of retrieving the buffer and the output code in Go to assert the results, not to mention that due to good testing practices it makes no sense to test the main method, since the other individual components of the application already have their respective unit tests.

## Caveats

Handling Cron expressions and simulating a Cron scheduler without using specific libraries is quite challenging and therefore not all non-standard Cron features have been implemented.

To validate the Cron expressions I created a specific regular expression, however it may not be very accurate, as well as all the code must have flaws, as it has not been thoroughly tested and not analyzed to ensure that every possible combination of fields in the Cron expressions are scheduled and run correctly.

## References:
- https://en.wikipedia.org/wiki/Cron
- https://www.quartz-scheduler.net/documentation/quartz-3.x/tutorial/crontriggers.html
