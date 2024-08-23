# Estonian Companies and Persons Data

Latest data about every Estonian company and associated persons in one SQLite database. 
The data is gathered from many public sources like [this](https://avaandmed.ariregister.rik.ee/et/avaandmete-allalaadimine).
The database is automatically regenerated every day, and the [latest version is available for download here](https://github.com/karelnagel/avaandmed/releases/latest), or if you'd like to automate the process, you can use the [download.sh](download.sh) script or the following command:

```bash
curl -s https://api.github.com/repos/karelnagel/avaandmed/releases/latest | jq -r '.assets[0].browser_download_url' | xargs -I {} curl -s -L {} | gunzip -c > out.db
```


## Parser usage

Parser is written in Go, so you need to have Go installed.

```bash
go install
```

```bash
go run main.go
```

options:
```bash
  -batch int
     Batch size (default 600)
  -fail-quietly
     Fail quietly
  -force-download
     Force downloading the latest data again, eg. it deletes the data directory
  -sources string
     Sources to process (comma separated) (default "yldandmed,kaardile_kantud,kandevalised,kasusaajad,osanikud,majandusaasta,emta,debt,lihtandmed")
  -sqlite string
     Path to the SQLite database (default "out.db")
```
