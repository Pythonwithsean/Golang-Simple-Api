# Use official MySQL image as base image
FROM mysql:latest

# Set the root password for MySQL
ENV MYSQL_ROOT_PASSWORD=root

# Create a database
ENV MYSQL_DATABASE=mydatabase

# Expose MySQL port
EXPOSE 3306
# Define a volume for MySQL data persistence
VOLUME /var/lib/mysql

# Health check for MySQL container
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD mysqladmin ping -uroot -proot || exit 1

# Copy the SQL script to initialize the database
COPY init.sql /docker-entrypoint-initdb.d/
