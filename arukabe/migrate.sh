#!/bin/bash
# usage migrate.sh up|down

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

export DATABASE="postgresql://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:5432/${PG_DB}?sslmode=disable"

# Function to print usage
print_usage() {
    echo "Usage: $0 <command> [<n>]"
    echo "Commands:"
    echo "  up [<n>]     Migrate up (apply n migrations if specified)"
    echo "  down [<n>]   Migrate down (revert n migrations if specified)"
    echo "  goto <v>     Migrate to version v"
    echo "  force <v>    Set version v but don't run migration (ignores dirty state)"
    echo "  version      Print current migration version"
}

if [ $# -eq 0 ]; then
		print_usage
		exit 1
fi

command=$1
n=$2

case $command in
	up|down)
		if [ -n "$n" ]; then
			migrate -database ${DATABASE} -path migrations ${command} ${n}
		else
			migrate -database ${DATABASE} -path migrations ${command}
		fi
		;;
	goto|force|version)
		if [ -n "$n" ]; then
			migrate -database ${DATABASE} -path migrations ${command} ${n}
		else
			echo "Error: $command requires a version number"
			print_usage
			exit 1
		fi
		;;
	*)
		echo "Error: Unknown command '$command'"
		print_usage
		exit 1
		;;
esac

# ./migrate.sh up        # Apply all pending migrations
# ./migrate.sh up 1      # Apply the next 1 migration
# ./migrate.sh down      # Revert the last applied migration
# ./migrate.sh down 2    # Revert the last 2 migrations
# ./migrate.sh goto 5    # Migrate to version 5
# ./migrate.sh version   # Print the current migration version
