
FROM beevelop/ionic:latest

# -----------------------------------------------------------------------------
# General environment variables
# -----------------------------------------------------------------------------
ENV DEBIAN_FRONTEND=noninteractive


# -----------------------------------------------------------------------------
# WORKDIR is the generic /app folder. All volume mounts of the actual project
# code need to be put into /app.
# -----------------------------------------------------------------------------
WORKDIR /front_end

COPY . .

# Expose port 8100 to the outside world
EXPOSE 8100

CMD ["ionic", "serve", "-b", "-p", "8100", "--address", "0.0.0.0"]
