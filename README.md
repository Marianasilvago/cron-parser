
# Cron Expression Parser

This project is a command-line application written in Golang that parses cron expressions and outputs the times at which the cron jobs will run. It utilizes the Cobra library for building the CLI.

The application can handle a subset of standard cron expressions with 5 fields: minute, hour, day of month, month, and day of week, followed by a command. It expands the cron fields and prints the expanded schedule in a readable format.

## Features

- Parses cron strings and outputs expanded schedules.
- Supports ranges, steps, and specific values in cron fields.
- Provides a CLI interface with a `parse` command to process cron expressions.
- Unit tests to validate the parser logic.

## Getting Started

### Prerequisites

To run this project, you need the following installed:

- **Golang** (version 1.18+ recommended)

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/Marianasilvago/cron-parser.git
   ```

2. Navigate into the project directory:

   ```bash
   cd cron-parser
   ```

3. Install dependencies:

   Since Cobra is already imported, you can tidy up the dependencies:

   ```bash
   go mod tidy
   ```

### Usage

To run the parser, use the `parse` command with a cron expression. The cron expression should be provided as a single argument in quotes.

#### Example:

To parse a cron expression that runs every 15 minutes:

```bash
go run main.go parse "*/15 0 1,15 * 1-5 /usr/bin/find"
```

To Run the compiled application:

```bash
go ./cron-parser parse "*/15 0 1,15 * 1-5 /usr/bin/find"
```

#### Output:

```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

In this case, the cron job runs every 15 minutes (00, 15, 30, 45) at midnight on the 1st and 15th of the month, from Monday to Friday.

### Commands

- **parse**: The main command to parse and expand cron expressions.
  
  Usage:
  
  ```bash
  go run main.go parse "*/15 0 1,15 * 1-5 /usr/bin/find"
  ```

  Example output:

  ```bash
  minute        0 15 30 45
  hour          0
  day of month  1 15
  month         1 2 3 4 5 6 7 8 9 10 11 12
  day of week   1 2 3 4 5
  command       /usr/bin/find
  ```

### Testing

Unit tests are provided to validate the cron parser functionality. You can run the tests with the following command:

```bash
go test ./...
```

This will run all the tests and display the results.

### Example Output:

```bash
ok  	cron-parser/cmd	0.005s
```

If any test fails, the output will detail the failed test case.

## Cron Syntax Reference

The application supports cron expressions that include:

- **Minute (0-59)**: The minute field specifies the minute(s) the job should run.
- **Hour (0-23)**: The hour field specifies the hour(s) the job should run.
- **Day of Month (1-31)**: The day of the month field specifies the day(s) of the month.
- **Month (1-12)**: The month field specifies the month(s) of the year.
- **Day of Week (0-6)**: The day of the week field specifies the day(s) of the week (0 = Sunday).

### Examples

- `*/15 0 1,15 * 1-5 echo 'Hello World'`: Runs every 15 minutes at midnight on the 1st and 15th of the month, Monday to Friday.
- `0 12 * * 0 echo 'Run at noon on Sundays'`: Runs at 12:00 noon every Sunday.
  
For more information about cron syntax, refer to [crontab.guru](https://crontab.guru).

## Project Structure

```
.
├── cmd
│   └── parse.go         # Contains the cron parser logic
├── main.go              # Entry point for the CLI application
├── parse_test.go        # Unit tests for the parser
└── go.mod               # Module dependencies
```

## Future Improvements

- Add support for handling cron special strings like `@yearly`, `@daily`, `@hourly`.
- Provide detailed error messages for invalid cron expressions.
- Add support for different time zones.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
