


```shell
# Ready
$ mongoman user list [--dbname <database name>] --uri <mongodb-uri>
$ mongoman user create --username <user name> --password <password> --roles <roles> [--dbname <database name>] --uri <mongodb-uri>
$ mongoman user delete --username <user name> [--dbname <database name>] --uri <mongodb-uri>

# Todo
$ mongoman db list
$ mongoman db create <database name> --uri <mongodb-uri>
$ mongoman db del <database name> --uri <mongodb-uri>
$ mongoman coll list
$ mongoman coll create <collection name> --dbname <database name> --uri <mongodb-uri>
$ mongoman coll del <collection name> --dbname <database name> --uri <mongodb-uri>

```