# Blueprints
Blueprints are a way of defining reusable directory and file structures. They can be used to create consistent setups for projects, applications, or any other purpose.

## Blueprint Files
Blueprints are defined in YAML files. (both `.yaml` and `.yml` are supported).

Blueprints are created at `~/.config/atools/blueprints`.

Make sure to give your blueprint file a unique, easy-to-remember name. This name will be used to reference the blueprint.

## Variables
### Custom Variables
Blueprints can contain variables that can be replaced with actual values when the blueprint is used. Variables are defined in the format `{{ variable_name }}`.

### Predefined Variables
Predefined variables are automatically replaced with their values when the blueprint is used.

- `{{ year }}` - the current year.
- `{{ quarter }}` - the quarter of the year we are in (e.g. Q1).
- `{{ month }}` - the current month.
- `{{ monthnumber }}` - the number of the current month.
- `{{ weeknumber }}` - the number of the current week.
- `{{ day }}` - the current day.
- `{{ daynumber }}` - the number of the current date.
- `{{ hour }}` - the current hour.
- `{{ minute }}` - the current minute.

## Example
```yaml
blueprint:
    directories:
        - 'test'
        - '{{ year }}-{{ monthnumber }}-{{ daynumber }}-test'
    files:
        - path: 'test/test.txt'
        - path: '{{ year }}-{{ monthnumber }}-{{ daynumber }}-test/test.md'
          content: |
            # This is a test file
            This file was created on {{ year }}-{{ monthnumber }}-{{ daynumber }}
```
