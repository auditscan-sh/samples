# Intentionally bad multi-stage Java build.
FROM maven:latest AS build
WORKDIR /src
COPY . .
RUN mvn -DskipTests package

FROM openjdk:latest
USER root
RUN apt-get update && apt-get install -y curl netcat-openbsd sudo && rm -rf /var/lib/apt/lists/*
ENV SPRING_PROFILES_ACTIVE=prod
ENV JWT_SECRET=9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV
ENV DB_PASSWORD=Pr0d-B1ll1ng-2024!
ENV STRIPE_SECRET_KEY=sk_live_REDACTED_FOR_PUSH
COPY --from=build /src/api/target/*.jar /app/api.jar
EXPOSE 22 8080 5005
CMD ["java", "-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005", "-jar", "/app/api.jar"]
