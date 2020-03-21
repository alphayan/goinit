package temp

const DOCKER_COMPOSE = `version: '3.7'
services:
  mysql:
    image: mysql
    ports:
      - "3306:3306"
      - "33060:33060"
    volumes:
      - ~/data/mysql:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: root123456
    restart: unless-stopped
  redis:
    image: redis
    ports:
      - "6379:6379"
    restart: unless-stopped
    volumes:
      - ~/data/redis:/data
  mongo:
    image: mongo
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: root123456
    ports:
      - "27017:27017"
      - "27018:27018"
      - "27019:27019"
      - "28017:28017"
    restart: unless-stopped
    volumes:
      - ~/data/mongo:/data/db
  postgres:
    image: postgres
    ports:
      - "5432:5432"
    environment:
      POSTGRES_PASSWORD: root123456
    restart: unless-stopped
    volumes:
      - ~/data/postgresql:/var/lib/postgresql/data
  {{.APP}}:
    image: {{.APP}}
    ports:
      - "8080:8080"
    links:
      - mysql
      - redis
      - mongo
      - postgres
    depends_on:
      - {{.DB}}
      - redis
    restart: unless-stopped
`
