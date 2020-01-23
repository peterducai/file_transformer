# file_transformer
utility to quickly replace text in files, configs and so on..

## DONE:

- simple text conversion with single value

## INSTALL

> go get github.com/peterducai/file_transformer


which will transform testfile.txt to test_new.txt

## TODO:

### JSON

accept JSON as input

For example JSON like

```json
{
    "name":"John",
    "age":30,
    "cars": {
        "car1":"Ford",
        "car2":"BMW",
    }
 }
 ```

 will replace **${NAME.CARS.CAR1}** with value *Ford*.

### YAML

accept YAML as input
