# **Final Project University**

Final project journey to get Bachelor Degree of Computer Science with topic  `Text Summarization` in domain Data Science.

Stacks : Golang, MySQL, Nginx, Ubuntu...

> "Otw Get Bachelor Degree of Computer Science at Telkom University" - 
> Ananda Affan Fattahila

**Migrate Create**
Create a migration file. You can find the file at `migration` folder
```
make migrate-create NAME=namefile
```

**Migrate Up**
To migrate all your migration file
```
make migrate-up
```

**Migrate Down**
To delete all your schema with migration
```
make migrate-down
```

**Migrate Rollback**
to run migration down only `N` step(s) behind
```
make migrate-rollback N=yournumberrunmigrationdown
```

**Fixing your Migration**  
What happend if your database is dirty?
You can fix your migration first and then using foce command with the version you want.
If you're happend to get `error: Dirty database version 16. Fix and force version.`
Then you want to run:
```
make migrate-force VERSION=15
```
Reference: https://github.com/golang-migrate/migrate/issues/282#issuecomment-530743258