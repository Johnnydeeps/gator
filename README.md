![readme title image](images/readme_image.png)

# Gator

A CLI tool for aggregating RSS feeds.

## Prerequesites 

- [Go](https://go.dev/doc/install)
- [Postgres](https://www.postgresql.org/download/)

## Installation

```bash
go install github.com/Johnnydeeps/gator@latest
```

## Configuration

This program requires a config file named `.gatorconfig.json` located in your home directory, with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "yourusername"
}
```

When you register for the first time the `"current_user_name": "yourusername"` will be populated automatically.

## Commands & Usage

In your terminal/CLI use the following commands:

- `gator register <username>` - registers a new user, automatically logs you in.
- `gator login <username>` - login as an existing user.
- `gator reset` - resets the tables that are stored in the database.
- `gator users` - lists all stored users that are in the database.
- `gator addfeed <URL>` - adds the URL of a specific RSS feed to the database. 
- `gator feeds` - lists all stored feeds that are in the database.
- `gator follow <URL>` - when you are logged in, allows the user to follow a feed stored in the database.
- `gator following` - when you are logged in, displays all the RSS feeds you are following.
- `gator unfollow <URL>` -  unfollow 
- `gator agg <time>` - suggest 10m as a default starting time, will retrieve all followed feeds for a given
                   user that is logged in.
- `gator browse` - browse all retrieved and aggregated posts from the RSS feeds that are followed by a user.

## Future Improvements

Currently requires manual creation of the config file; a future version could add a gator config init command. This would remove the required configuration steps listed above.
