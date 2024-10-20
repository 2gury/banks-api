FROM gomicro/goose

# Copy the migrations directory to /migrations in the container
COPY ./migrations /migrations
# Copy the migration.sh script to /migration/migration.sh in the container
COPY migration.sh /migration/migration.sh

WORKDIR /migration

# Ensure the migration.sh script has executable permissions
RUN chmod +x /migration/migration.sh

ENTRYPOINT ["/bin/bash", "-c", "./migration.sh"]