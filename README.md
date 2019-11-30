# Tracker

Measure weekly tasks in 2.5 hour blocks


## Usage
```
tracker <command> [options]

Commands:
  list
  reset
  start
Options:
  -t, --task int   The task you want to start.
```

## Example tasks.json
Each task requires three attributes: name, limit, and current. The name is just a string that comes up in the list of tasks. The limit is the maximum number of 2.5 hour blocks you want to spend working on that task each week. The current should be zero in your config, it is used by tracker to measure how many blocks you have completed.
```
{
	"tasks": [{
		"name": "Example Task 1",
		"limit": 2,
		"current": 0
	}, {
		"name": "Example Task 2",
		"limit": 3,
		"current": 0
	}]
}
```

