# Konnec File Declations

This file describes Konnec Features

## Inventory.yaml

This topic describes inventory decleartions in inventory.yaml

```
inventories:
  - name: active_directory_1
	ip: 192.168.1.33
	domain: dc01.example.com
	description: Primary active directory server

  - name: sql_server
	ip: 192.168.1.22
	description: Primary SQL Server

  - name: mongodb
	ip: mongodb.example.com
	description: Single Node Mongo DB Server

## Secrets.yaml

This topic describes secret declerations in secrets.yaml

### Clear Text Values

```yaml
sql:
 - username: admin
   password: 123456
```

### Base64 Values

```yaml
active_directory:
  - username: admin
    password: b:MTIzNDU2
mongo_db:
  - username: b:YWRtaW4=
    password: b:MTIzNDU2
```

## Check.yaml

Theese sub-topics describes service declerations in check.yaml

### DNS Check

Checking DNS resolution answers. 

```yaml
dns:
   - name: gateway
     target: gateway.example.com
     value: 192.168.1.43
   - name: active-directory
     target: dc.example.com
     value: 192.168.1.33
```

### PING Check

Checking any machine accessiblity by ICMP.

```yaml
ping:
   - name: sql-server
     target: 192.168.1.22
   - name: git-server
     target: git.example.com
```

### MSSQL Check

Check SQL Server connection.

```yaml
mssql:
 - name: primary-sql-server
   target: 192.168.1.22
   port: 1433
   username: {{ sql.username }}
   password: {{ sql.password }}
```

### MongoDB Check

### Active Directory Check

### Docker Image Registry

### Kafka Check


